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

	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// setEnvironmentSecretCmd represents the setEnvironmentSecretCmd command
var setEnvironmentSecretCmd = &cobra.Command{
	Use:   "environment-secret",
	Short: "Will set an environment secret",
	Long:  `Will set an environment secret`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("Application name is required")
		}

		secretName, _ := cmd.Flags().GetString("secret")
		secretValue, _ := cmd.Flags().GetString("value")

		if secretName == "" {
			return errors.New("Secret is required")
		}

		if secretValue == "" {
			return errors.New("Value is required")
		}

		environmentName, _ := cmd.Flags().GetString("environment")
		branch, _ := cmd.Flags().GetString("branch")

		if (environmentName != "" && branch != "") || (environmentName == "" && branch == "") {
			return errors.New("Either `environment` or `branch` is required, but both cannot be provided at the same time")
		}

		if branch != "" {
			environmentBranch, err := getEnvironmentFromConfig(cmd, branch)
			if err != nil {
				return err
			}

			environmentName = *environmentBranch
		}

		component, _ := cmd.Flags().GetString("component")

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
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

func init() {
	setEnvironmentSecretCmd.Flags().StringP("application", "a", "", "Name of the application to set secret for")
	setEnvironmentSecretCmd.Flags().StringP("environment", "e", "", "Environment to set secret in")
	setEnvironmentSecretCmd.Flags().StringP("branch", "b", "", "Branch of the repository. Can be used together with --from-config to get the environment")
	setEnvironmentSecretCmd.Flags().String("component", "", "Component to set the secret for")
	setEnvironmentSecretCmd.Flags().StringP("secret", "s", "", "Name of the secret to set")
	setEnvironmentSecretCmd.Flags().StringP("value", "v", "", "Value of the secret to set")
}
