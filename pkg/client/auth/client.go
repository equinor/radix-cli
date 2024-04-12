package auth

import (
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
)

// newPublicClient creates a new authentication client
func newPublicClient(radixConfig *radixconfig.RadixConfig, clientID, authority string) (*public.Client, error) {
	cacheAccessor := NewTokenCache(radixConfig)
	cache := public.WithCache(cacheAccessor)
	client, err := public.New(clientID, cache, public.WithAuthority(authority))
	if err != nil {
		return nil, err
	}
	return &client, nil
}
