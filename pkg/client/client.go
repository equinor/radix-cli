package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/pkg/client/auth"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/settings"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

const (
	apiEndpointPatternForContext = "api.%sradix.equinor.com"
	apiEndpointPatternForCluster = "server-radix-api-%s.%s.dev.radix.equinor.com"

	// TokenEnvironmentName Name of environment variable to load token from
	TokenEnvironmentName = "APP_SERVICE_ACCOUNT_TOKEN"
)

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
	if token != nil && *token != "" {
		return GetForToken(context, cluster, apiEnvironment, *token, verboseOutput)
	}
	if cluster != "" {
		return GetForCluster(cluster, apiEnvironment, verboseOutput)
	}
	return GetForContext(context, verboseOutput)
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
func LogoutCommand() error {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return err
	}
	provider, err := getAuthProvider(radixConfig)
	if err != nil {
		return err
	}
	return provider.Logout()
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
func GetForToken(context, cluster, environment, token string, verboseOutput bool) (*apiclient.Radixapi, error) {
	var apiEndpoint string

	if cluster != "" {
		apiEndpoint = getAPIEndpointForCluster(cluster, environment)
	} else {
		radixConfig, err := radixconfig.GetRadixConfig()
		if err != nil {
			return nil, err
		}
		apiEndpoint = getAPIEndpointForContext(context, radixConfig)
	}

	transport := httptransport.New(apiEndpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = httptransport.BearerToken(token)
	transport.Debug = verboseOutput
	return apiclient.New(transport, strfmt.Default), nil
}

// GetForContext Gets API client for set context
func GetForContext(context string, verbose bool) (*apiclient.Radixapi, error) {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return nil, err
	}
	apiEndpoint := getAPIEndpointForContext(context, radixConfig)
	return getClientForEndpoint(apiEndpoint, verbose)
}

// LoginContext Performs login
func LoginContext(context string) error {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return err
	}
	apiEndpoint := getAPIEndpointForContext(context, radixConfig)
	transport, _ := getTransport(apiEndpoint, radixConfig)
	if _, err := transport.Transport.RoundTrip(&http.Request{}); err != nil && err.Error() != "http: nil Request.URL" {
		return err
	}
	return nil
}

// GetForCluster Gets API client for cluster
func GetForCluster(cluster, environment string, verboseOutput bool) (*apiclient.Radixapi, error) {
	apiEndpoint := getAPIEndpointForCluster(cluster, environment)
	return getClientForEndpoint(apiEndpoint, verboseOutput)
}

func getClientForEndpoint(apiEndpoint string, verbose bool) (*apiclient.Radixapi, error) {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return nil, err
	}
	transport, err := getTransport(apiEndpoint, radixConfig)
	if err != nil {
		return nil, err
	}
	transport.Debug = verbose

	return apiclient.New(transport, strfmt.Default), nil
}

func getTransport(apiEndpoint string, radixConfig *radixconfig.RadixConfig) (*httptransport.Runtime, error) {
	provider, err := getAuthProvider(radixConfig)
	if err != nil {
		return nil, err
	}
	schema := "https"
	if os.Getenv("USE_LOCAL_RADIX_API") == "true" {
		schema = "http"
		apiEndpoint = "localhost:3002"
	}
	transport := httptransport.New(apiEndpoint, "/api/v1", []string{schema})
	transport.Transport = provider.WrapTransport(transport.Transport)
	return transport, nil
}

func getAuthProvider(radixConfig *radixconfig.RadixConfig) (auth.MSALAuthProvider, error) {
	provider, err := auth.NewMSALAuthProvider(radixConfig)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func getAPIEndpointForContext(context string, radixConfig *radixconfig.RadixConfig) string {
	if strings.TrimSpace(context) == "" {
		context = radixConfig.CustomConfig.Context
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
