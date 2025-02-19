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

	"github.com/equinor/radix-cli/generated/radixapi/client/component"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// stopComponentCmd represents the stop component command
var stopComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Stop a component",
	Long: `Stop a component
  - Stops the component running container`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString(flagnames.Environment)

		if err != nil || appName == "" || envName == "" {
			return errors.New("environment name and application name are required fields")
		}

		cmpName, err := cmd.Flags().GetString(flagnames.Component)
		if err != nil {
			return errors.New("component name is a required field")
		}

		cmd.SilenceUsage = true

		parameters := component.NewStopComponentParams().
			WithAppName(appName).
			WithEnvName(envName).
			WithComponentName(cmpName)

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Component.StopComponent(parameters, nil)
		return err
	},
}

func init() {
	stopCmd.AddCommand(stopComponentCmd)
	stopComponentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	stopComponentCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")
	stopComponentCmd.Flags().StringP(flagnames.Component, "n", "", "Name of the component to stop")
	_ = stopComponentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = stopComponentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = stopComponentCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	setContextSpecificPersistentFlags(stopComponentCmd)
}
