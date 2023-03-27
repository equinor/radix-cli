package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/settings"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	apiEndpointPatternForContext = "api.%sradix.equinor.com"
	apiEndpointPatternForCluster = "server-radix-api-%s.%s.dev.radix.equinor.com"

	// TokenEnvironmentName Name of environment variable to load token from
	TokenEnvironmentName = "APP_SERVICE_ACCOUNT_TOKEN"
)

// Get Gets API client for set context
func Get() *apiclient.Radixapi {
	return GetForContext("", false)
}

// GetForCommand Gets client for command
func GetForCommand(cmd *cobra.Command) (*apiclient.Radixapi, error) {
	token, err := getTokenFromFlagSet(cmd)
	if err != nil {
		return nil, err
	}

	context, cluster, err := getContextAndCluster(cmd)
	if err != nil {
		return nil, err
	}
	apiEnvironment, _ := cmd.Flags().GetString(settings.ApiEnvironmentOption)
	verboseOutput, _ := cmd.Flags().GetBool(settings.VerboseOption)
	var apiClient *apiclient.Radixapi
	if token != nil && *token != "" {
		apiClient = GetForToken(context, cluster, apiEnvironment, *token, verboseOutput)
	} else if cluster != "" {
		apiClient = GetForCluster(cluster, apiEnvironment, verboseOutput)
	} else {
		apiClient = GetForContext(context, verboseOutput)
	}

	return apiClient, nil
}

// LoginCommand Login client for command
func LoginCommand(cmd *cobra.Command) error {
	context, _, err := getContextAndCluster(cmd)
	if err != nil {
		return err
	}
	return LoginContext(context)
}

// LogoutCommand Logout command
func LogoutCommand(cmd *cobra.Command) error {
	context, _, err := getContextAndCluster(cmd)
	if err != nil {
		return err
	}
	config := radixconfig.GetDefaultRadixConfig()
	config.CustomConfig.Context = context
	return radixconfig.Save(*config)
}

func getContextAndCluster(cmd *cobra.Command) (string, string, error) {
	context, _ := cmd.Flags().GetString("context")
	cluster, _ := cmd.Flags().GetString(settings.ClusterOption)

	if context != "" && cluster != "" {
		return "", "", errors.New("cannot use both context and cluster as arguments at the same time")
	}
	return context, cluster, nil
}

// GetForToken Gets API client with passed token
func GetForToken(context, cluster, environment, token string, verboseOutput bool) *apiclient.Radixapi {
	var apiEndpoint string

	if cluster != "" {
		apiEndpoint = getAPIEndpointForCluster(cluster, environment)
	} else {
		radixConfig := radixconfig.RadixConfigAccess{}
		apiEndpoint = getAPIEndpointForContext(context, radixConfig.GetStartingConfig())
	}

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = httptransport.BearerToken(token)
	transport.Debug = verboseOutput
	return apiclient.New(transport, strfmt.Default)
}

// GetForContext Gets API client for set context
func GetForContext(context string, verbose bool) *apiclient.Radixapi {
	radixConfig := radixconfig.RadixConfigAccess{}
	apiEndpoint := getAPIEndpointForContext(context, radixConfig.GetStartingConfig())
	return getClientForEndpoint(apiEndpoint, verbose)
}

// LoginContext Performs login
func LoginContext(context string) error {
	radixConfig := radixconfig.RadixConfigAccess{}
	defaultConfig := radixConfig.GetDefaultConfig()
	apiEndpoint := getAPIEndpointForContext(context, defaultConfig)
	transport := getTransport(apiEndpoint, radixconfig.RadixConfigAccess{}, defaultConfig)
	_, err := transport.Transport.RoundTrip(&http.Request{})
	if err != nil && err.Error() == "http: nil Request.URL" {
		return nil
	}
	return err
}

// GetForCluster Gets API client for cluster
func GetForCluster(cluster, environment string, verboseOutput bool) *apiclient.Radixapi {
	apiEndpoint := getAPIEndpointForCluster(cluster, environment)
	return getClientForEndpoint(apiEndpoint, verboseOutput)
}

func getClientForEndpoint(apiEndpoint string, verbose bool) *apiclient.Radixapi {
	radixConfig := radixconfig.RadixConfigAccess{}
	startingConfig := radixConfig.GetStartingConfig()
	transport := getTransport(apiEndpoint, radixConfig, startingConfig)
	transport.Debug = verbose
	return apiclient.New(transport, strfmt.Default)
}

func getTransport(apiEndpoint string, radixConfig radixconfig.RadixConfigAccess, startingConfig *clientcmdapi.AuthProviderConfig) *httptransport.Runtime {
	persister := radixconfig.PersisterForRadix(radixConfig)
	provider, _ := rest.GetAuthProvider("", startingConfig, persister)

	schema := "https"
	if os.Getenv("USE_LOCAL_RADIX_API") == "true" {
		schema = "http"
		apiEndpoint = "localhost:3002"
	}
	transport := httptransport.New(apiEndpoint, "/api/v1", []string{schema})
	transport.Transport = provider.WrapTransport(transport.Transport)
	return transport
}

func getAPIEndpointForContext(context string, config *clientcmdapi.AuthProviderConfig) string {
	if strings.TrimSpace(context) == "" {
		context = config.Config["context"]
	}
	return fmt.Sprintf(apiEndpointPatternForContext, getPatternForContext(context))
}

func getAPIEndpointForCluster(cluster, environment string) string {
	return fmt.Sprintf(apiEndpointPatternForCluster, environment, cluster)
}

func getPatternForContext(context string) string {
	contextToPattern := make(map[string]string)
	contextToPattern[radixconfig.ContextDevelopment] = "dev."
	contextToPattern[radixconfig.ContextPlayground] = "playground."
	contextToPattern[radixconfig.ContextPlatform2] = "c2."
	contextToPattern[radixconfig.ContextProduction] = ""
	contextToPattern[radixconfig.ContextPlatform] = ""
	return contextToPattern[context]
}

func getTokenFromFlagSet(cmd *cobra.Command) (*string, error) {
	var token string
	tokenFromStdIn, _ := cmd.Flags().GetBool(settings.TokenStdinOption)
	tokenFromEnvironment, _ := cmd.Flags().GetBool(settings.TokenEnvironmentOption)

	if tokenFromStdIn && tokenFromEnvironment {
		return nil, errors.New("`token-stdin` and `token-environment` cannot both be set")
	}

	if tokenFromStdIn {
		contents, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return nil, err
		}

		token = strings.TrimSuffix(string(contents), "\n")
		token = strings.TrimSuffix(token, "\r")
	} else if tokenFromEnvironment {
		token = os.Getenv(TokenEnvironmentName)
		if strings.EqualFold(token, "") {
			return nil, fmt.Errorf("environment variable `%s` should be set", TokenEnvironmentName)
		}
	}

	return &token, nil
}
