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

	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// stopApplicationCmd represents the stop application command
var stopApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Stop an application",
	Long: `Stop an application
  - Stops the application components running containers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if err != nil || appName == "" {
			return errors.New("application name is required fields")
		}

		cmd.SilenceUsage = true

		parameters := application.NewStopApplicationParams().
			WithAppName(appName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Application.StopApplication(parameters, nil)
		return err
	},
}

func init() {
	stopCmd.AddCommand(stopApplicationCmd)
	stopApplicationCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	_ = stopApplicationCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(stopApplicationCmd)
}
