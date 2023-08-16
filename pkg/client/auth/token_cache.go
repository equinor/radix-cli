package auth

import (
	"context"
	"encoding/base64"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
)

// TokenCache is a token cache
type TokenCache struct {
	radixConfig *radixconfig.RadixConfig
}

// NewTokenCache creates a new token cache
func NewTokenCache(radixConfig *radixconfig.RadixConfig) *TokenCache {
	return &TokenCache{
		radixConfig: radixConfig,
	}
}

// Replace replaces the cache with what is in external storage. Implementors should honor
// Context cancellations and return context.Canceled or context.DeadlineExceeded in those cases.
func (t *TokenCache) Replace(ctx context.Context, cache cache.Unmarshaler, hints cache.ReplaceHints) error {
	var (
		data []byte
		err  error
	)
	if len(t.radixConfig.MSAL) > 0 {
		data, err = base64.RawStdEncoding.DecodeString(t.radixConfig.MSAL)
		// TODO: Should we print the error of decoing fails?
		if err != nil {
			data = nil // DecodeString can return a non-empty slice on error. Need to set it to nil to avoid errors in cache.Unmarshal
		}
	}
	return cache.Unmarshal(data)
}

// Export writes the binary representation of the cache (cache.Marshal()) to external storage.
// This is considered opaque. Context cancellations should be honored as in Replace.
func (t *TokenCache) Export(ctx context.Context, cache cache.Marshaler, hints cache.ExportHints) error {
	data, err := cache.Marshal()
	if err != nil {
		return err
	}

	t.radixConfig.MSAL = base64.StdEncoding.EncodeToString(data)
	return radixconfig.Save(t.radixConfig)
}
