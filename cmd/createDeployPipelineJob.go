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
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

var createDeployPipelineJobCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Will trigger deploy of a Radix application",
	Long: `Triggers deploy of a Radix application according to the radix config in its repository's master branch.

Examples:
  # Create a Radix pipeline deploy-only job to deploy an application "radix-test" to an environment "dev" 
  rx create job deploy --application radix-test --environment dev

  # Create a Radix pipeline deploy-only job, short option versions 
  rx create job deploy -a radix-test -e dev

  # Create a Radix pipeline deploy-only job to deploy an application with re-defined image-tags. These image tags will re-define  
  rx create job deploy --application radix-test --environment dev --image-tag web-app=web-app-v2.1

  # Create a Radix pipeline deploy-only job with re-defined image-tags for components, short option versions 
  rx create job deploy -a radix-test -e dev -t web-app=web-app-v2.1 -t api-server=api-v1.0

  # Create a Radix pipeline deploy-only job to deploy only specific components 
  rx create job deploy -a radix-test -e dev --component web-app --component api-server
`,
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
		imageTagNames, err := cmd.Flags().GetStringToString("image-tag-name")
		if err != nil {
			errs = append(errs, err)
		}
		if len(errs) > 0 {
			return errors.Join(errs...)
		}
		if appName == nil || *appName == "" || targetEnvironment == "" {
			return errors.New("application name and target environment are required")
		}
		commitID, _ := cmd.Flags().GetString("commitID")
		err = validateCommitID(commitID)
		if err != nil {
			return err
		}
		componentNames, err := cmd.Flags().GetStringSlice("component")
		if err != nil {
			errs = append(errs, err)
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineDeployParams()
		triggerPipelineParams.SetAppName(*appName)
		parametersDeploy := models.PipelineParametersDeploy{
			ToEnvironment: targetEnvironment,
			ImageTagNames: imageTagNames,
			TriggeredBy:   triggeredByUser,
			CommitID:      commitID,
		}
		if components := strings.Join(componentNames, ","); len(components) > 0 {
			parametersDeploy.Components = components
		}
		triggerPipelineParams.SetPipelineParametersDeploy(&parametersDeploy)

		newJob, err := apiClient.Application.TriggerPipelineDeploy(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		jobName := newJob.GetPayload().Name
		log.Infof("Deploy pipeline job triggered with the name %s\n", jobName)
		if !follow {
			return nil
		}
		return getLogsJob(cmd, apiClient, *appName, jobName)
	},
}

func validateCommitID(commitID string) error {
	if len(commitID) == 0 {
		return nil
	}
	re, err := regexp.Compile("^[a-f0-9]{40}$")
	if err != nil {
		return fmt.Errorf("error compiling the regex for the GitHub CommitID: %w", err)
	}
	if !re.MatchString(commitID) {
		return fmt.Errorf("invalid GitHub CommitID format")
	}
	return nil
}

func init() {
	createJobCmd.AddCommand(createDeployPipelineJobCmd)
	createDeployPipelineJobCmd.Flags().StringP("application", "a", "", "Name of the application to deploy")
	createDeployPipelineJobCmd.Flags().StringP("environment", "e", "", "Target environment to deploy in ('prod', 'dev', 'playground')")
	createDeployPipelineJobCmd.Flags().StringP("user", "u", "", "The user who triggered the deploy")
	createDeployPipelineJobCmd.Flags().StringToStringP("image-tag-name", "t", map[string]string{}, "Image tag name for a component: component-name=tag-name. Multiple pairs can be specified.")
	createDeployPipelineJobCmd.Flags().StringP("commitID", "i", "", "An optional 40 character commit id to tag the new pipeline job")
	createDeployPipelineJobCmd.Flags().StringSlice("component", []string{}, "Optional component to deploy, if only specific component need to be deployed")
	createDeployPipelineJobCmd.Flags().BoolP("follow", "f", false, "Follow deploy")
	setContextSpecificPersistentFlags(createDeployPipelineJobCmd)
}
