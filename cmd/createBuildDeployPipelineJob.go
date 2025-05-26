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

var overrideUseBuildCacheForBuildDeploy, refreshBuildCacheForBuildDeploy model.BoolPtr

// createBuildDeployApplicationCmd represents the buildApplication command
var createBuildDeployApplicationCmd = &cobra.Command{
	Use:   "build-deploy",
	Short: "Will trigger build-deploy of a Radix application",
	Long:  `Triggers build-deploy of Radix application, if branch to environment map exists for the branch in the Radix config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var errs []error
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			errs = append(errs, err)
		}

		gitRef, gitRefType, err := getGitRefAndType(cmd)
		if err != nil {
			errs = append(errs, err)
		}
		commitID, err := cmd.Flags().GetString(flagnames.CommitID)
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

		if appName == "" || gitRef == "" {
			errs = append(errs, errors.New("application name and branch or tag are required"))
		}
		if len(errs) > 0 {
			return errors.Join(errs...)
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineBuildDeployParams()
		triggerPipelineParams.SetAppName(appName)
		triggerPipelineParams.SetPipelineParametersBuild(&models.PipelineParametersBuild{
			GitRef:                gitRef,
			GitRefType:            gitRefType,
			ToEnvironment:         targetEnvironment,
			CommitID:              commitID,
			OverrideUseBuildCache: overrideUseBuildCacheForBuildDeploy.Get(),
			RefreshBuildCache:     refreshBuildCacheForBuildDeploy.Get(),
		})

		newJob, err := apiClient.Application.TriggerPipelineBuildDeploy(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		jobName := newJob.GetPayload().Name
		log.Infof("Build-deploy pipeline job triggered with the name %s\n", *jobName)
		if !follow {
			return nil
		}
		return getLogsJob(cmd, apiClient, appName, *jobName)
	},
}

func init() {
	createJobCmd.AddCommand(createBuildDeployApplicationCmd)
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to build-deploy")
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.Branch, "b", "", "GitHub branch to build-deploy from")
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.Tag, "", "", "GitHub tag to build-deploy from")
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.Environment, "e", "", "Optional. Target environment to deploy in ('prod', 'dev', 'playground'), when multiple environments are built from the selected branch.")
	createBuildDeployApplicationCmd.Flags().StringP(flagnames.CommitID, "i", "", "Commit id")
	createBuildDeployApplicationCmd.Flags().BoolP(flagnames.Follow, "f", false, "Follow build-deploy")
	createBuildDeployApplicationCmd.Flags().Var(&overrideUseBuildCacheForBuildDeploy, flagnames.UseBuildCache, "Optional. Overrides configured or default useBuildCache option. It is applicable when the useBuildKit option is set as true.")
	createBuildDeployApplicationCmd.Flags().Var(&refreshBuildCacheForBuildDeploy, flagnames.RefreshBuildCache, "Optional. Refreshes the build cache. It is applicable when the useBuildKit option is set as true.")

	createBuildDeployApplicationCmd.MarkFlagsMutuallyExclusive(flagnames.Branch, flagnames.Tag)
	_ = createBuildDeployApplicationCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(createBuildDeployApplicationCmd)
}
