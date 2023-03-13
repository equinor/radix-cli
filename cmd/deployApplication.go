// Copyright © 2022
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
	"github.com/equinor/radix-common/utils/slice"
	"strings"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	commonErrors "github.com/equinor/radix-common/utils/errors"
	"github.com/spf13/cobra"
)

const deployApplicationEnabled = true

var deployApplicationCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Will trigger deploy of a Radix application",
	Long:  `Triggers deploy of a Radix application according to the radix config in its repository's master branch.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var errs []error
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			errs = append(errs, err)
		}
		targetEnvironment, err := cmd.Flags().GetString("environment")
		if err != nil {
			errs = append(errs, err)
		}
		triggeredByUser, err := cmd.Flags().GetString("user")
		if err != nil {
			errs = append(errs, err)
		}
		follow, err := cmd.Flags().GetBool("follow")
		if err != nil {
			errs = append(errs, err)
		}
		componentImageTags, err := cmd.Flags().GetStringSlice("image-tag")
		if len(errs) > 0 {
			return commonErrors.Concat(errs)
		}

		componentTagsMap := slice.Reduce(componentImageTags, make(map[string]string), func(componentTags map[string]string, componentTagPair string) map[string]string {
			if pair := strings.Split(componentTagPair, "="); len(pair) == 2 {
				componentTags[pair[0]] = pair[1] //component-name:tag-name
			}
			return componentTags
		})

		if appName == nil || *appName == "" || targetEnvironment == "" {
			return errors.New("application name and target environment are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineDeployParams()
		triggerPipelineParams.SetAppName(*appName)
		triggerPipelineParams.SetPipelineParametersDeploy(&models.PipelineParametersDeploy{
			ToEnvironment: targetEnvironment,
			ImageTags:     componentTagsMap,
			TriggeredBy:   triggeredByUser,
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
		deployApplicationCmd.Flags().StringP("user", "u", "", "The user who triggered the deploy")
		deployApplicationCmd.Flags().StringSliceP("image-tag", "t", nil, "Image tag for a component: component-name=tag-name. Multiple pairs can be specified.")
		deployApplicationCmd.Flags().BoolP("follow", "f", false, "Follow deploy")
	}
}
