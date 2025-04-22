package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/equinor/radix-cli/pkg/cache"
)

type AzureFederatedCredentials struct {
	cache cache.Cache
}

var _ GetAccessTokener = &AzureFederatedCredentials{}

func NewAzureFederatedCredentials(cache cache.Cache) *AzureFederatedCredentials {

	return &AzureFederatedCredentials{
		cache: cache,
	}
}

func (p *AzureFederatedCredentials) Authenticate(ctx context.Context, azureClientId, federatedTokenFile string) (string, error) {
	if token, ok := p.cache.GetItem(AccessTokenCacheKey); ok {
		return token, nil
	}

	if federatedTokenFile == "" || azureClientId == "" {
		return "", errors.New("please login again")
	}
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	cred, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
		ClientID:      azureClientId,
		TenantID:      AzureTenantID,
		TokenFilePath: federatedTokenFile,
	})
	if err != nil {
		return "", err
	}

	authResult, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes:   getScopes(),
		TenantID: AzureTenantID,
	})
	if err != nil {
		return "", err
	}

	p.cache.SetItem(azureClientIdCacheKey, azureClientId, time.Until(authResult.ExpiresOn))
	p.cache.SetItem(federatedTokenFileCacheKey, federatedTokenFile, time.Until(authResult.ExpiresOn))
	p.cache.SetItem(AccessTokenCacheKey, authResult.Token, time.Until(authResult.ExpiresOn))
	return authResult.Token, nil
}

func (p *AzureFederatedCredentials) GetAccessToken(ctx context.Context) (string, error) {
	if token, ok := p.cache.GetItem(AccessTokenCacheKey); ok {
		return token, nil
	}

	azureClientId, _ := p.cache.GetItem(azureClientIdCacheKey)
	federatedTokenFile, _ := p.cache.GetItem(federatedTokenFileCacheKey)
	return p.Authenticate(ctx, azureClientId, federatedTokenFile)
}
