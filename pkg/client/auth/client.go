package auth

import (
	"fmt"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
)

// New creates a new authentication client
func New(radixConfig *radixconfig.RadixConfig) (*public.Client, error) {
	cacheAccessor := NewTokenCache(radixConfig)
	cache := public.WithCache(cacheAccessor)
	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", radixConfig.TenantID)
	client, err := public.New(radixConfig.ClientID, cache, public.WithAuthority(authority))
	if err != nil {
		return nil, err
	}
	return &client, nil
}
