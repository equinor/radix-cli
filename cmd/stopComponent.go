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

	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const stopComponentEnabled = true

// stopComponentCmd represents the stop component command
var stopComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Stop a component",
	Long: `Stop a component
  - Stops the component running container`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString("environment")

		if err != nil || appName == nil || *appName == "" || envName == "" {
			return errors.New("environment name and application name are required fields")
		}

		cmpName, err := cmd.Flags().GetString("component")
		if err != nil {
			return errors.New("component name is a required field")
		}

		parameters := component.NewStopComponentParams().
			WithAppName(*appName).
			WithEnvName(envName).
			WithComponentName(cmpName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Component.StopComponent(parameters, nil)

		println(fmt.Sprintf("%v", err))

		return nil
	},
}

func init() {
	if stopComponentEnabled {
		stopCmd.AddCommand(stopComponentCmd)
		stopComponentCmd.Flags().StringP("application", "a", "", "Name of the application namespace")
		stopComponentCmd.Flags().StringP("environment", "e", "", "Name of the environment of the application")
		stopComponentCmd.Flags().StringP("component", "n", "", "Name of the component to stop")
	}
}
