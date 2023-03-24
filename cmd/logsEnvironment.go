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

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const logsEnvironmentEnabled = true

// logsEnvironmentCmd represents the followEnvironmentCmd command
var logsEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Get logs of all components in environment",
	Long: `Will get and follow logs of all components in an environment.

It may take few seconds to get the log.

Example:
  # Get logs for all components in an environment. Log lines from different components have different colors
  rx get logs environment --application radix-test --environment dev
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}

		environmentName, _ := cmd.Flags().GetString("environment")
		previousLog, _ := cmd.Flags().GetBool("previous")

		if environmentName == "" {
			return errors.New("both `environment` and `component` are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		componentReplicas, err := getComponentReplicasForEnvironment(apiClient, *appName, environmentName)
		if err != nil {
			return err
		}

		err = logForComponentReplicas(cmd, apiClient, *appName, environmentName, componentReplicas, previousLog)
		return err

	},
}

func getComponentReplicasForEnvironment(apiClient *apiclient.Radixapi, appName, environmentName string) (map[string][]string, error) {
	// Get active deployment
	environmentParams := environment.NewGetEnvironmentParams()
	environmentParams.SetAppName(appName)
	environmentParams.SetEnvName(environmentName)
	environmentDetails, err := apiClient.Environment.GetEnvironment(environmentParams, nil)

	if err != nil {
		return nil, err
	}

	if environmentDetails == nil || environmentDetails.Payload.ActiveDeployment == nil {
		return nil, errors.New("active deployment was not found in environment")
	}

	componentReplicas := make(map[string][]string)
	for _, component := range environmentDetails.Payload.ActiveDeployment.Components {
		if component.Name != nil {
			componentReplicas[*component.Name] = component.Replicas
		}
	}

	return componentReplicas, nil
}

func init() {
	if logsEnvironmentEnabled {
		logsCmd.AddCommand(logsEnvironmentCmd)

		logsEnvironmentCmd.Flags().StringP("application", "a", "", "Name of the application owning the component")
		logsEnvironmentCmd.Flags().StringP("environment", "e", "", "Environment the component runs in")
		logsEnvironmentCmd.Flags().BoolP("previous", "p", false, "If set, print the logs for the previous instances of containers in environment component pods, if they exist")
	}
}
