package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/equinor/radix-cli/pkg/cache"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

const (
	RadixCliClientID           = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	AzureTenantID              = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
	AccessTokenCacheKey        = "access_token"
	azureClientIdCacheKey      = "azure_client_id"
	federatedTokenFileCacheKey = "federated_token_file"
	ConfigCacheKey             = "auth"
	azureADAudience            = "api://AzureADTokenExchange"
	actionsIDTokenRequestToken = "ACTIONS_ID_TOKEN_REQUEST_TOKEN"
	actionsIDTokenRequestURL   = "ACTIONS_ID_TOKEN_REQUEST_URL"
	authProviderCacheKey       = "auth_provider"
	authProviderTypeCacheKey   = "auth_provider_type"

	// provider constants
	providerMsalInteractive           = "msal_interactive"
	providerMsalDevicecode            = "msal_devicecode"
	providerAzureClientSecret         = "azure_client_secret"
	providerAzureFederatedCredentials = "azure_federated_credentials"
	providerAzureGithub               = "azure_github"
)

var (
	errProviderNotSet  = errors.New("auth provider not set, please login")
	errProviderUnknown = errors.New("auth provider is unknown, please login")
)

type GetAccessTokener interface {
	// GetAccessToken returns a valid token
	GetAccessToken(context.Context) (string, error)
}

// Provider is an Provider that uses MSAL
type Provider interface {
	Login(ctx context.Context, useInteractiveLogin, useDeviceCode, useGithubCredentials bool, azureClientId, federatedTokenFile, azureClientSecret string) error
	Logout() error
	runtime.ClientAuthInfoWriter
}

type auth struct {
	authority string
	provider  GetAccessTokener
	cacheFn   func(namespace string) cache.Cache
	cache     cache.Cache
}

type githubTokenResponse struct {
	Value string `json:"value"`
}

// NewMSALAuthProvider creates a new Provider
func NewMSALAuthProvider(radixConfig *config.RadixConfig) (Provider, error) {
	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", AzureTenantID)
	authCacheFilename := fmt.Sprintf("%s/auth.%s.json", config.RadixConfigDir, radixConfig.CustomConfig.Context)
	globalCache := cache.New(authCacheFilename, "global")

	provider, err := LoadProviderFromCache(globalCache, authCacheFilename, authority)
	if err != nil && !errors.Is(err, errProviderNotSet) {
		return nil, err
	}

	return &auth{
		authority: authority,
		provider:  provider,
		cache:     globalCache,
		cacheFn:   func(namespace string) cache.Cache { return cache.New(authCacheFilename, namespace) },
	}, nil
}

// Login allows the plugin to initialize its configuration. It must not
// require direct user interaction.
func (a *auth) Login(ctx context.Context, useInteractiveLogin, useDeviceCode, useGithubCredentials bool, azureClientId, federatedTokenFile, azureClientSecret string) error {
	switch {
	case useInteractiveLogin:
		provider, err := NewMsalInteractive(NewMsalTokenCache(a.cacheFn(providerMsalInteractive), "msal"), a.authority)
		if err != nil {
			return err
		}
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerMsalInteractive, 365*24*time.Hour)

		_, err = provider.Authenticate(ctx)
		return err

	case useDeviceCode:
		provider, err := NewMsalDeviceCode(NewMsalTokenCache(a.cacheFn(providerMsalDevicecode), "msal"), a.authority)
		if err != nil {
			return err
		}
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerMsalDevicecode, 365*24*time.Hour)

		_, err = provider.Authenticate(ctx)
		return err

	case useGithubCredentials:
		provider := NewAzureGithub(a.cacheFn(providerAzureGithub), a.authority)
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerAzureGithub, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, azureClientSecret)
		return err

	case federatedTokenFile != "":
		provider := NewAzureFederatedCredentials(a.cacheFn(providerAzureFederatedCredentials))
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerAzureFederatedCredentials, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, azureClientId, federatedTokenFile)
		return err

	case azureClientSecret != "":
		provider := NewAzureClientSecret(a.cacheFn(providerAzureClientSecret), a.authority)
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerAzureClientSecret, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, azureClientId, azureClientSecret)
		return err

	}

	return errors.New("invalid auth arguments")
}

func LoadProviderFromCache(globalCache cache.Cache, authCacheFilename, authority string) (GetAccessTokener, error) {

	providerType, ok := globalCache.GetItem(authProviderTypeCacheKey)
	if !ok || providerType == "" {
		return nil, errProviderNotSet
	}

	switch providerType {
	case providerMsalInteractive:
		msalCache := cache.New(authCacheFilename, providerMsalInteractive)
		return NewMsalInteractive(NewMsalTokenCache(msalCache, "msal"), authority)

	case providerMsalDevicecode:
		msalCache := cache.New(authCacheFilename, providerMsalDevicecode)
		return NewMsalDeviceCode(NewMsalTokenCache(msalCache, "msal"), authority)

	case providerAzureGithub:
		localCache := cache.New(authCacheFilename, providerAzureGithub)
		return NewAzureGithub(localCache, authority), nil

	case providerAzureClientSecret:
		localCache := cache.New(authCacheFilename, providerAzureClientSecret)
		return NewAzureClientSecret(localCache, authority), nil

	case providerAzureFederatedCredentials:
		localCache := cache.New(authCacheFilename, providerAzureFederatedCredentials)
		return NewAzureFederatedCredentials(localCache), nil
	}

	return nil, errProviderUnknown
}

// Logout removes all cached accounts with tokens
func (a *auth) Logout() error {
	config.ClearCache(ConfigCacheKey)
	config.ClearCache(AccessTokenCacheKey)
	config.ClearCache(authProviderTypeCacheKey)
	config.ClearCache(authProviderCacheKey)

	authFilesGlob := fmt.Sprintf("%s/auth.*.json", config.RadixConfigDir)
	files, err := filepath.Glob(authFilesGlob)
	if err != nil {
		log.Printf("Error fetching auth files (%s): %s", authFilesGlob, err)

	}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			log.Printf("Error removing file %s: %s", file, err)
		}
	}

	// Legacy: Logout of previus MSAL state
	if rc, err := config.GetRadixConfig(); err == nil {
		rc.MSAL = ""
		err := config.Save(rc)
		if err != nil {
			log.Printf("Error deleting MSAL auth from file %s: %s", config.RadixConfigFileFullName, err)
		}
	}

	return nil
}

func (a *auth) AuthenticateRequest(r runtime.ClientRequest, _ strfmt.Registry) error {
	if a.provider == nil {
		return errProviderNotSet
	}

	token, err := a.provider.GetAccessToken(context.Background())
	if err != nil {
		return err
	}

	return r.SetHeaderParam(runtime.HeaderAuthorization, "Bearer "+token)
}

func getScopes() []string {
	return []string{"6dae42f8-4368-4678-94ff-3960e28e3630/.default"}
}
