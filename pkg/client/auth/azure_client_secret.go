package auth

import (
	"context"
	"errors"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	cache2 "github.com/equinor/radix-cli/pkg/cache"
)

type AzureClientSecret struct {
	Authority string
	cache     cache2.Cache
}

var _ GetAccessTokener = &AzureClientSecret{}

func NewAzureClientSecret(cache cache2.Cache, authority string) *AzureClientSecret {
	return &AzureClientSecret{
		Authority: authority,
		cache:     cache,
	}
}

func (p *AzureClientSecret) Authenticate(ctx context.Context, azureClientId, azureClientSecret string) (string, error) {
	if azureClientSecret == "" || azureClientId == "" {
		return "", errors.New("please login again")
	}

	cred, err := confidential.NewCredFromSecret(azureClientSecret)
	if err != nil {
		return "", err
	}

	confidentialClient, err := confidential.New(p.Authority, azureClientId, cred)
	if err != nil {
		return "", err
	}

	authResult, err := confidentialClient.AcquireTokenByCredential(ctx, getScopes(), confidential.WithTenantID(AzureTenantID))
	if err != nil {
		return "", err
	}

	p.cache.SetItem(azureClientIdCacheKey, azureClientId, time.Until(authResult.ExpiresOn))
	p.cache.SetItem(AccessTokenCacheKey, authResult.AccessToken, time.Until(authResult.ExpiresOn))
	return authResult.AccessToken, nil
}

func (p *AzureClientSecret) GetAccessToken(_ context.Context) (string, error) {
	if token, ok := p.cache.GetItem(AccessTokenCacheKey); ok {
		return token, nil
	}

	return "", errors.New("please login again")
}
