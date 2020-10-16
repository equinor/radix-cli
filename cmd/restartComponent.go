// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const restartComponentEnabled = true

// restartComponentCmd represents the restart component command
var restartComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Restart a component",
	Long: `Restart a component
  - Stops running the component container
  - Pulls new image from image hub in radix configuration
  - Starts the container again using up to date image`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString("environment")

		if err != nil || appName == nil || *appName == "" {
			return errors.New("Environment name and application name are required fields")
		}

		cmpName, err := cmd.Flags().GetString("component")
		if err != nil {
			return errors.New("Component name is a required field")
		}

		parameters := component.NewRestartComponentParams().
			WithAppName(*appName).
			WithEnvName(envName).
			WithComponentName(cmpName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Component.RestartComponent(parameters, nil)

		println(fmt.Sprintf("%v", err))

		return nil
	},
}

func init() {
	if restartComponentEnabled {
		restartCmd.AddCommand(restartComponentCmd)
		restartComponentCmd.Flags().StringP("application", "a", "", "Name of the application namespace")
		restartComponentCmd.Flags().StringP("environment", "e", "", "Name of the environment of the application")
		restartComponentCmd.Flags().StringP("component", "n", "", "Name of the component to restart")
	}
}
