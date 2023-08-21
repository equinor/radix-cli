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
	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// startApplicationCmd represents the start application command
var startApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Start an application",
	Long: `Start an application
  - Pulls new images from image hub in radix configuration
  - Starts the application containers using up to date images`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if err != nil || appName == nil || *appName == "" {
			return errors.New("application name is required fields")
		}

		cmd.SilenceUsage = true

		parameters := application.NewStartApplicationParams().
			WithAppName(*appName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Application.StartApplication(parameters, nil)
		return err
	},
}

func init() {
	startCmd.AddCommand(startApplicationCmd)
	startApplicationCmd.Flags().StringP("application", "a", "", "Name of the application namespace")
	setContextSpecificPersistentFlags(startApplicationCmd)
}
