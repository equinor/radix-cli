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
	"github.com/spf13/cobra"
)

// setComponentTLSSecretCmd represents the setComponentTLSSecretCmd command
var setComponentTLSSecretCmd = &cobra.Command{
	Use:   "tls",
	Short: "Set TLS certificate and private key for DNS external alias",
	Long:  `Set TLS certificate and private key for DNS external alias`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}
		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}

		environmentName, _ := cmd.Flags().GetString("environment")
		if environmentName == "" {
			return errors.New("`environment` is required")
		}

		componentName, _ := cmd.Flags().GetString("component")
		if componentName == "" {
			return errors.New("`component` is required")
		}

		fqdn, _ := cmd.Flags().GetString("alias")
		if fqdn == "" {
			return errors.New("`alias` is required")
		}

		certificate, _ := cmd.Flags().GetString("certificate")
		if certificate == "" {
			return errors.New("`certificate` is required")
		}

		privateKey, _ := cmd.Flags().GetString("private-key")
		if privateKey == "" {
			return errors.New("`private-key` is required")
		}

		skipValidation, _ := cmd.Flags().GetBool("skip-validation")
		awaitReconcile, _ := cmd.Flags().GetBool("await-reconcile")
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
	setCmd.AddCommand(setComponentTLSSecretCmd)
	setComponentTLSSecretCmd.Flags().StringP("application", "a", "", "Name of the application")
	setComponentTLSSecretCmd.Flags().StringP("environment", "e", "", "Name of the environment")
	setComponentTLSSecretCmd.Flags().String("component", "", "Name of the component")
	setComponentTLSSecretCmd.Flags().String("alias", "", "External DNS alias to update TLS for")
	setComponentTLSSecretCmd.Flags().String("certificate", "", "Certificate in PEM format")
	setComponentTLSSecretCmd.Flags().String("private-key", "", "Private key in PEM format")
	setComponentTLSSecretCmd.Flags().Bool("skip-validation", false, "Skip validation of certificate and private key")
	setComponentTLSSecretCmd.Flags().Bool("await-reconcile", true, "Await reconciliation in Radix. Default is true")
	setContextSpecificPersistentFlags(setComponentTLSSecretCmd)
}
