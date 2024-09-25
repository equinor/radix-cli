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

	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// restartEnvironmentCmd represents the restart environment command
var restartEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Restart an environment",
	Long: `Restart an environment
  - Starts the environment's containers, using up to date images
  - Stops the application environment's old containers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString(flagnames.Environment)

		if err != nil || appName == nil || *appName == "" || envName == "" {
			return errors.New("environment name and application name are required fields")
		}

		cmd.SilenceUsage = true

		parameters := environment.NewRestartEnvironmentParams().
			WithAppName(*appName).
			WithEnvName(envName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Environment.RestartEnvironment(parameters, nil)
		return err
	},
}

func init() {
	restartCmd.AddCommand(restartEnvironmentCmd)
	restartEnvironmentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	restartEnvironmentCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")

	_ = restartComponentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = restartComponentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(restartEnvironmentCmd)
}
