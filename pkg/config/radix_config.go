package config

import (
	"os"
	"path"
	"reflect"

	"encoding/json"

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

	recommendedHomeDir           = ".radix"
	recommendedFileName          = "config"
	recommendedMsalCredsFileName = "creds"

	clientID    = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	tenantID    = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
	apiServerID = "6dae42f8-4368-4678-94ff-3960e28e3630"
	configMode  = "1" // Config mode "1" omits spn prefix from the aud (audience) in the token. "0" includes spn prefix

	defaultContext = ContextPlatform

	cfgContext      = "context"
	cfgAccessToken  = "access-token"
	cfgRefreshToken = "refresh-token"
	cfgExpiresIn    = "expires-in"
	cfgExpiresOn    = "expires-on"
	cfgEnvironment  = "environment"
	cfgConfigMode   = "config-mode"
)

var (
	RecommendedConfigDir         = path.Join(homedir.HomeDir(), recommendedHomeDir)
	RecommendedHomeFile          = path.Join(RecommendedConfigDir, recommendedFileName)
	RecommendedHomeMsalCredsFile = path.Join(RecommendedConfigDir, recommendedMsalCredsFileName)
	ValidContexts                = []string{ContextProduction, ContextPlatform, ContextPlatform2, ContextPlayground, ContextDevelopment}
)

type RadixConfig struct {
	CustomConfig  *CustomConfig  `json:"customConfig"`
	SessionConfig *SessionConfig `json:"sessionConfig"`
	ClientID      string         `json:"-"`
	TenantID      string         `json:"-"`
	APIServerID   string         `json:"-"`
	MsalContract  *Contract      `json:"-"`
}

type CustomConfig struct {
	Context string `json:"Context"`
}

type SessionConfig struct {
	RefreshToken string      `json:"refreshToken"`
	AccessToken  string      `json:"accessToken"`
	ExpiresIn    json.Number `json:"expiresIn"`
	ExpiresOn    json.Number `json:"expiresOn"`
	Environment  string      `json:"environment"`
	ConfigMode   string      `json:"configMode"`
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
	return getAzureAuthProvider(radixConfig)
}

// GetDefaultConfig Gets AuthProviderConfig with default properties
func (c RadixConfigAccess) GetDefaultConfig() *clientcmdapi.AuthProviderConfig {
	return getAzureAuthProvider(GetDefaultRadixConfig())
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
		MsalContract: NewContract(),
		SessionConfig: &SessionConfig{
			ConfigMode: configMode,
		},
	}
}

func getAzureAuthProvider(radixConfig *RadixConfig) *clientcmdapi.AuthProviderConfig {
	return &clientcmdapi.AuthProviderConfig{
		Name:   "msal-radix",
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

	config[cfgRefreshToken] = radixConfig.SessionConfig.RefreshToken
	config[cfgAccessToken] = radixConfig.SessionConfig.AccessToken
	config[cfgExpiresIn] = radixConfig.SessionConfig.ExpiresIn.String()
	config[cfgExpiresOn] = radixConfig.SessionConfig.ExpiresOn.String()
	config[cfgEnvironment] = radixConfig.SessionConfig.Environment
	config[cfgConfigMode] = radixConfig.SessionConfig.ConfigMode
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

	return NewRadixConfig(customConfig,
		&SessionConfig{
			RefreshToken: config[cfgRefreshToken],
			AccessToken:  config[cfgAccessToken],
			ExpiresIn:    json.Number(config[cfgExpiresIn]),
			ExpiresOn:    json.Number(config[cfgExpiresOn]),
			Environment:  config[cfgEnvironment],
			ConfigMode:   config[cfgConfigMode],
		})
}

func NewRadixConfig(config *CustomConfig, sessionConfig *SessionConfig) *RadixConfig {
	return &RadixConfig{
		CustomConfig:  config,
		SessionConfig: sessionConfig,
		ClientID:      clientID,
		TenantID:      tenantID,
		APIServerID:   apiServerID,
		MsalContract:  NewContract(),
	}
}
