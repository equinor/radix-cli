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
	"github.com/spf13/cobra"
)

// followEnvironmentComponentCmd represents the followEnvironmentComponentCmd command
var followEnvironmentComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Will follow a component in an environment",
	Long:  `Will follow a component in an environment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("Application name is required")
		}

		environmentName, _ := cmd.Flags().GetString("environment")
		componentName, _ := cmd.Flags().GetString("component")

		if environmentName == "" || componentName == "" {
			return errors.New("Both `environment` and `component` are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		deploymentName, replicas, err := getReplicasForComponent(apiClient, *appName, environmentName, componentName)
		if err != nil {
			return err
		}

		refreshLog := time.Tick(settings.DeltaRefreshApplication)
		loggedForReplica := make(map[string]int)

		for {
			select {
			case <-refreshLog:

				for i, replica := range replicas {
					logParameters := component.NewLogParams()
					logParameters.WithAppName(*appName)
					logParameters.WithDeploymentName(*deploymentName)
					logParameters.WithComponentName(componentName)
					logParameters.WithPodName(replica)

					logData, err := apiClient.Component.Log(logParameters, nil)
					if err != nil {
						// Replicas may have died
						deploymentName, replicas, err = getReplicasForComponent(apiClient, *appName, environmentName, componentName)
						if err != nil {
							return err
						}

					} else {
						totalLinesLogged := 0

						if _, contained := loggedForReplica[replica]; contained {
							totalLinesLogged = loggedForReplica[replica]
						}

						logLines := strings.Split(strings.Replace(logData.Payload, "\r\n", "\n", -1), "\n")
						logged := log.From(cmd, replica, totalLinesLogged, logLines, log.GetColor(i))

						totalLinesLogged += logged
						loggedForReplica[replica] = totalLinesLogged
					}
				}
			}

		}
	},
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
		return nil, nil, errors.New("Active deployment was not found in environment")
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
	followEnvironmentComponentCmd.Flags().StringP("application", "a", "", "Name of the application owning the component")
	followEnvironmentComponentCmd.Flags().StringP("environment", "e", "", "Environment the component runs in")
	followEnvironmentComponentCmd.Flags().String("component", "", "The component to follow")
}
