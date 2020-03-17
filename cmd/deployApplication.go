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

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const deployApplicationEnabled = false

var deployApplicationCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Will trigger deploy of a Radix application",
	Long:  `Triggers deploy of a Radix application according to the radix config in its repository's master branch.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}
		targetEnvironment, _ := cmd.Flags().GetString("environment")
		follow, _ := cmd.Flags().GetBool("follow")

		if appName == nil || *appName == "" || targetEnvironment == "" {
			return errors.New("Application name and target environment are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineDeployParams()
		triggerPipelineParams.SetAppName(*appName)
		triggerPipelineParams.SetPipelineParametersDeploy(&models.PipelineParametersDeploy{
			ToEnvironment: targetEnvironment,
		})

		newJob, err := apiClient.Application.TriggerPipelineDeploy(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		if follow {
			jobName := newJob.GetPayload().Name
			getLogsJob(cmd, apiClient, *appName, jobName)
		}

		return nil
	},
}

func init() {
	if deployApplicationEnabled {
		createJobCmd.AddCommand(deployApplicationCmd)
		deployApplicationCmd.Flags().StringP("application", "a", "", "Name of the application to deploy")
		deployApplicationCmd.Flags().StringP("environment", "e", "", "Target environment to deploy in ('prod', 'dev', 'playground')")
		deployApplicationCmd.Flags().BoolP("follow", "f", false, "Follow deploy")
	}
}
