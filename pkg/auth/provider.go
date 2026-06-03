package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/equinor/radix-cli/pkg/auth/cache"
	"github.com/equinor/radix-cli/pkg/config"
)

const (
	radixCliClientID = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	azureTenantID    = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"

	azureClientIdCacheKey = "azure_client_id"

	authProviderTypeCacheKey = "auth_provider_type"

	// provider constants
	providerMsalInteractive           = "msal_interactive"
	providerMsalDevicecode            = "msal_devicecode"
	providerAzureClientSecret         = "azure_client_secret"
	providerAzureFederatedCredentials = "azure_federated_credentials"
	providerAzureGithub               = "azure_github"

	authFileFormat = "%s/auth.json"
)

var (
	ErrProviderNotSet  = errors.New("auth provider not set, please login")
	ErrProviderUnknown = errors.New("auth provider is unknown, please login")

	defaultLoginScope = []string{"6dae42f8-4368-4678-94ff-3960e28e3630/.default"}
)

var _ Provider = &Auth{}

type AccessToken struct {
	// Token is the access token
	Token string
	// ExpiresOn indicates when the token expires
	ExpiresOn time.Time
}

type GetAccessTokener interface {
	// GetAccessToken returns a valid token
	GetAccessToken(ctx context.Context, scopes []string) (AccessToken, error)
}

// Provider is an Provider that uses MSAL
type Provider interface {
	GetAccessTokener
	Login(ctx context.Context, useInteractiveLogin, useDeviceCode, useGithubCredentials bool, azureClientId, federatedTokenFile, azureClientSecret string) error
	Logout() error
}
type Auth struct {
	authority string
	provider  GetAccessTokener
	cacheFn   func(namespace string) cache.Cache
	cache     cache.Cache
}

type githubTokenResponse struct {
	Value string `json:"value"`
}

// New creates a new Provider
func New() (*Auth, error) {
	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", azureTenantID)
	authCacheFilename := fmt.Sprintf(authFileFormat, config.RadixConfigDir)
	globalCache := cache.New(authCacheFilename, "global")

	provider, err := loadProviderFromCache(globalCache, authCacheFilename, authority)
	if err != nil && !errors.Is(err, ErrProviderNotSet) {
		return nil, err
	}

	return &Auth{
		authority: authority,
		provider:  provider,
		cache:     globalCache,
		cacheFn:   func(namespace string) cache.Cache { return cache.New(authCacheFilename, namespace) },
	}, nil
}

// Login allows the plugin to initialize its configuration. It must not
// require direct user interaction.
func (a *Auth) Login(ctx context.Context, useInteractiveLogin, useDeviceCode, useGithubCredentials bool, azureClientId, federatedTokenFile, azureClientSecret string) error {
	switch {
	case useInteractiveLogin:
		provider := NewMsalInteractive(NewMsalTokenCache(a.cacheFn(providerMsalInteractive), "msal"), a.authority)
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerMsalInteractive, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, defaultLoginScope)
		return err

	case useDeviceCode:
		provider := NewMsalDeviceCode(NewMsalTokenCache(a.cacheFn(providerMsalDevicecode), "msal"), a.authority)
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerMsalDevicecode, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, defaultLoginScope)
		return err

	case useGithubCredentials:
		provider := NewAzureGithub(a.cacheFn(providerAzureGithub), a.authority)
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerAzureGithub, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, azureClientId, defaultLoginScope)
		return err

	case federatedTokenFile != "":
		provider := NewAzureFederatedCredentials(a.cacheFn(providerAzureFederatedCredentials))
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerAzureFederatedCredentials, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, azureClientId, federatedTokenFile, defaultLoginScope)
		return err

	case azureClientSecret != "":
		provider := NewAzureClientSecret(a.cacheFn(providerAzureClientSecret), a.authority)
		a.provider = provider
		a.cache.SetItem(authProviderTypeCacheKey, providerAzureClientSecret, 365*24*time.Hour)

		_, err := provider.Authenticate(ctx, azureClientId, azureClientSecret, defaultLoginScope)
		return err

	}

	return errors.New("invalid auth arguments")
}

// Logout removes all cached accounts with tokens
func (a *Auth) Logout() error {
	authFilesGlob := fmt.Sprintf(authFileFormat, config.RadixConfigDir)
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

func (a *Auth) GetAccessToken(ctx context.Context, scopes []string) (AccessToken, error) {
	if a.provider == nil {
		return AccessToken{}, ErrProviderNotSet
	}

	return a.provider.GetAccessToken(ctx, scopes)
}

func loadProviderFromCache(globalCache cache.Cache, authCacheFilename, authority string) (GetAccessTokener, error) {

	providerType, ok := globalCache.GetItem(authProviderTypeCacheKey)
	if !ok || providerType.Content == "" {
		return nil, ErrProviderNotSet
	}

	switch providerType.Content {
	case providerMsalInteractive:
		msalCache := cache.New(authCacheFilename, providerMsalInteractive)
		return NewMsalInteractive(NewMsalTokenCache(msalCache, "msal"), authority), nil

	case providerMsalDevicecode:
		msalCache := cache.New(authCacheFilename, providerMsalDevicecode)
		return NewMsalDeviceCode(NewMsalTokenCache(msalCache, "msal"), authority), nil

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

	return nil, ErrProviderUnknown
}
