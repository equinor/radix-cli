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
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/component"
	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/client/consumer"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/streaminglog"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// logsEnvironmentCmd represents the followEnvironmentCmd command
var logsEnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Gets logs for all components in an environment",
	Long: `Gets and follows logs for all components in an environment.
	
It may take few seconds to get the log.
	`,
	Example: `# Get logs for all components in an environment. Log lines from different components have different colors
rx get logs environment --application radix-test --environment dev`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == "" {
			return errors.New("application name is required")
		}

		environmentName, _ := cmd.Flags().GetString(flagnames.Environment)
		previousLog, _ := cmd.Flags().GetBool(flagnames.Previous)
		since, _ := cmd.Flags().GetDuration(flagnames.Since)

		if environmentName == "" {
			return errors.New("both `environment` and `component` are required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		return streaminglog.New(
			cmd.OutOrStdout(),
			getComponentReplicasForEnvironment(apiClient, appName, environmentName),
			getComponentLog(apiClient, appName, since, previousLog),
		).StreamLogs(cmd.Context())
	},
}

type ComponentItem struct {
	Component string
	Replica   string
}

func (c ComponentItem) String() string {
	return c.Component + "/" + c.Replica
}

func getComponentLog(apiClient *radixapi.Radixapi, appName string, since time.Duration, previous bool) streaminglog.GetLogFunc[ComponentItem] {
	now := time.Now()
	sinceTime := now.Add(-since)
	sinceStr := strfmt.DateTime(sinceTime)
	previousStr := strconv.FormatBool(previous)

	return func(ctx context.Context, item ComponentItem, print func(text string)) error {
		logParameters := component.NewLogParamsWithContext(ctx)
		logParameters.WithAppName(appName)
		logParameters.WithDeploymentName("irrelevant")
		logParameters.WithComponentName(item.Component)
		logParameters.WithPodName(item.Replica)
		logParameters.WithFollow(pointers.Ptr("true"))
		logParameters.SetSinceTime(&sinceStr)
		logParameters.WithPrevious(&previousStr)

		resp, err := apiClient.Component.Log(logParameters, nil, consumer.NewEventSourceClientOptions(func(event consumer.Event) {
			switch event.Type {
			case "event":
				switch event.Message {
				case "started":
					print("stream started...")
				case "completed":
					print("stream closed.")
				}
			case "data":
				print(event.Message)
			}
		}))
		if err != nil {
			return err
		}

		lines := strings.Split(resp.Payload, "\n")
		for _, line := range lines {
			print(line)
		}
		print("stream closed.")

		return nil
	}
}

func getComponentReplicasForEnvironment(apiClient *radixapi.Radixapi, appName, environmentName string) streaminglog.GetReplicasFunc[ComponentItem] {
	return func() ([]ComponentItem, error) {
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

		componentReplicas := make([]ComponentItem, 0, 50)
		for _, component := range environmentDetails.Payload.ActiveDeployment.Components {
			if component.Name != nil {
				for _, replica := range component.ReplicaList {
					componentReplicas = append(componentReplicas, ComponentItem{
						Component: *component.Name,
						Replica:   *replica.Name,
					})
				}
				//
			}
		}

		return componentReplicas, nil
	}
}

func init() {
	logsCmd.AddCommand(logsEnvironmentCmd)
	logsEnvironmentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application owning the component")
	logsEnvironmentCmd.Flags().StringP(flagnames.Environment, "e", "", "Environment the component runs in")
	logsEnvironmentCmd.Flags().BoolP(flagnames.Previous, "p", false, "If set, print the logs for the previous instances of containers in environment component pods, if they exist")
	logsEnvironmentCmd.Flags().DurationP(flagnames.Since, "s", settings.DeltaRefreshApplication, "If set, start get logs from the specified time, eg. 5m or 12h")

	_ = logsEnvironmentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = logsEnvironmentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	setContextSpecificPersistentFlags(logsEnvironmentCmd)
}
