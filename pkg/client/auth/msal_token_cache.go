package auth

import (
	"context"
	"time"

	azurecache "github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	"github.com/equinor/radix-cli/pkg/cache"
)

// MsalTokenCache is a token azurecache
type MsalTokenCache struct {
	cache cache.Cache
	key   string
}

// NewMsalTokenCache creates a new token azurecache
func NewMsalTokenCache(cache cache.Cache, key string) *MsalTokenCache {
	return &MsalTokenCache{
		cache: cache,
		key:   key,
	}
}

// Replace replaces the cache with what is in external storage. Implementors should honor
// RadixCluster cancellations and return context.Canceled or context.DeadlineExceeded in those cases.
func (t *MsalTokenCache) Replace(ctx context.Context, cache azurecache.Unmarshaler, hints azurecache.ReplaceHints) error {
	content, ok := t.cache.GetItem(t.key)
	if !ok {
		return cache.Unmarshal(nil)
	}
	return cache.Unmarshal([]byte(content))
}

// Export writes the binary representation of the cache (cache.Marshal()) to external storage.
// This is considered opaque. RadixCluster cancellations should be honored as in Replace.
func (t *MsalTokenCache) Export(ctx context.Context, cache azurecache.Marshaler, hints azurecache.ExportHints) error {
	data, err := cache.Marshal()
	if err != nil {
		return err
	}

	t.cache.SetItem(t.key, string(data), 24*365*time.Hour)
	return nil
}
