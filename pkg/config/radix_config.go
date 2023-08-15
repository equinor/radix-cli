package config

import (
	"fmt"
	"os"
	"path"

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

	clientID    = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	tenantID    = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
	apiServerID = "6dae42f8-4368-4678-94ff-3960e28e3630"

	defaultContext = ContextPlatform
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

type RadixConfig struct {
	// CustomConfig is the custom environment config
	CustomConfig *CustomConfig `json:"customConfig"`
	// ClientID is the ID of the Azure AD app registration
	ClientID string `json:"-"`
	// TenantID is the ID of the tenant that the client will authenticate with
	TenantID string `json:"-"`
	// APIServerID is the ID of the API server that the client will use to get tokens
	APIServerID string `json:"-"`
	// MSALContract is the MSAL internal structure that is written to any storage medium when serializing the cache
	MSALContract *Contract `json:"contract"`
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
	radixConfig := getDefaultRadixConfig()
	err := jsonutils.Load(RadixConfigFileFullName, radixConfig)
	if err == nil {
		return radixConfig, nil
	}
	fmt.Println("Cannot load a RadixConfig, creating a new one with the context 'platform'.")
	radixConfig = getDefaultRadixConfig()
	if err = jsonutils.Save(RadixConfigFileFullName, radixConfig); err != nil {
		return nil, err
	}
	return radixConfig, nil
}

func getDefaultRadixConfig() *RadixConfig {
	return &RadixConfig{
		CustomConfig: &CustomConfig{
			Context: defaultContext,
		},
		ClientID:     clientID,
		TenantID:     tenantID,
		APIServerID:  apiServerID,
		MSALContract: NewContract(),
	}
}

// Save Saves RadixConfig
func (radixConfig *RadixConfig) Save() error {
	return jsonutils.Save(RadixConfigFileFullName, *radixConfig)
}
