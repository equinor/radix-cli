package adapter

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/equinor/radix-cli/pkg/auth"
)

func NewAzureTokenCredentialAdapter(auth *auth.Auth) *AzureTokenCredentialAdapter {
	return &AzureTokenCredentialAdapter{
		auth: auth,
	}
}

type AzureTokenCredentialAdapter struct {
	auth *auth.Auth
}

func (a *AzureTokenCredentialAdapter) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	if a.auth == nil {
		return azcore.AccessToken{}, fmt.Errorf("auth not set")
	}

	t, err := a.auth.GetAccessToken(ctx, options.Scopes)
	if err != nil {
		return azcore.AccessToken{}, err
	}

	return azcore.AccessToken{Token: t.Token, ExpiresOn: t.ExpiresOn}, nil
}
