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
	"strings"
	"time"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/log"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// logsEnvironmentComponentCmd represents the logsEnvironmentComponentCmd command
var logsEnvironmentComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Get logs of specific components in environment",
	Long:  `Will get and follow logs of component in an environment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}

		environmentName, _ := cmd.Flags().GetString("environment")
		componentName, _ := cmd.Flags().GetString("component")

		if environmentName == "" || componentName == "" {
			return errors.New("both `environment` and `component` are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, replicas, err := getReplicasForComponent(apiClient, *appName, environmentName, componentName)
		if err != nil {
			return err
		}

		componentReplicas := make(map[string][]string)
		componentReplicas[componentName] = replicas

		err = logForComponentReplicas(cmd, apiClient, *appName, environmentName, componentReplicas)
		return err

	},
}

func logForComponentReplicas(cmd *cobra.Command, apiClient *apiclient.Radixapi, appName, environmentName string, componentReplicas map[string][]string) error {
	refreshLog := time.Tick(settings.DeltaRefreshApplication)

	// Somtimes, even though we get delta, the log is the same as previous
	previousLogForReplica := make(map[string]string)

	for range refreshLog {
		i := 0
		for componentName, replicas := range componentReplicas {
			for _, replica := range replicas {
				now := time.Now()
				sinceTime := now.Add(-settings.DeltaRefreshApplication)
				since := strfmt.DateTime(sinceTime)

				logParameters := component.NewLogParams()
				logParameters.WithAppName(appName)
				logParameters.WithDeploymentName("irrelevant")
				logParameters.WithComponentName(componentName)
				logParameters.WithPodName(replica)
				logParameters.SetSinceTime(&since)

				logData, err := apiClient.Component.Log(logParameters, nil)
				if err != nil {
					// Replicas may have died
					_, newReplicas, err := getReplicasForComponent(apiClient, appName, environmentName, componentName)
					if err != nil {
						return err
					}

					componentReplicas[componentName] = newReplicas
					break

				} else {
					// Somtimes, even though we get delta, the log is the same as previous
					if len(logData.Payload) > 0 && !strings.EqualFold(logData.Payload, previousLogForReplica[replica]) {
						logLines := strings.Split(strings.Replace(strings.TrimRight(logData.Payload, "\r\n"), "\r\n", "\n", -1), "\n")
						if len(logLines) > 0 {
							log.PrintLines(cmd, replica, []string{}, logLines, log.GetColor(i))
							previousLogForReplica[replica] = logData.Payload
						}
					}
				}

				i++
			}
		}
	}
	return nil
}

func getReplicasForComponent(apiClient *apiclient.Radixapi, appName, environmentName, componentName string) (*string, []string, error) {
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
	deploymentName = environmentDetails.Payload.ActiveDeployment.Name
	for _, component := range environmentDetails.Payload.ActiveDeployment.Components {
		if component.Name != nil &&
			*component.Name == componentName {
			replicas = component.Replicas
			break
		}
	}

	return &deploymentName, replicas, nil
}

func init() {
	if logsEnvironmentEnabled {
		logsCmd.AddCommand(logsEnvironmentComponentCmd)
		logsEnvironmentComponentCmd.Flags().StringP("application", "a", "", "Name of the application owning the component")
		logsEnvironmentComponentCmd.Flags().StringP("environment", "e", "", "Environment the component runs in")
		logsEnvironmentComponentCmd.Flags().String("component", "", "The component to follow")
	}
}
