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
	"github.com/spf13/cobra"
)

// createPromotePipelineJobCmd represents the buildApplication command
var createPromotePipelineJobCmd = &cobra.Command{
	Use:   "promote",
	Short: "Will trigger promote of a Radix application",
	Long:  `Triggers promote of a Radix application deployment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		useActiveDeployment, _ := cmd.Flags().GetBool("use-active-deployment")
		deploymentName, _ := cmd.Flags().GetString("deployment")
		fromEnvironment, _ := cmd.Flags().GetString("from-environment")
		toEnvironment, _ := cmd.Flags().GetString("to-environment")
		triggeredByUser, _ := cmd.Flags().GetString("user")
		follow, _ := cmd.Flags().GetBool("follow")

		if !useActiveDeployment && deploymentName == "" {
			return errors.New("Specifying deployment name or setting use-active-deployment is required")
		}
		if useActiveDeployment && deploymentName != "" {
			return errors.New("You must specify either deployment name or set use-active-deployment")
		}

		if appName == nil || *appName == "" || fromEnvironment == "" || toEnvironment == "" {
			return errors.New("application name, deployment name, from and to environments are required")
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
	createPromotePipelineJobCmd.Flags().StringP("application", "a", "", "Name of the application to be promoted")
	createPromotePipelineJobCmd.Flags().StringP("deployment", "d", "", "Name of a deployment to be promoted")
	createPromotePipelineJobCmd.Flags().StringP("from-environment", "", "", "The deployment source environment")
	createPromotePipelineJobCmd.Flags().StringP("to-environment", "", "", "The deployment target environment")
	createPromotePipelineJobCmd.Flags().StringP("user", "u", "", "The user who triggered the promote pipeline job")
	createPromotePipelineJobCmd.Flags().BoolP("follow", "f", false, "Follow the promote pipeline job log")
	createPromotePipelineJobCmd.Flags().BoolP("use-active-deployment", "", false, "Promote the active deployment")
	setContextSpecificPersistentFlags(createPromotePipelineJobCmd)
}
