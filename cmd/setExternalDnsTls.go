// Copyright © 2023
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// setExternalDnsTlsCmd represents the setExternalDnsTlsCmd command
var setExternalDnsTlsCmd = &cobra.Command{
	Use:   "external-dns-tls",
	Short: "Set TLS certificate and private key for a component's external DNS alias",
	Long:  "Set TLS certificate and private key for a component's external DNS alias",
	Example: `# Read certificate and private key from file
rx set external-dns-tls --application myapp --environment prod --component web --alias myapp.example.com --certificate-from-file "cert.crt" --private-key-from-file "cert.key" `,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}
		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}

		environmentName, _ := cmd.Flags().GetString(flagnames.Environment)
		if environmentName == "" {
			return errors.New("`environment` is required")
		}

		componentName, _ := cmd.Flags().GetString(flagnames.Component)
		if componentName == "" {
			return errors.New("`component` is required")
		}

		fqdn, _ := cmd.Flags().GetString(flagnames.Alias)
		if fqdn == "" {
			return errors.New("`alias` is required")
		}

		certificate, err := getStringFromFlagValueOrFlagFile(cmd, flagnames.Certificate, flagnames.CertificateFromFile)
		if err != nil {
			return err
		}
		if certificate == "" {
			return errors.New("certificate value cannot be empty")
		}

		privateKey, err := getStringFromFlagValueOrFlagFile(cmd, flagnames.PrivateKey, flagnames.PrivateKeyFromFile)
		if err != nil {
			return err
		}
		if privateKey == "" {
			return errors.New("private key value cannot be empty")
		}

		skipValidation, _ := cmd.Flags().GetBool(flagnames.SkipValidation)
		awaitReconcile, _ := cmd.Flags().GetBool(flagnames.AwaitReconcile)
		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		if awaitReconcile {
			reconciledOk := awaitReconciliation(func() bool {
				return isComponentExternalDNSReconciled(apiClient, *appName, environmentName, componentName, fqdn)
			})

			if !reconciledOk {
				return fmt.Errorf("component was not reconciled within time: either component %s does not exist in the environment %s or external DNS alias %s is not defined",
					componentName, environmentName, fqdn)
			}
		}
		updateExternalDNSTLS := component.NewUpdateComponentExternalDNSTLSParams().
			WithAppName(*appName).
			WithEnvName(environmentName).
			WithComponentName(componentName).
			WithFqdn(fqdn).
			WithTLSData(&models.UpdateExternalDNSTLSRequest{
				Certificate:    &certificate,
				PrivateKey:     &privateKey,
				SkipValidation: skipValidation,
			})

		_, err = apiClient.Component.UpdateComponentExternalDNSTLS(updateExternalDNSTLS, nil)
		return err
	},
}

func isComponentExternalDNSReconciled(apiClient *apiclient.Radixapi, appName, environmentName, componentName, fqdn string) bool {
	getEnvironmentParameters := environment.NewGetEnvironmentParams().
		WithAppName(appName).
		WithEnvName(environmentName)

	env, err := apiClient.Environment.GetEnvironment(getEnvironmentParameters, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	if env.Payload != nil &&
		env.Payload.ActiveDeployment != nil &&
		env.Payload.ActiveDeployment.Components != nil {
		for _, component := range env.Payload.ActiveDeployment.Components {
			if *component.Name == componentName {
				for _, externalDns := range component.ExternalDNS {
					if *externalDns.FQDN == fqdn {
						return true
					}
				}
			}
		}
	}

	return false
}

func init() {
	setExternalDnsTlsCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	setExternalDnsTlsCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment")
	setExternalDnsTlsCmd.Flags().String(flagnames.Component, "", "Name of the component")
	setExternalDnsTlsCmd.Flags().String(flagnames.Alias, "", "External DNS alias to update")
	setExternalDnsTlsCmd.Flags().String(flagnames.Certificate, "", "Certificate (PEM format)")
	setExternalDnsTlsCmd.Flags().String(flagnames.CertificateFromFile, "", "Read certificate (PEM format) from file")
	setExternalDnsTlsCmd.Flags().String(flagnames.PrivateKey, "", "Private key (PEM format)")
	setExternalDnsTlsCmd.Flags().String(flagnames.PrivateKeyFromFile, "", "Read private key (PEM format) from file")
	setExternalDnsTlsCmd.Flags().Bool(flagnames.SkipValidation, false, "Skip validation of certificate and private key")
	setExternalDnsTlsCmd.Flags().Bool(flagnames.AwaitReconcile, true, "Await reconciliation in Radix. Default is true")

	setExternalDnsTlsCmd.MarkFlagsOneRequired(flagnames.Certificate, flagnames.CertificateFromFile)
	setExternalDnsTlsCmd.MarkFlagsMutuallyExclusive(flagnames.Certificate, flagnames.CertificateFromFile)
	setExternalDnsTlsCmd.MarkFlagsOneRequired(flagnames.PrivateKey, flagnames.PrivateKeyFromFile)
	setExternalDnsTlsCmd.MarkFlagsMutuallyExclusive(flagnames.PrivateKey, flagnames.PrivateKeyFromFile)

	_ = setExternalDnsTlsCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = setExternalDnsTlsCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)

	setContextSpecificPersistentFlags(setExternalDnsTlsCmd)
	setCmd.AddCommand(setExternalDnsTlsCmd)
}
