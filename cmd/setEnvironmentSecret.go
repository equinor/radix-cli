// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	globalSettings "github.com/equinor/radix-cli/pkg/settings"
	"github.com/spf13/cobra"
)

const (
	applicationOption = "application"
	environmentOption = "environment"
	componentOption   = "component"
	secretOption      = "secret"
	valueOption       = "value"
)

// setEnvironmentSecretCmd represents the setEnvironmentSecretCmd command
var setEnvironmentSecretCmd = &cobra.Command{
	Use:   "environment-secret",
	Short: "Will set an environment secret",
	Long:  `Will set an environment secret`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, applicationOption)
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("Application name is required")
		}

		secretName, _ := cmd.Flags().GetString(secretOption)
		secretValue, _ := cmd.Flags().GetString(valueOption)

		if secretName == "" {
			return errors.New("Secret is required")
		}

		if secretValue == "" {
			return errors.New("Value is required")
		}

		environmentName, _ := cmd.Flags().GetString(environmentOption)

		if environmentName == "" {
			return errors.New("`environment` is required")
		}

		component, _ := cmd.Flags().GetString(componentOption)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		awaitReconcile, _ := cmd.Flags().GetBool(globalSettings.AwaitReconcileOption)
		if awaitReconcile {
			reconciledOk := awaitReconciliation(func() bool {
				return isComponentSecretReconciled(apiClient, *appName, environmentName, component, secretName)
			})

			if !reconciledOk {
				return errors.New("Component was not reconciled within time")
			}
		}

		componentSecret := models.SecretParameters{}
		componentSecret.SecretValue = &secretValue

		changeComponentSecretParameters := environment.NewChangeEnvironmentComponentSecretParams()
		changeComponentSecretParameters.SetAppName(*appName)
		changeComponentSecretParameters.SetEnvName(environmentName)
		changeComponentSecretParameters.SetComponentName(component)
		changeComponentSecretParameters.SetSecretName(secretName)
		changeComponentSecretParameters.SetComponentSecret(&componentSecret)

		_, err = apiClient.Environment.ChangeEnvironmentComponentSecret(changeComponentSecretParameters, nil)
		if err != nil {
			return err
		}

		return nil
	},
}

func isComponentSecretReconciled(apiClient *apiclient.Radixapi, appName, environmentName, componentName, secretName string) bool {
	getEnvironmentParameters := environment.NewGetEnvironmentParams()
	getEnvironmentParameters.SetAppName(appName)
	getEnvironmentParameters.SetEnvName(environmentName)

	environment, err := apiClient.Environment.GetEnvironment(getEnvironmentParameters, nil)
	if err != nil {
		return false
	}

	if environment.Payload != nil &&
		environment.Payload.ActiveDeployment != nil &&
		environment.Payload.ActiveDeployment.Components != nil {
		for _, component := range environment.Payload.ActiveDeployment.Components {
			if *component.Name == componentName {
				for _, secret := range component.Secrets {
					if secret == secretName {
						return true
					}
				}
			}
		}

	}

	return false
}

func init() {
	setEnvironmentSecretCmd.Flags().StringP(applicationOption, "a", "", "Name of the application to set secret for")
	setEnvironmentSecretCmd.Flags().StringP(environmentOption, "e", "", "Environment to set secret in")
	setEnvironmentSecretCmd.Flags().String(componentOption, "", "Component to set the secret for")
	setEnvironmentSecretCmd.Flags().StringP(secretOption, "s", "", "Name of the secret to set")
	setEnvironmentSecretCmd.Flags().StringP(valueOption, "v", "", "Value of the secret to set")
}
