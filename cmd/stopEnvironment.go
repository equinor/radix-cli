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

// stopEnvironmentCmd represents the stop environment command
var stopEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Stop an environment",
	Long: `Stop an environment
  - Stops the environment components running containers`,
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

		parameters := environment.NewStopEnvironmentParams().
			WithAppName(*appName).
			WithEnvName(envName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Environment.StopEnvironment(parameters, nil)
		return err
	},
}

func init() {
	stopCmd.AddCommand(stopEnvironmentCmd)
	stopEnvironmentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	stopEnvironmentCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")

	_ = stopEnvironmentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = stopEnvironmentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	setContextSpecificPersistentFlags(stopEnvironmentCmd)
}
