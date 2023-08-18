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
	"github.com/equinor/radix-cli/pkg/utils/json"

	"github.com/equinor/radix-cli/generated-client/client/deployment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// getDeploymentCmd represents the getDeploymentCmd command
var getDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Will get a deployment for a given application",
	Long: `Will get a deployment for a given application

Examples:
  # Get a deployment deployment-abc for an application radix-test 
  rx get deployment --application radix-test --deployment deployment-abc
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}
		deploymentName, err := cmd.Flags().GetString("deployment")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" || deploymentName == "" {
			return errors.New("application and deployment names are required fields")
		}

		cmd.SilenceUsage = true

		params := deployment.NewGetDeploymentParams()
		params.WithAppName(*appName)
		if deploymentName != "" {
			params.WithDeploymentName(deploymentName)
		}
		resp, err := apiClient.Deployment.GetDeployment(params, nil)
		if err != nil {
			return err
		}
		prettyJSON, err := json.Pretty(resp.Payload)
		if err != nil {
			return err
		}
		fmt.Println(*prettyJSON)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getDeploymentCmd)
	getDeploymentCmd.Flags().StringP("application", "a", "", "Name of the application")
	getDeploymentCmd.Flags().StringP("deployment", "d", "", "Optional, name of a deployment")
	setContextSpecificPersistentFlags(getDeploymentCmd)
}
