// Copyright © 2022
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

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const restartApplicationEnabled = true

// restartApplicationCmd represents the restart application command
var restartApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Restart an application",
	Long: `Restart an application
  - Stops running the application containers
  - Pulls new images from image hub in radix configuration
  - Starts the application containers again using up to date images`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if err != nil || appName == nil || *appName == "" {
			return errors.New("application name is required fields")
		}

		parameters := application.NewRestartApplicationParams().
			WithAppName(*appName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Application.RestartApplication(parameters, nil)

		println(fmt.Sprintf("%v", err))

		return nil
	},
}

func init() {
	if restartApplicationEnabled {
		restartCmd.AddCommand(restartApplicationCmd)
		restartApplicationCmd.Flags().StringP("application", "a", "", "Name of the application namespace")
	}
}
