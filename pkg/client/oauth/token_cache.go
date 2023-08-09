package oauth

import (
	"context"
	"encoding/json"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	customJSON "github.com/equinor/radix-cli/pkg/client/oauth/internal/json"
	"github.com/equinor/radix-cli/pkg/client/oauth/internal/storage"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

type TokenCache struct {
	file        string
	radixConfig *radixconfig.RadixConfig
}

func (t *TokenCache) Replace(ctx context.Context, cache cache.Unmarshaler, hints cache.ReplaceHints) error {
	data, err := os.ReadFile(t.file)
	if err != nil {
		log.Println(err)
	}
	// contract := storage.NewContract()
	// if err := customJSON.Unmarshal(data, contract); err != nil {
	// 	return err
	// }
	return cache.Unmarshal(data)
}

func (t *TokenCache) Export(ctx context.Context, cache cache.Marshaler, hints cache.ExportHints) error {
	data, err := cache.Marshal()
	if err != nil {
		log.Println(err)
	}
	contract := storage.NewContract()
	if err := customJSON.Unmarshal(data, contract); err != nil {
		return err
	}
	err = t.updateRadixSessionConfig(contract)
	if err != nil {
		return err
	}
	return os.WriteFile(t.file, data, 0600)
}

func (t *TokenCache) updateRadixSessionConfig(contract *storage.Contract) error {
	key := getKey(contract.Accounts)
	if len(key) == 0 {
		t.radixConfig.SessionConfig = &radixconfig.SessionConfig{
			TenantID:    t.radixConfig.SessionConfig.TenantID,
			ClientID:    t.radixConfig.SessionConfig.ClientID,
			APIServerID: t.radixConfig.SessionConfig.APIServerID,
		}
		return nil
	}
	if token, ok := contract.AccessTokens[key]; ok {
		t.radixConfig.SessionConfig.AccessToken = token.Secret
		t.radixConfig.SessionConfig.ExpiresOn = json.Number(token.ExpiresOn)
		t.radixConfig.SessionConfig.ExpiresIn = json.Number(getExpiresIn(token.ExpiresOn))
	}
	if token, ok := contract.RefreshTokens[key]; ok {
		t.radixConfig.SessionConfig.RefreshToken = token.Secret
	}

	config := t.radixConfig
	return radixconfig.Save(*config)
}

func getExpiresIn(expiresOn string) string {
	if expiresOnNum, err := strconv.Atoi(expiresOn); err == nil {
		expiresIn := time.Unix(int64(expiresOnNum), 0).Unix() - time.Now().Unix()
		if expiresIn > 0 {
			return string(expiresIn)
		}
	}
	return "0"
}

func getKey(accounts map[string]storage.Account) string {
	for key := range accounts {
		return key
	}
	return ""
}
