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

	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const restartEnvironmentEnabled = true

// restartEnvironmentCmd represents the restart environment command
var restartEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Restart an environment",
	Long: `Restart an environment
  - Starts the environment's containers, using up to date images
  - Stops the application environment's old containers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString("environment")

		if err != nil || appName == nil || *appName == "" || envName == "" {
			return errors.New("environment name and application name are required fields")
		}

		parameters := environment.NewRestartEnvironmentParams().
			WithAppName(*appName).
			WithEnvName(envName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Environment.RestartEnvironment(parameters, nil)

		println(fmt.Sprintf("%v", err))

		return nil
	},
}

func init() {
	if restartEnvironmentEnabled {
		restartCmd.AddCommand(restartEnvironmentCmd)
		restartEnvironmentCmd.Flags().StringP("application", "a", "", "Name of the application namespace")
		restartEnvironmentCmd.Flags().StringP("environment", "e", "", "Name of the environment of the application")
		setContextSpecificPersistentFlags(restartEnvironmentCmd)
	}
}
