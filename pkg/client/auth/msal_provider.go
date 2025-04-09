package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
	RadixCliClientID           = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	AzureTenantID              = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
	AccessTokenCacheKey        = "access_token"
	ConfigCacheKey             = "auth"
	azureADAudience            = "api://AzureADTokenExchange"
	actionsIDTokenRequestToken = "ACTIONS_ID_TOKEN_REQUEST_TOKEN"
	actionsIDTokenRequestURL   = "ACTIONS_ID_TOKEN_REQUEST_URL"
)

// MSALAuthProvider is an AuthProvider that uses MSAL
type MSALAuthProvider interface {
	Login(ctx context.Context, useInteractiveLogin, useDeviceCode, useGithubCredentials bool, azureClientId, federatedTokenFile, azureClientSecret string) error
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
	UseInteractiveLogin  bool
	UseGithubCredentials bool
	UseDeviceCode        bool
	AzureClientId        string
	FederatedTokenFile   string
}

type githubTokenResponse struct {
	Value string `json:"value"`
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
func (provider *msalAuthProvider) Login(ctx context.Context, useInteractiveLogin, useDeviceCode, useGithubCredentials bool, azureClientId, federatedTokenFile, azureClientSecret string) error {
	provider.config = LoginConfig{
		UseInteractiveLogin:  useInteractiveLogin,
		UseDeviceCode:        useDeviceCode,
		AzureClientId:        azureClientId,
		FederatedTokenFile:   federatedTokenFile,
		UseGithubCredentials: useGithubCredentials,
	}
	config.SetCache(ConfigCacheKey, provider.config, 365*24*time.Hour)

	switch {
	case provider.config.UseInteractiveLogin:
		_, err := provider.loginInteractive(ctx)
		return err
	case provider.config.UseDeviceCode:
		_, err := provider.loginDeviceCode(ctx)
		return err
	case provider.config.UseGithubCredentials:
		if provider.config.AzureClientId == "" {
			return errors.New("missing Azure Client ID")
		}
		_, err := provider.loginGithubFederatedCredentials(ctx)
		return err
	case provider.config.FederatedTokenFile != "":
		if provider.config.AzureClientId == "" {
			return errors.New("missing Azure Client ID")
		}
		_, err := provider.loginFederatedCredentials(ctx)
		return err
	case azureClientSecret != "":
		if provider.config.AzureClientId == "" {
			return errors.New("missing Azure Client ID")
		}
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
	case provider.config.UseGithubCredentials:
		token, err = provider.loginGithubFederatedCredentials(context.Background())
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

func (provider *msalAuthProvider) loginGithubFederatedCredentials(ctx context.Context) (string, error) {
	// Mostly copied from kubelogin\pkg\internal\token\githubactionscredential.go:newGithubActionsCredential()
	if token, ok := config.GetCache[string](AccessTokenCacheKey); ok {
		return token, nil
	}

	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	cred := confidential.NewCredFromAssertionCallback(func(ctx context.Context, _ confidential.AssertionRequestOptions) (string, error) {
		return getGithubFedCred(ctx)
	})

	client, err := confidential.New(provider.authority, provider.config.AzureClientId, cred)
	if err != nil {
		return "", fmt.Errorf("failed to create github actions credential: %w", err)
	}

	authResult, err := client.AcquireTokenByCredential(ctx, getScopes())
	if err != nil {
		return "", err
	}

	config.SetCache(AccessTokenCacheKey, authResult.AccessToken, authResult.ExpiresOn.Sub(time.Now()))
	return authResult.AccessToken, nil
}

func getGithubFedCred(ctx context.Context) (string, error) {
	// All code copied from kubelogins kubelogin\pkg\internal\token\githubactionscredential.go:getGithubToken()
	reqToken := os.Getenv(actionsIDTokenRequestToken)
	reqURL := os.Getenv(actionsIDTokenRequestURL)

	if reqToken == "" || reqURL == "" {
		return "", errors.New("ACTIONS_ID_TOKEN_REQUEST_TOKEN or ACTIONS_ID_TOKEN_REQUEST_URL is not set")
	}

	u, err := url.Parse(reqURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse ACTIONS_ID_TOKEN_REQUEST_URL: %w", err)
	}
	q := u.Query()
	q.Set("audience", azureADAudience)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	// reference:
	// https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", reqToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; api-version=2.0")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body string
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			body = err.Error()
		} else {
			body = string(b)
		}

		return "", fmt.Errorf("github actions ID token request failed with status code: %d, response body: %s", resp.StatusCode, body)
	}

	var tokenResp githubTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	if tokenResp.Value == "" {
		return "", errors.New("github actions ID token is empty")
	}

	return tokenResp.Value, nil
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
