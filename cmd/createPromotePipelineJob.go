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
	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/spf13/cobra"
)

// createPromotePipelineJobCmd represents the buildApplication command
var createPromotePipelineJobCmd = &cobra.Command{
	Use:   "promote",
	Short: "Will trigger promote of a Radix application",
	Long:  `Triggers promote of a Radix application deployment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		useActiveDeployment, _ := cmd.Flags().GetBool(flagnames.UseActiveDeployment)
		deploymentName, _ := cmd.Flags().GetString(flagnames.Deployment)
		fromEnvironment, _ := cmd.Flags().GetString(flagnames.FromEnvironment)
		toEnvironment, _ := cmd.Flags().GetString(flagnames.ToEnvironment)
		triggeredByUser, _ := cmd.Flags().GetString(flagnames.User)
		follow, _ := cmd.Flags().GetBool(flagnames.Follow)

		if !useActiveDeployment && deploymentName == "" {
			return errors.New("Specifying deployment name or setting use-active-deployment is required")
		}
		if useActiveDeployment && deploymentName != "" {
			return errors.New("You cannot set use-active-deployment and specify deployment name at the same time")
		}

		if appName == nil || *appName == "" || fromEnvironment == "" || toEnvironment == "" {
			return errors.New("application name, from and to environments are required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		if useActiveDeployment {
			d, err := getActiveDeploymentName(apiClient, *appName, fromEnvironment)
			if err != nil {
				return err
			}

			deploymentName = d
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

		jobName := newJob.GetPayload().Name
		log.Infof("Promote pipeline job triggered with the name %s\n", jobName)
		if !follow {
			return nil
		}

		return getLogsJob(cmd, apiClient, *appName, jobName)
	},
}

func getActiveDeploymentName(apiClient *apiclient.Radixapi, appName, envName string) (string, error) {
	params := environment.NewGetEnvironmentParams()
	params.SetAppName(appName)
	params.SetEnvName(envName)

	resp, err := apiClient.Environment.GetEnvironment(params, nil)
	if err != nil {
		return "", errors.Wrap(err, "Failed to get environment details")
	}

	if resp.Payload.ActiveDeployment == nil || resp.Payload.ActiveDeployment.Name == "" {
		return "", errors.Errorf("Environment '%s' does not have any active deployments", envName)
	}

	return resp.Payload.ActiveDeployment.Name, nil
}

func init() {
	createJobCmd.AddCommand(createPromotePipelineJobCmd)
	createPromotePipelineJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to be promoted")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.Deployment, "d", "", "Name of a deployment to be promoted")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.FromEnvironment, "", "", "The deployment source environment")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.ToEnvironment, "", "", "The deployment target environment")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.User, "u", "", "The user who triggered the promote pipeline job")
	createPromotePipelineJobCmd.Flags().BoolP(flagnames.Follow, "f", false, "Follow the promote pipeline job log")
	createPromotePipelineJobCmd.Flags().BoolP(flagnames.UseActiveDeployment, "", false, "Promote the active deployment")
	setContextSpecificPersistentFlags(createPromotePipelineJobCmd)
}
