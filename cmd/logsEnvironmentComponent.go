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
	"io"
	"strconv"
	"sync"
	"time"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/component"
	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/client/consumer"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/log"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/go-openapi/strfmt"
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

		_, replicas, err := getReplicasForComponent(apiClient, appName, environmentName, componentName)
		if err != nil {
			return err
		}

		componentReplicas := make(map[string][]string)
		componentReplicas[componentName] = replicas

		return logForComponentReplicas(cmd, apiClient, appName, environmentName, since, componentReplicas, previousLog)
	},
}

func logForComponentReplicas(cmd *cobra.Command, apiClient *radixapi.Radixapi, appName, environmentName string, since time.Duration, componentReplicas map[string][]string, previousLog bool) error {
	previous := strconv.FormatBool(previousLog)
	now := time.Now()
	sinceTime := now.Add(-since)
	sinceDt := strfmt.DateTime(sinceTime)
	wg := sync.WaitGroup{}
	i := 0
	for componentName, replicas := range componentReplicas {
		for _, replica := range replicas {
			i++
			colorIndex := i
			wg.Go(func() {
				logParameters := component.NewLogParamsWithContext(cmd.Context())
				logParameters.WithAppName(appName)
				logParameters.WithDeploymentName("irrelevant")
				logParameters.WithComponentName(componentName)
				logParameters.WithPodName(replica)
				logParameters.WithFollow(pointers.Ptr("true"))
				logParameters.SetSinceTime(&sinceDt)
				logParameters.WithPrevious(&previous)
				_, err := apiClient.Component.Log(logParameters, nil, consumer.CreateEventSourceClientOptions(func(event consumer.Event) {
					switch event.Type {
					case "error":
						log.PrintLine(cmd, "error", fmt.Sprintf("Could not get log for component %s, replica %s: %s", componentName, replica, event.Message), log.Red)
					case "event":
						switch event.Message {
						case "started":
							log.PrintLine(cmd, replica, "stream started...", log.GetColor(colorIndex))
						case "completed":
							log.PrintLine(cmd, replica, "stream closed.", log.GetColor(colorIndex))
						}
					case "data":
						log.PrintLine(cmd, replica, event.Message, log.GetColor(colorIndex))
					}
				}))

				if err != nil && !errors.Is(err, io.EOF) {
					log.PrintLine(cmd, replica, log.Red(fmt.Sprintf("error: Could not get log: %s", err.Error())), log.GetColor(colorIndex))
				}
			})
		}
	}

	// wg.Go(func() {
	// 	ticker := time.NewTicker(time.Second * 15)
	// 	for range ticker.C {
	// 		// update replicas
	// 		for componentName, replicas := range componentReplicas {
	// 			_, newReplicas, err := getReplicasForComponent(apiClient, appName, environmentName, componentName)
	// 			if err != nil {
	// 				logrus.Infof("Failed to get replicas for component %s. %v", componentName, err)
	// 				continue
	// 			}

	// 			addedReplicas := slice.Difference(newReplicas, replicas)
	// 			if len(addedReplicas) > 0 {
	// 				log.PrintLine(cmd, "info", fmt.Sprintf("New replicas for component %s: %v", componentName, addedReplicas), log.Yellow)
	// 			}
	// 			componentReplicas[componentName] = newReplicas

	// 			for _, replica := range addedReplicas {
	// 				i++
	// 				wg.Go(func() {
	// 					logParameters := component.NewLogParamsWithContext(cmd.Context())
	// 					logParameters.WithAppName(appName)
	// 					logParameters.WithDeploymentName("irrelevant")
	// 					logParameters.WithComponentName(componentName)
	// 					logParameters.WithPodName(replica)
	// 					logParameters.WithFollow(pointers.Ptr("true"))
	// 					logParameters.SetSinceTime(&sinceDt)
	// 					logParameters.WithPrevious(&previous)

	// 					_, err := apiClient.Component.Log(logParameters, nil, func(co *runtime.ClientOperation) {
	// 						co.Reader = createLogReader(cmd, replica, i)
	// 					})

	// 					if err != nil {
	// 						log.PrintLine(cmd, "error", fmt.Sprintf("Could not get log for component %s, replica %s: %v", componentName, replica, err), log.Red)
	// 					}
	// 				})
	// 			}
	// 		}
	// 	}
	// })

	wg.Wait()
	return nil
}

func getReplicasForComponent(apiClient *radixapi.Radixapi, appName, environmentName, componentName string) (*string, []string, error) {
	// Get active deployment
	start := time.Now()
	environmentParams := environment.NewGetEnvironmentParams()
	environmentParams.SetAppName(appName)
	environmentParams.SetEnvName(environmentName)
	environmentDetails, err := apiClient.Environment.GetEnvironment(environmentParams, nil)
	duration := time.Since(start)
	fmt.Printf("Fetched environment details for %s in %s in %v\n", environmentName, appName, duration)

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
