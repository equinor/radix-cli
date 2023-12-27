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

// setEnvironmentVariableCmd represents the setEnvironmentVariableCmd command
var setEnvironmentVariableCmd = &cobra.Command{
	Use:   "environment-variable",
	Short: "Will set an environment variable",
	Long:  `Will set an environment variable`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}

		variableName, _ := cmd.Flags().GetString("variable")
		variableValue, _ := cmd.Flags().GetString("value")

		if variableName == "" {
			return errors.New("variable is required")
		}

		if variableValue == "" {
			return errors.New("value is required")
		}

		environmentName, _ := cmd.Flags().GetString("environment")

		if environmentName == "" {
			return errors.New("`environment` is required")
		}

		componentName, _ := cmd.Flags().GetString("component")
		if componentName == "" {
			return errors.New("`component` is required")
		}

		awaitReconcile, _ := cmd.Flags().GetBool("await-reconcile")

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		if awaitReconcile {
			reconciledOk := awaitReconciliation(func() bool {
				return isComponentVariableReconciled(apiClient, *appName, environmentName, componentName, variableName)
			})

			if !reconciledOk {
				return fmt.Errorf("component was not reconciled within time: either component %s does not exist in the environment %s or the component does not have a variable %s",
					componentName, environmentName, variableName)
			}
		}

		componentVariable := models.EnvVarParameter{}
		componentVariable.Name = &variableName
		componentVariable.Value = &variableValue

		changeComponentVariableParameters := component.NewChangeEnvVarParams()
		changeComponentVariableParameters.SetAppName(*appName)
		changeComponentVariableParameters.SetEnvName(environmentName)
		changeComponentVariableParameters.SetComponentName(componentName)
		changeComponentVariableParameters.SetEnvVarParameter([]*models.EnvVarParameter{&componentVariable})

		_, err = apiClient.Component.ChangeEnvVar(changeComponentVariableParameters, nil)
		return err
	},
}

func isComponentVariableReconciled(apiClient *apiclient.Radixapi, appName, environmentName, componentName, variableName string) bool {
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
				for _, variable := range component.Variables {
					if variable == variableName {
						return true
					}
				}
			}
		}

	}

	return false
}

func init() {
	setCmd.AddCommand(setEnvironmentVariableCmd)
	setEnvironmentVariableCmd.Flags().StringP("application", "a", "", "Name of the application to set variable for")
	setEnvironmentVariableCmd.Flags().StringP("environment", "e", "", "Environment to set variable in")
	setEnvironmentVariableCmd.Flags().String("component", "", "Component to set the variable for")
	setEnvironmentVariableCmd.Flags().StringP("variable", "", "", "Name of the variable to set")
	setEnvironmentVariableCmd.Flags().StringP("value", "v", "", "Value of the variable to set")
	setEnvironmentVariableCmd.Flags().Bool("await-reconcile", true, "Await reconciliation in Radix. Default is true")
	setContextSpecificPersistentFlags(setEnvironmentVariableCmd)
}
