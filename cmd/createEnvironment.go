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

	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const createEnvironmentEnabled = true

// createEnvironmentCmd represents the create environment command
var createEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Create environment",
	Long:  `Creates a Radix environment for the application`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		envName, err := cmd.Flags().GetString("environment")

		if err != nil || appName == nil || *appName == "" {
			return errors.New("Environment name and application name are required fields")
		}

		parameters := environment.NewCreateEnvironmentParams().
			WithAppName(*appName).
			WithEnvName(envName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Environment.CreateEnvironment(parameters, nil)

		println(fmt.Sprintf("%v", err))

		return nil
	},
}

func init() {
	if createEnvironmentEnabled {
		createCmd.AddCommand(createEnvironmentCmd)
		createEnvironmentCmd.Flags().StringP("application", "a", "", "Name of the application namespace")
		createEnvironmentCmd.Flags().StringP("environment", "e", "", "Name of the environment to create")
	}
}
