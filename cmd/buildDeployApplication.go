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

const buildDeployApplicationEnabled = false

// buildDeployApplicationCmd represents the buildApplication command
var buildDeployApplicationCmd = &cobra.Command{
	Use:   "build-deploy",
	Short: "Will trigger build-deploy of a Radix application",
	Long:  `Triggers build-deploy of Radix application, if branch to environment map exists for the branch in the Radix config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		branch, _ := cmd.Flags().GetString("branch")
		commitID, _ := cmd.Flags().GetString("commitID")
		follow, _ := cmd.Flags().GetBool("follow")

		if appName == nil || *appName == "" || branch == "" {
			return errors.New("Application name and branch are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineBuildDeployParams()
		triggerPipelineParams.SetAppName(*appName)
		triggerPipelineParams.SetPipelineParametersBuild(&models.PipelineParametersBuild{
			Branch:   branch,
			CommitID: commitID,
		})

		newJob, err := apiClient.Application.TriggerPipelineBuildDeploy(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		jobName := newJob.GetPayload().Name
		if follow {
			followJob(cmd, apiClient, *appName, jobName)
		}

		return nil
	},
}

func init() {
	if buildDeployApplicationEnabled {
		triggerCmd.AddCommand(buildDeployApplicationCmd)
		buildDeployApplicationCmd.Flags().StringP("application", "a", "", "Name of the application to build-deploy")
		buildDeployApplicationCmd.Flags().StringP("branch", "b", "master", "Branch to build-deploy from")
		buildDeployApplicationCmd.Flags().StringP("commitID", "i", "", "Commit id")
		buildDeployApplicationCmd.Flags().BoolP("follow", "f", false, "Follow build-deploy")
	}
}
