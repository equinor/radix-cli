package auth

import (
	"context"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	customJSON "github.com/equinor/radix-cli/pkg/client/auth/internal/json"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	jsonutils "github.com/equinor/radix-cli/pkg/utils/json"
	log "github.com/sirupsen/logrus"
	"os"
)

type TokenCache struct {
	file        string
	radixConfig *radixconfig.RadixConfig
}

func NewTokenCache(radixConfig *radixconfig.RadixConfig) *TokenCache {
	return &TokenCache{
		file:        radixconfig.RecommendedHomeMsalContractFile,
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
	t.radixConfig.MSALContract = contract
	return cache.Unmarshal(data)
}

func (t *TokenCache) Export(ctx context.Context, cache cache.Marshaler, hints cache.ExportHints) error {
	data, err := cache.Marshal()
	if err != nil {
		log.Println(err)
	}
	msalContract := radixconfig.NewContract()
	if err := customJSON.Unmarshal(data, msalContract); err != nil {
		return err
	}
	t.radixConfig.MSALContract = msalContract
	err = radixconfig.Save(t.radixConfig)
	if err != nil {
		return err
	}
	err = ensureMsalContractFileExists()
	if err != nil {
		return err
	}
	return os.WriteFile(t.file, data, 0600)
}

func ensureMsalContractFileExists() error {
	if _, err := os.Stat(radixconfig.RecommendedConfigDir); os.IsNotExist(err) {
		err := os.MkdirAll(radixconfig.RecommendedConfigDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(radixconfig.RecommendedHomeMsalContractFile); err == nil {
		return jsonutils.Save(radixconfig.RecommendedHomeMsalContractFile, radixconfig.NewContract())
	}
	return nil
}
