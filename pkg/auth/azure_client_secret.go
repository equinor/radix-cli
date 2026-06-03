package auth

import (
	"context"
	"errors"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"github.com/equinor/radix-cli/pkg/auth/cache"
)

const (
	azureClientSecretCacheKey = "azure_client_secret"
)

type AzureClientSecret struct {
	Authority string
	cache     cache.Cache
}

var _ GetAccessTokener = &AzureClientSecret{}

func NewAzureClientSecret(cache cache.Cache, authority string) *AzureClientSecret {
	return &AzureClientSecret{
		Authority: authority,
		cache:     cache,
	}
}

func (p *AzureClientSecret) Authenticate(ctx context.Context, azureClientId, azureClientSecret string, scopes []string) (AccessToken, error) {
	if azureClientSecret == "" || azureClientId == "" {
		return AccessToken{}, errors.New("please login again")
	}

	cred, err := confidential.NewCredFromSecret(azureClientSecret)
	if err != nil {
		return AccessToken{}, err
	}

	confidentialClient, err := confidential.New(p.Authority, azureClientId, cred)
	if err != nil {
		return AccessToken{}, err
	}

	authResult, err := confidentialClient.AcquireTokenByCredential(ctx, scopes, confidential.WithTenantID(azureTenantID))
	if err != nil {
		return AccessToken{}, err
	}

	p.cache.SetItem(azureClientIdCacheKey, azureClientId, 365*24*time.Hour)
	p.cache.SetItem(azureClientSecretCacheKey, azureClientSecret, 365*24*time.Hour)
	p.cache.SetItem(accessTokenCacheKey(scopes), authResult.AccessToken, time.Until(authResult.ExpiresOn))
	return AccessToken{Token: authResult.AccessToken, ExpiresOn: authResult.ExpiresOn}, nil
}

func (p *AzureClientSecret) GetAccessToken(ctx context.Context, scopes []string) (AccessToken, error) {
	if token, ok := p.cache.GetItem(accessTokenCacheKey(scopes)); ok {
		return AccessToken{Token: token.Content, ExpiresOn: token.ExpiresAt}, nil
	}

	azureClientId, _ := p.cache.GetItem(azureClientIdCacheKey)
	azureClientSecret, _ := p.cache.GetItem(azureClientSecretCacheKey)
	return p.Authenticate(ctx, azureClientId.Content, azureClientSecret.Content, scopes)
}
