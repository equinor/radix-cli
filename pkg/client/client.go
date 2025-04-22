package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	vulnscanapi "github.com/equinor/radix-cli/generated/vulnscanapi/client"
	"github.com/equinor/radix-cli/pkg/client/auth"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

const (
	// TokenEnvironmentName Name of environment variable to load token from
	TokenEnvironmentName = "APP_SERVICE_ACCOUNT_TOKEN"
)

// GetRadixApiForCommand Gets radixapi for command
func GetRadixApiForCommand(cmd *cobra.Command) (*radixapi.Radixapi, error) {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return nil, err
	}
	authWriter, err := getAuthWriter(cmd)
	if err != nil {
		return nil, nil
	}
	context, cluster, apiEnvironment := getContextClusterApiEnvironment(cmd, radixConfig)
	endpoint := getEndpoint("server-radix-api", apiEnvironment, context, cluster)
	verboseOutput, _ := cmd.Flags().GetBool(flagnames.Verbose)
	transport := getTransport(endpoint, authWriter, verboseOutput)
	silenceError, _ := cmd.Flags().GetBool(flagnames.SilenceError)
	cmd.SilenceErrors = silenceError
	return radixapi.New(transport, strfmt.Default), nil
}

// GetVulnerabilityScanApiForCommand Gets radixapi for command
func GetVulnerabilityScanApiForCommand(cmd *cobra.Command) (*vulnscanapi.Vulnscanapi, error) {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return nil, err
	}
	authWriter, err := getAuthWriter(cmd)
	if err != nil {
		return nil, nil
	}
	context, cluster, apiEnvironment := getContextClusterApiEnvironment(cmd, radixConfig)
	endpoint := getEndpoint("server-radix-vulnerability-scanner-api", apiEnvironment, context, cluster)
	verboseOutput, _ := cmd.Flags().GetBool(flagnames.Verbose)
	transport := getTransport(endpoint, authWriter, verboseOutput)
	return vulnscanapi.New(transport, strfmt.Default), nil
}

func getContextClusterApiEnvironment(cmd *cobra.Command, config *radixconfig.RadixConfig) (string, string, string) {
	context, _ := cmd.Flags().GetString("context")
	cluster, _ := cmd.Flags().GetString(flagnames.Cluster)
	apiEnvironment, _ := cmd.Flags().GetString(flagnames.ApiEnvironment)

	if strings.TrimSpace(context) == "" {
		context = config.CustomConfig.Context
	}
	return context, cluster, apiEnvironment
}

func getEndpoint(service, env, context, cluster string) string {
	zoneDomain, defaultEnv := getPatternForContext(context)
	if strings.TrimSpace(env) == "" {
		env = defaultEnv
	}

	if cluster != "" {
		return fmt.Sprintf("%s-%s.%s.%sradix.equinor.com", service, env, cluster, zoneDomain)
	}

	return fmt.Sprintf("%s-%s.%sradix.equinor.com", service, env, zoneDomain)
}

func getTransport(endpoint string, authWriter runtime.ClientAuthInfoWriter, verbose bool) *httptransport.Runtime {
	transport := httptransport.New(endpoint, "/api/v1", []string{"https"})
	transport.DefaultAuthentication = authWriter
	transport.Debug = verbose
	return transport
}

func getAuthWriter(cmd *cobra.Command) (runtime.ClientAuthInfoWriter, error) {
	token, err := getTokenFromFlagSet(cmd)
	if err != nil {
		return nil, err
	}

	if token != nil && *token != "" {
		return httptransport.BearerToken(*token), nil
	}

	return getAuthProvider()
}

// LoginCommand Login radixapi for command
func LoginCommand(ctx context.Context, useInteractiveLogin, useDeviceCode, useGithubCredentials bool, azureClientId, federatedTokenFile, azureClientSecret string) error {
	radixConfig, err := radixconfig.GetRadixConfig()
	if err != nil {
		return err
	}
	contextName := radixConfig.CustomConfig.Context
	radixConfig = radixconfig.GetDefaultRadixConfig()
	radixConfig.CustomConfig.Context = contextName
	provider, err := getAuthProvider()
	if err != nil {
		return err
	}
	return provider.Login(ctx, useInteractiveLogin, useDeviceCode, useGithubCredentials, azureClientId, federatedTokenFile, azureClientSecret)
}

// LogoutCommand Logout command
func LogoutCommand() error {
	provider, err := getAuthProvider()
	if err != nil {
		return err
	}
	return provider.Logout()
}

func getAuthProvider() (auth.Provider, error) {
	provider, err := auth.New()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func getPatternForContext(context string) (string, string) {
	switch context {
	case radixconfig.ContextDevelopment:
		return "dev.", "qa"
	case radixconfig.ContextPlayground:
		return "playground.", "prod"
	case radixconfig.ContextPlatform2:
		return "c2.", "prod"
	case radixconfig.ContextProduction, radixconfig.ContextPlatform:
		return "", "prod"
	default:
		return "", "prod"
	}
}

func getTokenFromFlagSet(cmd *cobra.Command) (*string, error) {
	var token string
	tokenFromStdIn, _ := cmd.Flags().GetBool(flagnames.TokenStdin)
	tokenFromEnvironment, _ := cmd.Flags().GetBool(flagnames.TokenEnvironment)

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
