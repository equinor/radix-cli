package config

import (
	"os"
	"path"
	"reflect"

	jsonutils "github.com/equinor/radix-cli/pkg/utils/json"
	restclient "k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

const (
	ContextProduction  = "production"
	ContextPlatform    = "platform"
	ContextPlayground  = "playground"
	ContextDevelopment = "development"
	ContextPlatform2   = "platform2"

	recommendedHomeDir              = ".radix"
	recommendedFileName             = "config"
	recommendedMsalContractFileName = "contract"

	clientID    = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	tenantID    = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
	apiServerID = "6dae42f8-4368-4678-94ff-3960e28e3630"

	defaultContext = ContextPlatform

	cfgContext = "context"
)

var (
	RecommendedConfigDir            = path.Join(homedir.HomeDir(), recommendedHomeDir)
	RecommendedHomeFile             = path.Join(RecommendedConfigDir, recommendedFileName)
	RecommendedHomeMsalContractFile = path.Join(RecommendedConfigDir, recommendedMsalContractFileName)
	ValidContexts                   = []string{ContextProduction, ContextPlatform, ContextPlatform2, ContextPlayground, ContextDevelopment}
)

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
	MSALContract *Contract `json:"-"`
}

// CustomConfig is the custom environment config
type CustomConfig struct {
	// Context is the environment context: platform (default), playground, development, platform2
	Context string `json:"Context"`
}

type RadixConfigAccess struct {
}

func IsValidContext(context string) bool {
	for _, validContext := range ValidContexts {
		if validContext == context {
			return true
		}
	}
	return false
}

func (c RadixConfigAccess) GetStartingConfig() *clientcmdapi.AuthProviderConfig {
	var radixConfig *RadixConfig
	if _, err := os.Stat(RecommendedHomeFile); err == nil {
		radixConfig = &RadixConfig{}
		jsonutils.Load(RecommendedHomeFile, radixConfig)
	} else {
		radixConfig = GetDefaultRadixConfig()
	}
	return GetAuthProviderConfig(radixConfig)
}

// GetDefaultConfig Gets AuthProviderConfig with default properties
func (c RadixConfigAccess) GetDefaultConfig() *clientcmdapi.AuthProviderConfig {
	return GetAuthProviderConfig(GetDefaultRadixConfig())
}

// GetDefaultRadixConfig Gets RadixConfig with default properties
func GetDefaultRadixConfig() *RadixConfig {
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

// GetAuthProviderConfig Gets AuthProviderConfig with properties from RadixConfig
func GetAuthProviderConfig(radixConfig *RadixConfig) *clientcmdapi.AuthProviderConfig {
	return &clientcmdapi.AuthProviderConfig{
		Name:   "msal",
		Config: toMap(radixConfig),
	}
}

func (c RadixConfigAccess) GetExplicitFile() string {
	return "radix_config"
}

func PersisterForRadix(radixConfig RadixConfigAccess) restclient.AuthProviderConfigPersister {
	return &radixConfigPersister{
		radixConfig,
	}
}

type radixConfigPersister struct {
	radixConfig RadixConfigAccess
}

// Persist Persists config to file
func (p *radixConfigPersister) Persist(config map[string]string) error {
	startingConfig := ToConfig(p.radixConfig.GetStartingConfig().Config)
	newConfig := ToConfig(config)

	if newConfig.CustomConfig == nil {
		// When token is expired the newconfig doesn't come with the custom config set
		newConfig.CustomConfig = startingConfig.CustomConfig
	}

	if reflect.DeepEqual(startingConfig, newConfig) {
		return nil
	}

	return Save(newConfig)
}

// Save Saves RadixConfig
func Save(radixConfig *RadixConfig) error {
	if _, err := os.Stat(RecommendedConfigDir); os.IsNotExist(err) {
		os.MkdirAll(RecommendedConfigDir, os.ModePerm)
	}
	return jsonutils.Save(RecommendedHomeFile, *radixConfig)
}

func toMap(radixConfig *RadixConfig) map[string]string {
	config := make(map[string]string)
	if radixConfig.CustomConfig != nil {
		config[cfgContext] = radixConfig.CustomConfig.Context
	}
	return config
}

// ToConfig create RadixConfig from a map
func ToConfig(config map[string]string) *RadixConfig {
	var customConfig *CustomConfig
	if _, ok := config[cfgContext]; ok {
		customConfig = &CustomConfig{
			Context: config[cfgContext],
		}
	}

	radixConfig := NewRadixConfig()
	radixConfig.CustomConfig = customConfig
	return radixConfig
}

// NewRadixConfig Gets RadixConfig with default properties
func NewRadixConfig() *RadixConfig {
	return &RadixConfig{
		CustomConfig: &CustomConfig{},
		ClientID:     clientID,
		TenantID:     tenantID,
		APIServerID:  apiServerID,
		MSALContract: NewContract(),
	}
}
