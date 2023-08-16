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

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// promoteApplicationCmd represents the buildApplication command
var promoteApplicationCmd = &cobra.Command{
	Use:   "promote",
	Short: "Will trigger promote of a Radix application",
	Long:  `Triggers promote of a Radix application deployment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		deploymentName, _ := cmd.Flags().GetString("deployment")
		fromEnvironment, _ := cmd.Flags().GetString("from-environment")
		toEnvironment, _ := cmd.Flags().GetString("to-environment")
		triggeredByUser, _ := cmd.Flags().GetString("user")
		follow, _ := cmd.Flags().GetBool("follow")

		if appName == nil || *appName == "" || deploymentName == "" || fromEnvironment == "" || toEnvironment == "" {
			return errors.New("application name, deployment name, from and to environments are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}
		triggerPipelineParams := application.NewTriggerPipelinePromoteParams()
		triggerPipelineParams.SetAppName(*appName)
		triggerPipelineParams.SetPipelineParametersPromote(&models.PipelineParametersPromote{
			DeploymentName:  deploymentName,
			FromEnvironment: fromEnvironment,
			ToEnvironment:   toEnvironment,
			TriggeredBy:     triggeredByUser,
		})

		newJob, err := apiClient.Application.TriggerPipelinePromote(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		if !follow {
			return nil
		}

		jobName := newJob.GetPayload().Name
		return getLogsJob(cmd, apiClient, *appName, jobName)
	},
}

func init() {
	createJobCmd.AddCommand(promoteApplicationCmd)
	promoteApplicationCmd.Flags().StringP("application", "a", "", "Name of the application to be promoted")
	promoteApplicationCmd.Flags().StringP("deployment", "d", "", "Name of a deployment to be promoted")
	promoteApplicationCmd.Flags().StringP("from-environment", "", "", "The deployment source environment")
	promoteApplicationCmd.Flags().StringP("to-environment", "", "", "The deployment target environment")
	promoteApplicationCmd.Flags().StringP("user", "u", "", "The user who triggered the promote pipeline job")
	promoteApplicationCmd.Flags().BoolP("follow", "f", false, "Follow the promote pipeline job log")
	setContextSpecificPersistentFlags(promoteApplicationCmd)
}
