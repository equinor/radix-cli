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

	"github.com/equinor/radix-cli/pkg/model"
	log "github.com/sirupsen/logrus"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/spf13/cobra"
)

var overrideUseBuildCache model.BoolPtr

// createBuildDeployApplicationCmd represents the buildApplication command
var createBuildDeployApplicationCmd = &cobra.Command{
	Use:   "build-deploy",
	Short: "Will trigger build-deploy of a Radix application",
	Long:  `Triggers build-deploy of Radix application, if branch to environment map exists for the branch in the Radix config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		branch, _ := cmd.Flags().GetString(flagnames.Branch)
		commitID, _ := cmd.Flags().GetString(flagnames.CommitID)
		follow, _ := cmd.Flags().GetBool(flagnames.Follow)

		if appName == nil || *appName == "" || branch == "" {
			return errors.New("application name and branch are required")
		}
		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineBuildDeployParams()
		triggerPipelineParams.SetAppName(*appName)
		triggerPipelineParams.SetPipelineParametersBuild(&models.PipelineParametersBuild{
			Branch:                branch,
			CommitID:              commitID,
			OverrideUseBuildCache: overrideUseBuildCache.Get(),
		})

		newJob, err := apiClient.Application.TriggerPipelineBuildDeploy(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		jobName := newJob.GetPayload().Name
		log.Infof("Build-deploy pipeline job triggered with the name %s\n", jobName)
		if !follow {
			return nil
		}
		return getLogsJob(cmd, apiClient, *appName, jobName)
	},
}

func init() {
	createJobCmd.AddCommand(createBuildDeployApplicationCmd)
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to build-deploy")
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.Branch, "b", "master", "Branch to build-deploy from")
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.CommitID, "i", "", "Commit id")
	createBuildDeployApplicationCmd.Flags().BoolP(flagnames.Follow, "f", false, "Follow build-deploy")
	createBuildDeployApplicationCmd.Flags().Var(&overrideUseBuildCache, flagnames.UseBuildCache, "Optional. Overrides configured or default useBuildCache option. It is applicable when the useBuildKit option is set as true.")

	setContextSpecificPersistentFlags(createBuildDeployApplicationCmd)
}
