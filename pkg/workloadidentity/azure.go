package workloadidentity

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	"github.com/equinor/radix-common/utils/slice"
	kiotaauth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
)

type ServicePrincipalType string

const (
	ManagedIdentity ServicePrincipalType = "managed-identity"
	AppRegistration ServicePrincipalType = "app-registration"
)

var (
	graphScopes = []string{"https://graph.microsoft.com/.default"} //"Application.ReadUpdate.All"
)

type ServicePrincipal struct {
	ClientID             string                `json:"clientID"`
	Type                 ServicePrincipalType  `json:"type"`
	DisplayName          string                `json:"displayName"`
	SubscriptionID       string                `json:"subscriptionID,omitempty"`
	ResourceGroup        string                `json:"resourceGroup,omitempty"`
	FederatedCredentials []FederatedCredential `json:"existingFederatedCredentials"`
}

type FederatedCredential struct {
	Name      string   `json:"name"`
	Issuer    string   `json:"issuer"`
	Subject   string   `json:"subject"`
	Audiences []string `json:"audiences"`
}

type AzureServicePrincipalHelper struct {
	credential azcore.TokenCredential
	// azCreds         *azidentity.AzureCLICredential
	graphClient *msgraphsdk.GraphServiceClient
}

func NewAzureServicePrincipalHelper(credential azcore.TokenCredential) (*AzureServicePrincipalHelper, error) {
	g := &AzureServicePrincipalHelper{
		credential: credential,
	}

	if err := g.initGraphClient(); err != nil {
		return nil, err
	}

	return g, nil
}

// GetServicePrincipalDetails resolves a service principal appId and fetches
// the backing resource details as either a managed identity or app registration.
func (g *AzureServicePrincipalHelper) GetServicePrincipal(ctx context.Context, clientId string) (*ServicePrincipal, error) {
	servicePrincipal, servicePrincipalType, err := g.getGraphServicePrincipalAndType(ctx, clientId)
	if err != nil {
		return nil, err
	}

	switch servicePrincipalType {
	case ManagedIdentity:
		return g.getManagedIdentityServicePrincipal(ctx, servicePrincipal)
	case AppRegistration:
		return g.getAppRegistrationServicePrincipal(ctx, servicePrincipal)
	}

	return nil, fmt.Errorf("unhandled service principal type %q", servicePrincipalType)
}

func (g *AzureServicePrincipalHelper) getGraphServicePrincipalAndType(ctx context.Context, clientId string) (models.ServicePrincipalable, ServicePrincipalType, error) {
	clientId = strings.TrimSpace(clientId)
	if clientId == "" {
		return nil, "", errors.New("clientId is required")
	}

	servicePrincipal, err := g.queryServicePrincipalByAppId(ctx, clientId)
	if err != nil {
		return nil, "", err
	}
	servicePrincipalType := classifyServicePrincipalResourceType(servicePrincipal)

	return servicePrincipal, servicePrincipalType, nil
}

func (g *AzureServicePrincipalHelper) getAppRegistrationServicePrincipal(ctx context.Context, sp models.ServicePrincipalable) (*ServicePrincipal, error) {
	application, err := g.getApplicationByAppID(ctx, stringOrEmpty(sp.GetAppId()))
	if err != nil {
		return nil, err
	}

	applicationFederatedCredentials, err := g.getApplicationFederatedCredentials(ctx, application)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipal{
		ClientID:             stringOrEmpty(sp.GetAppId()),
		Type:                 AppRegistration,
		DisplayName:          stringOrEmpty(sp.GetDisplayName()),
		FederatedCredentials: slice.Map(applicationFederatedCredentials, mapAppRegistrationManagedIdentity),
	}, nil
}

func (g *AzureServicePrincipalHelper) getManagedIdentityServicePrincipal(ctx context.Context, sp models.ServicePrincipalable) (*ServicePrincipal, error) {
	subscriptionID, resourceGroupName, resourceName, err := parseManagedIdentityResourceID(sp.GetAlternativeNames())
	if err != nil {
		return nil, err
	}

	managedIdentityFederatedCredentials, err := g.getManagedIdentityFederatedCredentials(ctx, subscriptionID, resourceGroupName, resourceName)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipal{
		ClientID:             stringOrEmpty(sp.GetAppId()),
		SubscriptionID:       subscriptionID,
		ResourceGroup:        resourceGroupName,
		Type:                 ManagedIdentity,
		DisplayName:          stringOrEmpty(sp.GetDisplayName()),
		FederatedCredentials: slice.Map(managedIdentityFederatedCredentials, mapManagedIdentityFederatedCredential),
	}, nil
}

func (g *AzureServicePrincipalHelper) initGraphClient() error {
	// Create an auth provider using the credential
	authProvider, err := kiotaauth.NewAzureIdentityAuthenticationProviderWithScopes(g.credential, graphScopes)
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
	g.graphClient = client

	return nil
}

func (g *AzureServicePrincipalHelper) queryServicePrincipalByAppId(ctx context.Context, appId string) (models.ServicePrincipalable, error) {
	appIDFilter := fmt.Sprintf("appId eq '%s'", appId)

	spResponse, err := g.graphClient.ServicePrincipals().Get(
		ctx,
		&serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
			QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
				Filter: &appIDFilter,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed querying service principals for appId %q: %w", appId, err)
	}

	if spResponse == nil || len(spResponse.GetValue()) == 0 {
		return nil, fmt.Errorf("no service principal found for appId %q", appId)
	}

	return spResponse.GetValue()[0], nil
}

func (g *AzureServicePrincipalHelper) getApplicationByAppID(ctx context.Context, appID string) (models.Applicationable, error) {
	appIDFilter := fmt.Sprintf("appId eq '%s'", appID)

	appResponse, err := g.graphClient.Applications().Get(
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

func (g *AzureServicePrincipalHelper) getManagedIdentityFederatedCredentials(ctx context.Context, subscriptionID, resourceGroupName, resourceName string) ([]*armmsi.FederatedIdentityCredential, error) {
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

func (g *AzureServicePrincipalHelper) getApplicationFederatedCredentials(ctx context.Context, application models.Applicationable) ([]models.FederatedIdentityCredentialable, error) {
	applicationObjectID := strings.TrimSpace(stringOrEmpty(application.GetId()))
	if applicationObjectID == "" {
		return nil, errors.New("application object id is missing")
	}

	federatedCredentialsResponse, err := g.graphClient.Applications().ByApplicationId(applicationObjectID).FederatedIdentityCredentials().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed querying app registration federated credentials for application object id %q: %w", applicationObjectID, err)
	}

	if federatedCredentialsResponse == nil {
		return []models.FederatedIdentityCredentialable{}, nil
	}

	return federatedCredentialsResponse.GetValue(), nil
}

func mapAppRegistrationManagedIdentity(appFedCred models.FederatedIdentityCredentialable) FederatedCredential {
	return FederatedCredential{
		Name:      stringOrEmpty(appFedCred.GetName()),
		Issuer:    stringOrEmpty(appFedCred.GetIssuer()),
		Subject:   stringOrEmpty(appFedCred.GetSubject()),
		Audiences: appFedCred.GetAudiences(),
	}
}

func mapManagedIdentityFederatedCredential(miFedCred *armmsi.FederatedIdentityCredential) FederatedCredential {
	if miFedCred == nil {
		return FederatedCredential{}
	}

	return FederatedCredential{
		Name:      stringOrEmpty(miFedCred.Name),
		Issuer:    stringOrEmpty(miFedCred.Properties.Issuer),
		Subject:   stringOrEmpty(miFedCred.Properties.Subject),
		Audiences: slice.Map(miFedCred.Properties.Audiences, stringOrEmpty),
	}
}

func classifyServicePrincipalResourceType(servicePrincipal models.ServicePrincipalable) ServicePrincipalType {
	servicePrincipalType := strings.TrimSpace(strings.ToLower(stringOrEmpty(servicePrincipal.GetServicePrincipalType())))

	switch servicePrincipalType {
	case "managedidentity":
		return ManagedIdentity
	case "application":
		return AppRegistration
	}

	return ""
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

	return "", "", "", errors.New("no managed identity resource id found in service principal alternativeNames")
}

func stringOrEmpty(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}
