// Copyright © 2023
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/equinor/radix-cli/pkg/flagnames"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createBuildPipelineJobCmd represents the createBuildPipelineJob command
var createBuildPipelineJobCmd = &cobra.Command{
	Use:   "build",
	Short: "Will trigger build of a Radix application",
	Long:  `Triggers build of Radix application, for branches that are mapped to a environment in the Radix config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}
		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}
		branch, _ := cmd.Flags().GetString(flagnames.Branch)
		follow, _ := cmd.Flags().GetBool(flagnames.Follow)

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}
		triggerDeployParams := application.NewTriggerPipelineBuildParams()
		triggerDeployParams.SetAppName(*appName)
		triggerDeployParams.SetPipelineParametersBuild(&models.PipelineParametersBuild{
			Branch: branch,
		})
		newJob, err := apiClient.Application.TriggerPipelineBuild(triggerDeployParams, nil)
		if err != nil {
			return err
		}

		jobName := newJob.GetPayload().Name
		log.SetOutput(cmd.OutOrStdout())
		log.Infof("Build pipeline job triggered with the name %s\n", jobName)
		if !follow {
			return nil
		}

		return getLogsJob(cmd, apiClient, *appName, jobName)
	},
}

func init() {
	createJobCmd.AddCommand(createBuildPipelineJobCmd)
	createBuildPipelineJobCmd.Flags().StringP(flagnames.Branch, "b", "master", "Branch to build from")
	createBuildPipelineJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to build")
	createBuildPipelineJobCmd.Flags().BoolP(flagnames.Follow, "f", false, "Follow build")
	if err := createBuildPipelineJobCmd.MarkFlagRequired(flagnames.Branch); err != nil {
		log.Fatalf("Error during command initialization: %v", err)
	}
	setContextSpecificPersistentFlags(createBuildPipelineJobCmd)
}
