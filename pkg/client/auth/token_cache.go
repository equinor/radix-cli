package auth

import (
	"context"
	"encoding/json"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
)

// TokenCache is a token cache
type TokenCache struct {
	// file        string
	radixConfig *radixconfig.RadixConfig
}

// NewTokenCache creates a new token cache
func NewTokenCache(radixConfig *radixconfig.RadixConfig) *TokenCache {
	return &TokenCache{
		// file:        radixconfig.MsalContractFileFullName,
		radixConfig: radixConfig,
	}
}

// Replace replaces the cache with what is in external storage. Implementors should honor
// Context cancellations and return context.Canceled or context.DeadlineExceeded in those cases.
func (t *TokenCache) Replace(ctx context.Context, cache cache.Unmarshaler, hints cache.ReplaceHints) error {
	// data, err := os.ReadFile(t.file)
	// if err != nil {
	// 	log.Println(err)
	// }
	var (
		data []byte
		err  error
	)
	if t.radixConfig.MSALContract != nil {
		data, err = json.Marshal(t.radixConfig.MSALContract)
		if err != nil {
			return err
		}
	}
	// contract := radixconfig.NewContract()
	// if err := json.Unmarshal(data, contract); err != nil {
	// 	fmt.Printf("error loading the MSAL contract: %v\nClean it.\n", err)
	// 	if err = jsonutils.Save(radixconfig.MsalContractFileFullName, radixconfig.NewContract()); err != nil {
	// 		return err
	// 	}
	// }
	// t.radixConfig.MSALContract = contract
	return cache.Unmarshal(data)
}

// Export writes the binary representation of the cache (cache.Marshal()) to external storage.
// This is considered opaque. Context cancellations should be honored as in Replace.
func (t *TokenCache) Export(ctx context.Context, cache cache.Marshaler, hints cache.ExportHints) error {
	data, err := cache.Marshal()
	if err != nil {
		// log.Println(err)
		return err
	}
	var contract radixconfig.Contract
	// msalContract := radixconfig.NewContract()
	if err := json.Unmarshal(data, &contract); err != nil {
		return err
	}
	t.radixConfig.MSALContract = &contract
	return radixconfig.Save(t.radixConfig)
	// radixConfigData, err := json.Marshal(t.radixConfig)
	// t.radixConfig.MSALContract = msalContract
	// err = radixconfig.Save(t.radixConfig)
	// if err != nil {
	// 	return err
	// }
	// return os.WriteFile(t.file, data, 0600)
}
