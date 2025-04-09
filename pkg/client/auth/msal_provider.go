package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/equinor/radix-cli/pkg/config"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

const (
	RadixCliClientID    = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	AzureTenantID       = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
	AccessTokenCacheKey = "access_token"
	ConfigCacheKey      = "auth"
)

// MSALAuthProvider is an AuthProvider that uses MSAL
type MSALAuthProvider interface {
	Login(ctx context.Context, useInteractiveLogin, useDeviceCode bool, azureClientId, federatedTokenFile, azureClientSecret string) error
	Logout(ctx context.Context) error
	runtime.ClientAuthInfoWriter
}

type msalAuthProvider struct {
	authority   string
	client      *public.Client
	radixConfig *radixconfig.RadixConfig
	config      LoginConfig
}

type LoginConfig struct {
	UseInteractiveLogin bool
	UseDeviceCode       bool
	AzureClientId       string
	FederatedTokenFile  string
}

// NewMSALAuthProvider creates a new MSALAuthProvider
func NewMSALAuthProvider(radixConfig *radixconfig.RadixConfig) (MSALAuthProvider, error) {
	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", AzureTenantID)
	client, err := newPublicClient(radixConfig, RadixCliClientID, authority)

	if err != nil {
		return nil, err
	}
	config, _ := config.GetCache[LoginConfig](ConfigCacheKey)
	return &msalAuthProvider{
		client:      client,
		authority:   authority,
		radixConfig: radixConfig,
		config:      config,
	}, nil
}

// Login allows the plugin to initialize its configuration. It must not
// require direct user interaction.
func (provider *msalAuthProvider) Login(ctx context.Context, useInteractiveLogin, useDeviceCode bool, azureClientId, federatedTokenFile, azureClientSecret string) error {
	provider.config = LoginConfig{
		UseInteractiveLogin: useInteractiveLogin,
		UseDeviceCode:       useDeviceCode,
		AzureClientId:       azureClientId,
		FederatedTokenFile:  federatedTokenFile,
	}
	config.SetCache(ConfigCacheKey, provider.config, 365*24*time.Hour)

	switch {
	case provider.config.UseInteractiveLogin:
		_, err := provider.loginInteractive(ctx)
		return err
	case provider.config.UseDeviceCode:
		_, err := provider.loginDeviceCode(ctx)
		return err
	case provider.config.AzureClientId != "" && provider.config.FederatedTokenFile != "":
		_, err := provider.loginFederatedCredentials(ctx)
		return err
	case provider.config.AzureClientId != "" && azureClientSecret != "":
		_, err := provider.loginClientSecret(ctx, azureClientSecret)
		return err
	}

	return fmt.Errorf("no valild authentication combinations found")
}

// Logout removes all cached accounts with tokens
func (provider *msalAuthProvider) Logout(ctx context.Context) error {
	config.ClearCache(ConfigCacheKey)
	config.ClearCache(AccessTokenCacheKey)

	accounts, err := provider.client.Accounts(ctx)
	if err != nil {
		return err
	}
	for _, account := range accounts {
		if err := provider.client.RemoveAccount(ctx, account); err != nil {
			return err
		}
	}
	return nil
}

func (provider *msalAuthProvider) AuthenticateRequest(r runtime.ClientRequest, _ strfmt.Registry) error {
	var token string
	var err error

	switch {
	case provider.config.UseInteractiveLogin:
		token, err = provider.GetMsalToken(context.Background())
	case provider.config.UseDeviceCode:
		token, err = provider.GetMsalToken(context.Background())
	case provider.config.AzureClientId != "" && provider.config.FederatedTokenFile != "":
		token, err = provider.loginFederatedCredentials(context.Background())
	case provider.config.AzureClientId != "":
		token, err = provider.loginClientSecret(context.Background(), "")
	}
	if err != nil {
		return err
	}

	return r.SetHeaderParam(runtime.HeaderAuthorization, "Bearer "+token)
}

// GetMsalToken returns a valid token for the given scopes
func (provider *msalAuthProvider) GetMsalToken(ctx context.Context) (string, error) {
	accounts, err := provider.client.Accounts(ctx)

	if err != nil {
		return "", err
	}
	if len(accounts) > 0 {
		// found a cached account, now see if an applicable token has been cached
		// NOTE: this API conflates error states, i.e. err is non-nil if an applicable token isn't
		//       cached or if something goes wrong (making the HTTP request, unmarshalling, etc).
		authResult, err := provider.client.AcquireTokenSilent(ctx, getScopes(), public.WithSilentAccount(accounts[0]))
		if err == nil {
			return authResult.AccessToken, nil
		}
	}

	// either there was no cached account/token or the call to AcquireTokenSilent() failed
	// make a new request to AAD
	return provider.loginInteractive(ctx)
}

func (provider *msalAuthProvider) loginInteractive(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	fmt.Printf("A web browser has been opened at %s/oauth2/v2.0/authorize. Please continue the login in the web browser.\n", provider.authority)
	result, err := provider.client.AcquireTokenInteractive(ctx, getScopes())
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

func (provider *msalAuthProvider) loginDeviceCode(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	devCode, err := provider.client.AcquireTokenByDeviceCode(ctx, getScopes())
	if err != nil {
		return "", fmt.Errorf("got error while waiting for user to input the device code: %s", err)
	}

	fmt.Println(devCode.Result.Message) // show authentication link with device code

	result, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

func (provider *msalAuthProvider) loginFederatedCredentials(ctx context.Context) (string, error) {
	if token, ok := config.GetCache[string](AccessTokenCacheKey); ok {
		return token, nil
	}

	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	cred, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
		//	Cache:         c,
		ClientID:      provider.config.AzureClientId,
		TenantID:      AzureTenantID,
		TokenFilePath: provider.config.FederatedTokenFile,
	})

	authResult, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes:   getScopes(),
		TenantID: AzureTenantID,
	})
	if err != nil {
		return "", err
	}

	config.SetCache(AccessTokenCacheKey, authResult.Token, authResult.ExpiresOn.Sub(time.Now()))
	return authResult.Token, nil
}

func (provider *msalAuthProvider) loginClientSecret(ctx context.Context, azureClientSecret string) (string, error) {
	if token, ok := config.GetCache[string](AccessTokenCacheKey); ok {
		return token, nil
	}

	if azureClientSecret == "" {
		return "", errors.New("please login again")
	}

	cred, err := confidential.NewCredFromSecret(azureClientSecret)
	if err != nil {
		return "", err
	}

	cache := NewTokenCache(provider.radixConfig)
	confidentialClient, err := confidential.New(provider.authority, provider.config.AzureClientId, cred, confidential.WithCache(cache))
	if err != nil {
		return "", err
	}

	authResults, err := confidentialClient.AcquireTokenByCredential(ctx, getScopes(), confidential.WithTenantID(AzureTenantID))
	if err != nil {
		return "", err
	}

	config.SetCache(AccessTokenCacheKey, authResults.AccessToken, authResults.ExpiresOn.Sub(time.Now()))
	return authResults.AccessToken, nil
}

func getScopes() []string {
	return []string{"6dae42f8-4368-4678-94ff-3960e28e3630/.default"}
}
