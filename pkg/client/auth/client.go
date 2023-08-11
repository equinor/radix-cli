package auth

import (
	"context"
	"fmt"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
)

func NewClient(radixConfig *radixconfig.RadixConfig) (*public.Client, error) {
	cacheAccessor := NewTokenCache(radixConfig)
	cache := public.WithCache(cacheAccessor)
	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", radixConfig.TenantID)
	client, err := public.New(radixConfig.ClientID, cache, public.WithAuthority(authority))
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func getExistingAccount(ctx context.Context, client *public.Client) (public.Account, error) {
	accounts, err := client.Accounts(ctx)
	if err != nil {
		return public.Account{}, err
	}
	account := public.Account{}
	if len(accounts) > 0 {
		account = accounts[0]
		return account, nil
	}
	return public.Account{}, nil
}
