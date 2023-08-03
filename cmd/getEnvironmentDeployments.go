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

	"github.com/equinor/radix-cli/generated-client/client/deployment"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/utils/json"
	"github.com/spf13/cobra"
)

// getEnvironmentDeploymentsCmd represents the getEnvironmentDeploymentsCmd command
var getEnvironmentDeploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Will get deployments for a given application and environment",
	Long: `Will get deployments for a given application and environment

Examples:
  # Get list of deployments for an application radix-test 
  rx get deployments --application radix-test

  # Get list of deployments for an application radix-test and environment test
  rx get deployments --application radix-test --environment test
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}
		if appName == nil || *appName == "" {
			return errors.New("application name is required field")
		}
		envName, err := cmd.Flags().GetString("environment")
		if err != nil {
			return err
		}
		if envName != "" {
			return getDeploymentForEnvironment(cmd, *appName, envName)
		}
		return getDeploymentForAllEnvironments(cmd, *appName)
	},
}

func getDeploymentForAllEnvironments(cmd *cobra.Command, appName string) error {
	params := deployment.NewGetDeploymentParams()
	params.WithAppName(appName)
	apiClient, err := client.GetForCommand(cmd)
	if err != nil {
		return err
	}
	resp, err := apiClient.Deployment.GetDeployment(params, nil)
	if err != nil {
		println(fmt.Sprintf("%v", err))
		return err
	}
	prettyJSON, err := json.Pretty(resp.Payload)
	if err != nil {
		println(fmt.Sprintf("%v", err))
		return err
	}
	fmt.Println(*prettyJSON)
	return nil
}

func getDeploymentForEnvironment(cmd *cobra.Command, appName, envName string) error {
	params := environment.NewGetApplicationEnvironmentDeploymentsParams()
	params.WithAppName(appName)
	params.WithEnvName(envName)
	apiClient, err := client.GetForCommand(cmd)
	if err != nil {
		return err
	}
	resp, err := apiClient.Environment.GetApplicationEnvironmentDeployments(params, nil)
	if err != nil {
		println(fmt.Sprintf("%v", err))
		return err
	}
	prettyJSON, err := json.Pretty(resp.Payload)
	if err != nil {
		println(fmt.Sprintf("%v", err))
		return err
	}
	fmt.Println(*prettyJSON)
	return nil
}

func init() {
	if getBranchEnvironmentEnabled {
		getCmd.AddCommand(getEnvironmentDeploymentsCmd)
		getEnvironmentDeploymentsCmd.Flags().StringP("application", "a", "", "Name of the application")
		getEnvironmentDeploymentsCmd.Flags().StringP("environment", "e", "", "Optional, name of the environment")
	}
}
