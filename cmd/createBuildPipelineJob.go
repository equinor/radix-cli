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

	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/model"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var overrideUseBuildCacheForBuild model.BoolPtr

var createBuildPipelineJobCmd = &cobra.Command{
	Use:   "build",
	Short: "Will trigger build of a Radix application",
	Long:  `Triggers build of Radix application, for branches that are mapped to a environment in the Radix config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var errs []error
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			errs = append(errs, err)
		}
		if appName == "" {
			errs = append(errs, errors.New("application name is required"))
		}
		branch, err := cmd.Flags().GetString(flagnames.Branch)
		if err != nil {
			errs = append(errs, err)
		}
		targetEnvironment, err := cmd.Flags().GetString(flagnames.Environment)
		if err != nil {
			errs = append(errs, err)
		}
		follow, err := cmd.Flags().GetBool(flagnames.Follow)
		if err != nil {
			errs = append(errs, err)
		}
		if len(errs) > 0 {
			return errors.Join(errs...)
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}
		triggerDeployParams := application.NewTriggerPipelineBuildParams()
		triggerDeployParams.SetAppName(appName)
		triggerDeployParams.SetPipelineParametersBuild(&models.PipelineParametersBuild{
			Branch:                branch,
			ToEnvironment:         targetEnvironment,
			OverrideUseBuildCache: overrideUseBuildCacheForBuild.Get(),
		})
		newJob, err := apiClient.Application.TriggerPipelineBuild(triggerDeployParams, nil)
		if err != nil {
			return err
		}

		jobName := newJob.GetPayload().Name
		log.Infof("Build pipeline job triggered with the name %s\n", *jobName)
		if !follow {
			return nil
		}

		return getLogsJob(cmd, apiClient, appName, *jobName)
	},
}

func init() {
	createJobCmd.AddCommand(createBuildPipelineJobCmd)
	createBuildPipelineJobCmd.Flags().StringP(flagnames.Branch, "b", "master", "Branch to build from")
	createBuildPipelineJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to build")
	createBuildPipelineJobCmd.Flags().StringP(flagnames.Environment, "e", "", "Optional. Target environment to deploy in ('prod', 'dev', 'playground'), when multiple environments are built from the selected branch.")
	createBuildPipelineJobCmd.Flags().BoolP(flagnames.Follow, "f", false, "Follow build")
	createBuildPipelineJobCmd.Flags().Var(&overrideUseBuildCacheForBuild, flagnames.UseBuildCache, "Optional. Overrides configured or default useBuildCache option. It is applicable when the useBuildKit option is set as true.")
	if err := createBuildPipelineJobCmd.MarkFlagRequired(flagnames.Branch); err != nil {
		log.Fatalf("Error during command initialization: %v", err)
	}
	_ = createPromotePipelineJobCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(createBuildPipelineJobCmd)
}
