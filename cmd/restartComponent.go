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

	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// restartComponentCmd represents the restart component command
var restartComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Restart a component",
	Long: `Restart a component
  - Starts the component's container, using up to date image
  - Stops the application component's old containers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString(flagnames.Environment)

		if err != nil || appName == nil || *appName == "" {
			return errors.New("environment name and application name are required fields")
		}

		cmpName, err := cmd.Flags().GetString(flagnames.Component)
		if err != nil {
			return errors.New("component name is a required field")
		}

		cmd.SilenceUsage = true

		parameters := component.NewRestartComponentParams().
			WithAppName(*appName).
			WithEnvName(envName).
			WithComponentName(cmpName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Component.RestartComponent(parameters, nil)
		return err
	},
}

func init() {
	restartCmd.AddCommand(restartComponentCmd)
	restartComponentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	restartComponentCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")
	restartComponentCmd.Flags().StringP(flagnames.Component, "n", "", "Name of the component to restart")

	_ = getApplicationCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(restartComponentCmd)
}
