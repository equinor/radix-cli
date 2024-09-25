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
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// createEnvironmentCmd represents the create environment command
var createEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Create environment",
	Long:  `Creates a Radix environment for the application`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString(flagnames.Environment)

		if err != nil || appName == nil || *appName == "" {
			return errors.New("environment name and application name are required fields")
		}

		cmd.SilenceUsage = true

		parameters := environment.NewCreateEnvironmentParams().
			WithAppName(*appName).
			WithEnvName(envName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Environment.CreateEnvironment(parameters, nil)
		return err
	},
}

func init() {
	createCmd.AddCommand(createEnvironmentCmd)
	createEnvironmentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	createEnvironmentCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment to create")
	_ = getApplicationCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(createEnvironmentCmd)
}
