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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startComponentCmd represents the start component command
var startComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Start a component",
	Long: `Start a component

Deprecated: Use 'rx scale component --reset' instead

Resets a manully scaled component to resume normal operations again.`,
	Deprecated: " Use 'rx scale component --reset' instead. Will be removed after September 2025",
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

		parameters := component.NewStartComponentParams().
			WithAppName(appName).
			WithEnvName(envName).
			WithComponentName(cmpName)

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Component.StartComponent(parameters, nil)
		if err != nil {
			return err
		}

		logrus.Infof("%s Successfully reset to normal scaling", cmpName)
		return nil
	},
}

func init() {
	startCmd.AddCommand(startComponentCmd)
	startComponentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	startComponentCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")
	startComponentCmd.Flags().String(flagnames.Component, "", "Name of the component to start")
	_ = startComponentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = startComponentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = startComponentCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	setContextSpecificPersistentFlags(startComponentCmd)
}
