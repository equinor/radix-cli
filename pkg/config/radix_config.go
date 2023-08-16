package config

import (
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
			Context: defaultContext,
		},
		// MSALContract: NewContract(),
	}
}

// Save Saves RadixConfig
func Save(radixConfig *RadixConfig) error {
	return jsonutils.Save(RadixConfigFileFullName, *radixConfig)
}
