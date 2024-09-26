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
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// setEnvironmentSecretCmd represents the setEnvironmentSecretCmd command
var setEnvironmentSecretCmd = &cobra.Command{
	Use:   "environment-secret",
	Short: "Will set an environment secret",
	Long:  `Will set an environment secret`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == "" {
			return errors.New("application name is required")
		}

		secretName, _ := cmd.Flags().GetString(flagnames.Secret)
		secretValue, _ := cmd.Flags().GetString(flagnames.Value)

		if secretName == "" {
			return errors.New("secret is required")
		}

		if secretValue == "" {
			return errors.New("value is required")
		}

		environmentName, _ := cmd.Flags().GetString(flagnames.Environment)

		if environmentName == "" {
			return errors.New("`environment` is required")
		}

		component, _ := cmd.Flags().GetString(flagnames.Component)
		if component == "" {
			return errors.New("`component` is required")
		}

		awaitReconcile, _ := cmd.Flags().GetBool(flagnames.AwaitReconcile)

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		if awaitReconcile {
			reconciledOk := awaitReconciliation(func() bool {
				return isComponentSecretReconciled(apiClient, appName, environmentName, component, secretName)
			})

			if !reconciledOk {
				return fmt.Errorf("component was not reconciled within time: either component %s does not exist in the environment %s or the component has not secret %s",
					component, environmentName, secretName)
			}
		}

		componentSecret := models.SecretParameters{}
		componentSecret.SecretValue = &secretValue

		changeComponentSecretParameters := environment.NewChangeComponentSecretParams()
		changeComponentSecretParameters.SetAppName(appName)
		changeComponentSecretParameters.SetEnvName(environmentName)
		changeComponentSecretParameters.SetComponentName(component)
		changeComponentSecretParameters.SetSecretName(secretName)
		changeComponentSecretParameters.SetComponentSecret(&componentSecret)

		_, err = apiClient.Environment.ChangeComponentSecret(changeComponentSecretParameters, nil)
		return err
	},
}

func isComponentSecretReconciled(apiClient *apiclient.Radixapi, appName, environmentName, componentName, secretName string) bool {
	getEnvironmentParameters := environment.NewGetEnvironmentParams()
	getEnvironmentParameters.SetAppName(appName)
	getEnvironmentParameters.SetEnvName(environmentName)

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
	setCmd.AddCommand(setEnvironmentSecretCmd)
	setEnvironmentSecretCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to set secret for")
	setEnvironmentSecretCmd.Flags().StringP(flagnames.Environment, "e", "", "Environment to set secret in")
	setEnvironmentSecretCmd.Flags().String(flagnames.Component, "", "Component to set the secret for")
	setEnvironmentSecretCmd.Flags().StringP(flagnames.Secret, "s", "", "Name of the secret to set")
	setEnvironmentSecretCmd.Flags().StringP(flagnames.Value, "v", "", "Value of the secret to set")
	setEnvironmentSecretCmd.Flags().Bool(flagnames.AwaitReconcile, true, "Await reconciliation in Radix. Default is true")

	_ = setEnvironmentSecretCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = setEnvironmentSecretCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = setEnvironmentSecretCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	_ = setEnvironmentSecretCmd.RegisterFlagCompletionFunc(flagnames.Secret, completion.SecretCompletion)
	setContextSpecificPersistentFlags(setEnvironmentSecretCmd)
}
