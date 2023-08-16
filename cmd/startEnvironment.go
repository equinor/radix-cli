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

const startEnvironmentEnabled = true

// startEnvironmentCmd represents the start environment command
var startEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Start an environment",
	Long: `Start an environment
  - Pulls new images from image hub in radix configuration
  - Starts the environment containers using up to date images`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString("environment")

		if err != nil || appName == nil || *appName == "" || envName == "" {
			return errors.New("environment name and application name are required fields")
		}

		parameters := environment.NewStartEnvironmentParams().
			WithAppName(*appName).
			WithEnvName(envName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Environment.StartEnvironment(parameters, nil)

		println(fmt.Sprintf("%v", err))

		return nil
	},
}

func init() {
	if startEnvironmentEnabled {
		startCmd.AddCommand(startEnvironmentCmd)
		startEnvironmentCmd.Flags().StringP("application", "a", "", "Name of the application namespace")
		startEnvironmentCmd.Flags().StringP("environment", "e", "", "Name of the environment of the application")
		setContextSpecificPersistentFlags(startEnvironmentCmd)
	}
}
