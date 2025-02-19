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

	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// deleteEnvironmentCmd represents the delete environment command
var deleteEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "delete environment",
	Long:  `deletes an orphaned Radix environment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString(flagnames.Environment)

		if err != nil || appName == "" {
			return errors.New("environment name and application name are required fields")
		}

		cmd.SilenceUsage = true

		parameters := environment.NewDeleteEnvironmentParams().
			WithAppName(appName).
			WithEnvName(envName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Environment.DeleteEnvironment(parameters, nil)
		return err
	},
}

func init() {
	deleteCmd.AddCommand(deleteEnvironmentCmd)
	deleteEnvironmentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	deleteEnvironmentCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment to delete")
	_ = deleteEnvironmentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = deleteEnvironmentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	setContextSpecificPersistentFlags(deleteEnvironmentCmd)
}
