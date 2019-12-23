package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	httptransport "github.com/go-openapi/runtime/client"
)

const (
	apiEndpointPatternForContext = "api.%sradix.equinor.com"
	apiEndpointPatternForCluster = "server-radix-api-%s.%s.dev.radix.equinor.com"

	// TokenEnvironmentName Name of environment variable to load token from
	TokenEnvironmentName = "APP_SERVICE_ACCOUNT_TOKEN"
)

// Get Gets API client for set context
func Get() *apiclient.Radixapi {
	return GetForContext("")
}

// GetForCommand Gets client for command
func GetForCommand(cmd *cobra.Command) (*apiclient.Radixapi, error) {
	context, _ := cmd.Flags().GetString("context")
	cluster, _ := cmd.Flags().GetString("cluster")
	environment, _ := cmd.Flags().GetString("environment")

	token, err := getTokenFromFlagSet(cmd)
	if err != nil {
		return nil, err
	}

	if context != "" && cluster != "" {
		return nil, errors.New("Cannot use both context and cluster as arguments at the same time")
	}

	var apiClient *apiclient.Radixapi
	if token != nil && *token != "" {
		apiClient = GetForToken(context, cluster, environment, *token)
	} else if cluster != "" {
		apiClient = GetForCluster(cluster, environment)
	} else {
		apiClient = GetForContext(context)
	}

	return apiClient, nil
}

// GetForToken Gets API client with passed token
func GetForToken(context, cluster, environment, token string) *apiclient.Radixapi {
	var apiEndpoint string

	if cluster != "" {
		apiEndpoint = getAPIEndpointForCluster(cluster, environment)
	} else {
		radixConfig := radixconfig.RadixConfigAccess{}
		startingConfig := radixConfig.GetStartingConfig()

		if strings.TrimSpace(context) == "" {
			context = startingConfig.Config["context"]
		}

		apiEndpoint = getAPIEndpointForContext(context)
	}

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = httptransport.BearerToken(token)
	return apiclient.New(transport, strfmt.Default)
}

// GetForContext Gets API client for set context
func GetForContext(context string) *apiclient.Radixapi {
	radixConfig := radixconfig.RadixConfigAccess{}
	startingConfig := radixConfig.GetStartingConfig()

	if strings.TrimSpace(context) == "" {
		context = startingConfig.Config["context"]
	}

	apiEndpoint := getAPIEndpointForContext(context)
	return getClientForEndpoint(apiEndpoint)
}

// GetForCluster Gets API client for cluster
func GetForCluster(cluster, environment string) *apiclient.Radixapi {
	apiEndpoint := getAPIEndpointForCluster(cluster, environment)
	return getClientForEndpoint(apiEndpoint)
}

func getClientForEndpoint(apiEndpoint string) *apiclient.Radixapi {
	radixConfig := radixconfig.RadixConfigAccess{}
	startingConfig := radixConfig.GetStartingConfig()
	persister := radixconfig.PersisterForRadix(radixConfig)
	provider, _ := rest.GetAuthProvider("", startingConfig, persister)

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	transport.Transport = provider.WrapTransport(transport.Transport)
	return apiclient.New(transport, strfmt.Default)
}

func getAPIEndpointForContext(context string) string {
	return fmt.Sprintf(apiEndpointPatternForContext, getPatternForContext(context))
}

func getAPIEndpointForCluster(cluster, environment string) string {
	return fmt.Sprintf(apiEndpointPatternForCluster, environment, cluster)
}

func getPatternForContext(context string) string {
	contextToPattern := make(map[string]string)
	contextToPattern[radixconfig.ContextDevelopment] = "dev."
	contextToPattern[radixconfig.ContextPlayground] = fmt.Sprintf("%s.", radixconfig.ContextPlayground)
	contextToPattern[radixconfig.ContextProdction] = ""
	return contextToPattern[context]
}

func getTokenFromFlagSet(cmd *cobra.Command) (*string, error) {
	var token string
	tokenFromStdIn, _ := cmd.Flags().GetBool("token-stdin")
	tokenFromEnvironment, _ := cmd.Flags().GetBool("token-environment")

	if tokenFromStdIn && tokenFromEnvironment {
		return nil, errors.New("`token-stdin` and `token-environment` cannot both be set")
	}

	if tokenFromStdIn {
		contents, err := ioutil.ReadAll(cmd.InOrStdin())
		if err != nil {
			return nil, err
		}

		token = strings.TrimSuffix(string(contents), "\n")
		token = strings.TrimSuffix(token, "\r")
	} else if tokenFromEnvironment {
		token = os.Getenv(TokenEnvironmentName)
		if strings.EqualFold(token, "") {
			return nil, fmt.Errorf("Environment variable`%s` should be set", TokenEnvironmentName)
		}
	}

	return &token, nil
}
