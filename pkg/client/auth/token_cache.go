package auth

import (
	"context"
	"encoding/json"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	log "github.com/sirupsen/logrus"
)

// TokenCache is a token cache
type TokenCache struct {
	file        string
	radixConfig *radixconfig.RadixConfig
}

// NewTokenCache creates a new token cache
func NewTokenCache(radixConfig *radixconfig.RadixConfig) *TokenCache {
	return &TokenCache{
		file:        radixconfig.RadixConfigFileFullName,
		radixConfig: radixConfig,
	}
}

// Replace replaces the cache with what is in external storage. Implementors should honor
// Context cancellations and return context.Canceled or context.DeadlineExceeded in those cases.
func (t *TokenCache) Replace(ctx context.Context, cache cache.Unmarshaler, hints cache.ReplaceHints) error {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return err
	}
	contractData, err := json.Marshal(radixConfig.MSALContract)
	if err != nil {
		return err
	}
	return cache.Unmarshal(contractData)
}

// Export writes the binary representation of the cache (cache.Marshal()) to external storage.
// This is considered opaque. Context cancellations should be honored as in Replace.
func (t *TokenCache) Export(ctx context.Context, cache cache.Marshaler, hints cache.ExportHints) error {
	contractData, err := cache.Marshal()
	if err != nil {
		log.Println(err)
	}
	msalContract := radixconfig.NewContract()
	if err := json.Unmarshal(contractData, msalContract); err != nil {
		return err
	}
	t.radixConfig.MSALContract = msalContract
	return t.radixConfig.Save()
}
