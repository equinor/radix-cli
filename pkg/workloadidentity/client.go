package workloadidentity

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	"github.com/equinor/radix-cli/pkg/auth"
	kiotaauth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
)

type ServicePrincipalResourceType string

const (
	ServicePrincipalResourceManagedIdentity ServicePrincipalResourceType = "managed-identity"
	ServicePrincipalResourceAppRegistration ServicePrincipalResourceType = "app-registration"
)

type ServicePrincipalResource struct {
	Type                                ServicePrincipalResourceType
	ServicePrincipal                    models.ServicePrincipalable
	ManagedIdentity                     *armmsi.Identity
	ManagedIdentityFederatedCredentials []*armmsi.FederatedIdentityCredential
	Application                         models.Applicationable
	ApplicationFederatedCredentials     []models.FederatedIdentityCredentialable
}

type GraphHelper struct {
	credential azcore.TokenCredential
	// azCreds         *azidentity.AzureCLICredential
	userClient      *msgraphsdk.GraphServiceClient
	graphUserScopes []string
}

func NewGraphHelper() *GraphHelper {
	g := &GraphHelper{}
	return g
}

func (graphHelper *GraphHelper) Client() *msgraphsdk.GraphServiceClient {
	return graphHelper.userClient
}

func (g *GraphHelper) InitializeGraphForUserAuth() error {
	scopes := "https://graph.microsoft.com/.default"
	// scopes := "Application.Read.All"
	g.graphUserScopes = strings.Split(scopes, ",")

	auth2, err := auth.New()
	if err != nil {
		return err
	}
	g.credential = &tokenCredentialAdapter{auth: auth2}
	// Create an auth provider using the credential
	authProvider, err := kiotaauth.NewAzureIdentityAuthenticationProviderWithScopes(g.credential, g.graphUserScopes)
	if err != nil {
		return err
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return err
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)
	g.userClient = client

	return nil
}

type tokenCredentialAdapter struct {
	auth *auth.Auth
}

func (a *tokenCredentialAdapter) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	if a.auth == nil {
		return azcore.AccessToken{}, fmt.Errorf("auth not set")
	}

	t, err := a.auth.GetAccessToken(ctx, options.Scopes)
	if err != nil {
		return azcore.AccessToken{}, err
	}

	return azcore.AccessToken{Token: t.Token, ExpiresOn: t.ExpiresOn}, nil
}

// ResolveServicePrincipalResource resolves a service principal appId and fetches
// the backing resource details as either a managed identity or app registration.
func (g *GraphHelper) ResolveServicePrincipalResource(ctx context.Context, appID string) (*ServicePrincipalResource, error) {
	appID = strings.TrimSpace(appID)
	if appID == "" {
		return nil, errors.New("service principal appId is required")
	}

	servicePrincipal, err := g.findServicePrincipalByAppID(ctx, appID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	resourceType := classifyServicePrincipalResourceType(servicePrincipal)
	result := &ServicePrincipalResource{
		Type:             resourceType,
		ServicePrincipal: servicePrincipal,
	}

	if resourceType == ServicePrincipalResourceManagedIdentity {
		subscriptionID, resourceGroupName, resourceName, err := parseManagedIdentityResourceID(servicePrincipal.GetAlternativeNames())
		if err != nil {
			return nil, err
		}

		managedIdentity, err := g.GetManagedIdentity(ctx, subscriptionID, resourceGroupName, resourceName)
		if err != nil {
			return nil, err
		}

		managedIdentityFederatedCredentials, err := g.getManagedIdentityFederatedCredentials(ctx, subscriptionID, resourceGroupName, resourceName)
		if err != nil {
			return nil, err
		}

		result.ManagedIdentity = managedIdentity
		result.ManagedIdentityFederatedCredentials = managedIdentityFederatedCredentials
		return result, nil
	}

	application, err := g.findApplicationByAppID(ctx, appID)
	if err != nil {
		return nil, err
	}

	applicationFederatedCredentials, err := g.getApplicationFederatedCredentials(ctx, application)
	if err != nil {
		return nil, err
	}

	result.Application = application
	result.ApplicationFederatedCredentials = applicationFederatedCredentials
	return result, nil
}

func (g *GraphHelper) findServicePrincipalByAppID(ctx context.Context, appID string) (models.ServicePrincipalable, error) {
	appIDFilter := fmt.Sprintf("appId eq '%s'", appID)

	spResponse, err := g.Client().ServicePrincipals().Get(
		ctx,
		&serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
			QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
				Filter: &appIDFilter,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed querying service principals for appId %q: %w", appID, err)
	}

	if spResponse == nil || len(spResponse.GetValue()) == 0 {
		return nil, fmt.Errorf("no service principal found for appId %q", appID)
	}

	return spResponse.GetValue()[0], nil
}

func (g *GraphHelper) findApplicationByAppID(ctx context.Context, appID string) (models.Applicationable, error) {
	appIDFilter := fmt.Sprintf("appId eq '%s'", appID)

	appResponse, err := g.Client().Applications().Get(
		ctx,
		&applications.ApplicationsRequestBuilderGetRequestConfiguration{
			QueryParameters: &applications.ApplicationsRequestBuilderGetQueryParameters{
				Filter: &appIDFilter,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed querying applications for appId %q: %w", appID, err)
	}

	if appResponse == nil || len(appResponse.GetValue()) == 0 {
		return nil, fmt.Errorf("no app registration found for appId %q", appID)
	}

	return appResponse.GetValue()[0], nil
}

func (g *GraphHelper) GetManagedIdentity(ctx context.Context, subscriptionID, resourceGroupName, resourceName string) (*armmsi.Identity, error) {
	managedIdentityClient, err := armmsi.NewUserAssignedIdentitiesClient(subscriptionID, g.credential, nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating managed identity client: %w", err)
	}

	managedIdentityResponse, err := managedIdentityClient.Get(ctx, resourceGroupName, resourceName, nil)
	if err != nil {
		return nil, fmt.Errorf("failed getting managed identity %q in resource group %q: %w", resourceName, resourceGroupName, err)
	}

	return &managedIdentityResponse.Identity, nil
}

func (g *GraphHelper) getManagedIdentityFederatedCredentials(ctx context.Context, subscriptionID, resourceGroupName, resourceName string) ([]*armmsi.FederatedIdentityCredential, error) {
	federatedIdentityCredentialsClient, err := armmsi.NewFederatedIdentityCredentialsClient(subscriptionID, g.credential, nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating federated identity credentials client: %w", err)
	}

	pager := federatedIdentityCredentialsClient.NewListPager(resourceGroupName, resourceName, nil)
	federatedCredentials := make([]*armmsi.FederatedIdentityCredential, 0)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed listing managed identity federated credentials for %q in resource group %q: %w", resourceName, resourceGroupName, err)
		}

		federatedCredentials = append(federatedCredentials, page.Value...)
	}

	return federatedCredentials, nil
}

func (g *GraphHelper) getApplicationFederatedCredentials(ctx context.Context, application models.Applicationable) ([]models.FederatedIdentityCredentialable, error) {
	applicationObjectID := strings.TrimSpace(stringOrEmpty(application.GetId()))
	if applicationObjectID == "" {
		return nil, errors.New("application object id is missing")
	}

	federatedCredentialsResponse, err := g.Client().Applications().ByApplicationId(applicationObjectID).FederatedIdentityCredentials().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed querying app registration federated credentials for application object id %q: %w", applicationObjectID, err)
	}

	if federatedCredentialsResponse == nil {
		return []models.FederatedIdentityCredentialable{}, nil
	}

	return federatedCredentialsResponse.GetValue(), nil
}

func classifyServicePrincipalResourceType(servicePrincipal models.ServicePrincipalable) ServicePrincipalResourceType {

	servicePrincipalType := strings.TrimSpace(strings.ToLower(stringOrEmpty(servicePrincipal.GetServicePrincipalType())))
	if servicePrincipalType == "managedidentity" {
		return ServicePrincipalResourceManagedIdentity
	}

	if servicePrincipalType == "application" {
		return ServicePrincipalResourceAppRegistration
	}

	_, _, _, err := parseManagedIdentityResourceID(servicePrincipal.GetAlternativeNames())
	if err == nil {
		return ServicePrincipalResourceManagedIdentity
	}

	return ServicePrincipalResourceAppRegistration
}

func parseManagedIdentityResourceID(alternativeNames []string) (subscriptionID, resourceGroupName, resourceName string, err error) {
	for _, alternativeName := range alternativeNames {
		pathParts := strings.Split(strings.Trim(alternativeName, "/"), "/")
		for i := 0; i+7 < len(pathParts); i++ {
			if strings.EqualFold(pathParts[i], "subscriptions") &&
				strings.EqualFold(pathParts[i+2], "resourceGroups") &&
				strings.EqualFold(pathParts[i+4], "providers") &&
				strings.EqualFold(pathParts[i+5], "Microsoft.ManagedIdentity") &&
				strings.EqualFold(pathParts[i+6], "userAssignedIdentities") {
				return pathParts[i+1], pathParts[i+3], pathParts[i+7], nil
			}
		}
	}

	return "", "", "", errors.New("no user-assigned managed identity resource id found in service principal alternativeNames")
}

func stringOrEmpty(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}
