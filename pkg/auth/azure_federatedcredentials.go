package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/equinor/radix-cli/pkg/auth/cache"
)

const (
	federatedTokenFileCacheKey = "federated_token_file"
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

func (p *AzureFederatedCredentials) Authenticate(ctx context.Context, azureClientId, federatedTokenFile string, scopes []string) (AccessToken, error) {
	if federatedTokenFile == "" || azureClientId == "" {
		return AccessToken{}, errors.New("please login again")
	}
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	cred, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
		ClientID:      azureClientId,
		TenantID:      azureTenantID,
		TokenFilePath: federatedTokenFile,
	})
	if err != nil {
		return AccessToken{}, err
	}

	authResult, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes:   scopes,
		TenantID: azureTenantID,
	})
	if err != nil {
		return AccessToken{}, err
	}

	p.cache.SetItem(azureClientIdCacheKey, azureClientId, 365*24*time.Hour)
	p.cache.SetItem(federatedTokenFileCacheKey, federatedTokenFile, 365*24*time.Hour)
	p.cache.SetItem(accessTokenCacheKey(scopes), authResult.Token, time.Until(authResult.ExpiresOn))
	return AccessToken{Token: authResult.Token, ExpiresOn: authResult.ExpiresOn}, nil
}

func (p *AzureFederatedCredentials) GetAccessToken(ctx context.Context, scopes []string) (AccessToken, error) {
	if token, ok := p.cache.GetItem(accessTokenCacheKey(scopes)); ok {
		return AccessToken{Token: token.Content, ExpiresOn: token.ExpiresAt}, nil
	}

	azureClientId, _ := p.cache.GetItem(azureClientIdCacheKey)
	federatedTokenFile, _ := p.cache.GetItem(federatedTokenFileCacheKey)
	return p.Authenticate(ctx, azureClientId.Content, federatedTokenFile.Content, scopes)
}
