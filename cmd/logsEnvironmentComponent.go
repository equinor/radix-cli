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

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/streaminglog"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

// logsEnvironmentComponentCmd represents the logsEnvironmentComponentCmd command
var logsEnvironmentComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Gets logs for a specific components in an environment",
	Long: `Gets and follows logs for a component in an environment.

It may take few seconds to get the log.

Examples:
  # Get logs for a component 
  rx get logs component --application radix-test --environment dev --component web-app

  # Get logs for a component previous (terminated or restarted) container
  rx get logs component --application radix-test --environment dev --component web-app --previous

  # Short version of get logs for a component previous (terminated or restarted) container
  rx get logs component -a radix-test -e dev --component web-app -p
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == "" {
			return errors.New("application name is required")
		}

		environmentName, _ := cmd.Flags().GetString(flagnames.Environment)
		componentName, _ := cmd.Flags().GetString(flagnames.Component)
		previousLog, _ := cmd.Flags().GetBool(flagnames.Previous)
		since, _ := cmd.Flags().GetDuration(flagnames.Since)

		if environmentName == "" || componentName == "" {
			return errors.New("both `environment` and `component` are required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		var getReplicas = func() (map[string][]string, error) {
			_, replicas, err := getReplicasForComponent(apiClient, appName, environmentName, componentName)
			componentReplicas := make(map[string][]string)
			componentReplicas[componentName] = replicas
			return componentReplicas, err
		}
		return streaminglog.NewStreamingReplicas(apiClient, cmd.OutOrStdout(), appName, since, previousLog, getReplicas).StreamLogs(cmd.Context())
	},
}

func getReplicasForComponent(apiClient *radixapi.Radixapi, appName, environmentName, componentName string) (*string, []string, error) {
	// Get active deployment
	environmentParams := environment.NewGetEnvironmentParams()
	environmentParams.SetAppName(appName)
	environmentParams.SetEnvName(environmentName)
	environmentDetails, err := apiClient.Environment.GetEnvironment(environmentParams, nil)

	if err != nil {
		return nil, nil, err
	}

	var deploymentName string
	if environmentDetails == nil || environmentDetails.Payload.ActiveDeployment == nil {
		return nil, nil, errors.New("active deployment was not found in environment")
	}

	var replicas []string
	deploymentName = *environmentDetails.Payload.ActiveDeployment.Name
	for _, comp := range environmentDetails.Payload.ActiveDeployment.Components {
		if comp.Name != nil &&
			*comp.Name == componentName {
			replicas = slice.Reduce(comp.ReplicaList, make([]string, 0), func(acc []string, replica *models.ReplicaSummary) []string {
				return append(acc, *replica.Name)
			})
			break
		}
	}

	return &deploymentName, replicas, nil
}

func init() {
	logsCmd.AddCommand(logsEnvironmentComponentCmd)
	logsEnvironmentComponentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application owning the component")
	logsEnvironmentComponentCmd.Flags().StringP(flagnames.Environment, "e", "", "Environment the component runs in")
	logsEnvironmentComponentCmd.Flags().String(flagnames.Component, "", "The component to follow")
	logsEnvironmentComponentCmd.Flags().BoolP(flagnames.Previous, "p", false, "If set, print the logs for the previous instance of the container in a component pod, if it exists")
	logsEnvironmentComponentCmd.Flags().DurationP(flagnames.Since, "s", settings.DeltaRefreshApplication, "If set, start get logs from the specified time, eg. 5m or 12h")

	_ = logsEnvironmentComponentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = logsEnvironmentComponentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = logsEnvironmentComponentCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	setContextSpecificPersistentFlags(logsEnvironmentComponentCmd)
}
