package config

import (
	"os"
	"path"
	"time"

	jsonutils "github.com/equinor/radix-cli/pkg/utils/json"
)

const (
	ContextProduction  = "production"
	ContextPlatform    = "platform"
	ContextPlayground  = "playground"
	ContextDevelopment = "development"
	ContextPlatform2   = "platform2"

	radixConfigDir      = ".radix"
	radixConfigFileName = "config"

	defaultContext       = ContextPlatform
	DefaultCacheDuration = 7 * 24 * time.Hour
)

var (
	RadixConfigDir          = path.Join(getUserHomeDir(), radixConfigDir)
	RadixConfigFileFullName = path.Join(RadixConfigDir, radixConfigFileName)
	ValidContexts           = []string{ContextProduction, ContextPlatform, ContextPlatform2, ContextPlayground, ContextDevelopment}
)

func getUserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir
}

type CacheItem struct {
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Content   string    `json:"content"`
}
type RadixConfig struct {
	// CustomConfig is the custom environment config
	CustomConfig *CustomConfig                   `json:"customConfig"`
	Cache        map[string]map[string]CacheItem `json:"cache"`

	// MSAL is the internal cache structure used by the MSAL module. The content is base64 encoded
	MSAL string `json:"msal,omitempty"`
}

// CustomConfig is the custom environment config
type CustomConfig struct {
	// Context is the environment context: platform (default), playground, development, platform2
	Context string `json:"Context"`
}

func IsValidContext(context string) bool {
	for _, validContext := range ValidContexts {
		if validContext == context {
			return true
		}
	}
	return false
}

func GetRadixConfig() (*RadixConfig, error) {
	radixConfig := &RadixConfig{}
	err := jsonutils.Load(RadixConfigFileFullName, radixConfig)
	if err == nil {
		return radixConfig, nil
	}

	radixConfig = GetDefaultRadixConfig()
	if err = jsonutils.Save(RadixConfigFileFullName, radixConfig); err != nil {
		return nil, err
	}
	return radixConfig, nil
}

func GetDefaultRadixConfig() *RadixConfig {
	return &RadixConfig{
		CustomConfig: &CustomConfig{
			Context: string(defaultContext),
		},
	}
}

func GetCache(key string) (string, bool) {
	config, err := GetRadixConfig()
	if err != nil {
		return "", false
	}
	item, ok := config.Cache[config.CustomConfig.Context][key]
	if !ok {
		return "", false
	}

	if time.Now().After(item.ExpiresAt) {
		return "", false
	}

	return item.Content, ok
}
func SetCache(key, content string, ttl time.Duration) {
	config, _ := GetRadixConfig()

	if config.Cache == nil {
		config.Cache = make(map[string]map[string]CacheItem)
	}
	if _, ok := config.Cache[config.CustomConfig.Context]; !ok {
		config.Cache[config.CustomConfig.Context] = make(map[string]CacheItem)
	}

	config.Cache[config.CustomConfig.Context][key] = CacheItem{
		UpdatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl),
		Content:   content,
	}

	_ = Save(config)
}

// Save Saves RadixConfig
func Save(radixConfig *RadixConfig) error {
	return jsonutils.Save(RadixConfigFileFullName, *radixConfig)
}
