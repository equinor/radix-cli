package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/pkg/client/auth"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

const (
	apiEndpointPatternForContext = "api.%sradix.equinor.com"
	apiEndpointPatternForCluster = "server-radix-api-%s.%s.dev.radix.equinor.com"

	// TokenEnvironmentName Name of environment variable to load token from
	TokenEnvironmentName = "APP_SERVICE_ACCOUNT_TOKEN"

	// MSAL authentication
	clientID = "ed6cb804-8193-4e55-9d3d-8b88688482b3"
	tenantID = "3aa4a235-b6e2-48d5-9195-7fcf05b459b0"
)

// GetForCommand Gets client for command
func GetForCommand(cmd *cobra.Command) (*apiclient.Radixapi, error) {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return nil, err
	}
	authWriter, err := getAuthWriter(cmd, radixConfig)
	if err != nil {
		return nil, nil
	}
	endpoint, err := getAPIEndpoint(cmd, radixConfig)
	if err != nil {
		return nil, err
	}
	verboseOutput, _ := cmd.Flags().GetBool(settings.VerboseOption)
	transport := getTransport(endpoint, authWriter, verboseOutput)
	return apiclient.New(transport, strfmt.Default), nil
}

func getTransport(endpoint string, authWriter runtime.ClientAuthInfoWriter, verbose bool) *httptransport.Runtime {
	transport := httptransport.New(endpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = authWriter
	transport.Debug = verbose
	return transport
}

func getAPIEndpoint(cmd *cobra.Command, config *radixconfig.RadixConfig) (string, error) {
	context, cluster, err := getContextAndCluster(cmd)
	if err != nil {
		return "", err
	}

	if cluster != "" {
		apiEnvironment, _ := cmd.Flags().GetString(settings.ApiEnvironmentOption)
		return getAPIEndpointForCluster(cluster, apiEnvironment), nil
	}

	return getAPIEndpointForContext(context, config), nil
}

func getAuthWriter(cmd *cobra.Command, config *radixconfig.RadixConfig) (runtime.ClientAuthInfoWriter, error) {
	token, err := getTokenFromFlagSet(cmd)
	if err != nil {
		return nil, err
	}

	if token != nil && *token != "" {
		return httptransport.BearerToken(*token), nil
	}

	return getAuthProvider(config)
}

// LoginCommand Login client for command
func LoginCommand(cmd *cobra.Command) error {
	return LoginContext()
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
	return provider.Logout(context.Background())
}

func getContextAndCluster(cmd *cobra.Command) (string, string, error) {
	context, _ := cmd.Flags().GetString("context")
	cluster, _ := cmd.Flags().GetString(settings.ClusterOption)

	if context != "" && cluster != "" {
		return "", "", errors.New("cannot use both context and cluster as arguments at the same time")
	}
	return context, cluster, nil
}

// LoginContext Performs login
func LoginContext() error {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return err
	}
	contextName := radixConfig.CustomConfig.Context
	radixConfig = radixconfig.GetDefaultRadixConfig()
	radixConfig.CustomConfig.Context = contextName
	provider, err := getAuthProvider(radixConfig)
	if err != nil {
		return err
	}
	return provider.Login(context.Background())
}

func getAuthProvider(radixConfig *radixconfig.RadixConfig) (auth.MSALAuthProvider, error) {
	provider, err := auth.NewMSALAuthProvider(radixConfig, clientID, tenantID)
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
