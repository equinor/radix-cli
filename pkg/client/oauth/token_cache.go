package oauth

import (
	"context"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	customJSON "github.com/equinor/radix-cli/pkg/client/oauth/internal/json"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	log "github.com/sirupsen/logrus"
	"os"
)

type TokenCache struct {
	file        string
	radixConfig *radixconfig.RadixConfig
}

func NewTokenCache(radixConfig *radixconfig.RadixConfig) *TokenCache {
	return &TokenCache{
		file:        radixconfig.RecommendedHomeMsalCredsFile,
		radixConfig: radixConfig,
	}
}

func (t *TokenCache) Replace(ctx context.Context, cache cache.Unmarshaler, hints cache.ReplaceHints) error {
	data, err := os.ReadFile(t.file)
	if err != nil {
		log.Println(err)
	}
	contract := radixconfig.NewContract()
	if err := customJSON.Unmarshal(data, contract); err != nil {
		return err
	}
	t.radixConfig.MsalContract = contract
	return cache.Unmarshal(data)
}

func (t *TokenCache) Export(ctx context.Context, cache cache.Marshaler, hints cache.ExportHints) error {
	data, err := cache.Marshal()
	if err != nil {
		log.Println(err)
	}
	contract := radixconfig.NewContract()
	if err := customJSON.Unmarshal(data, contract); err != nil {
		return err
	}
	t.radixConfig.MsalContract = contract
	return os.WriteFile(t.file, data, 0600)
}
