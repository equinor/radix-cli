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
	"sort"
	"time"

	"github.com/equinor/radix-common/utils/slice"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	log1 "github.com/equinor/radix-cli/pkg/utils/replicalog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createPromotePipelineJobCmd represents the buildApplication command
var createPromotePipelineJobCmd = &cobra.Command{
	Use:   "promote",
	Short: "Will trigger promote of a Radix application",
	Long:  `Triggers promote of a Radix application deployment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		useActiveDeployment, _ := cmd.Flags().GetBool(flagnames.UseActiveDeployment)
		deploymentName, _ := cmd.Flags().GetString(flagnames.Deployment)
		commitId, _ := cmd.Flags().GetString(flagnames.CommitID)
		fromEnvironment, _ := cmd.Flags().GetString(flagnames.FromEnvironment)
		toEnvironment, _ := cmd.Flags().GetString(flagnames.ToEnvironment)
		triggeredByUser, _ := cmd.Flags().GetString(flagnames.User)
		follow, _ := cmd.Flags().GetBool(flagnames.Follow)

		if !useActiveDeployment && deploymentName == "" && commitId == "" {
			return errors.New("specifying deployment name or setting use-active-deployment is required")
		}
		if useActiveDeployment && deploymentName != "" {
			return errors.New("you cannot set use-active-deployment and specify deployment name at the same time")
		}

		if appName == "" || fromEnvironment == "" || toEnvironment == "" {
			return errors.New("application name, from and to environments are required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		if useActiveDeployment {
			name, err := getActiveDeploymentName(apiClient, appName, fromEnvironment)
			if err != nil {
				return err
			}

			deploymentName = name
		}

		if commitId != "" {
			if deploymentName != "" {
				return errors.New("deployment name or use-active-deployment and commitID cannot be used at the same time")
			}
			name, err := getLastDeploymentNameByCommitId(apiClient, appName, fromEnvironment, commitId)
			if err != nil {
				return err
			}
			deploymentName = name
		}

		triggerPipelineParams := application.NewTriggerPipelinePromoteParams()
		triggerPipelineParams.SetAppName(appName)
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
		printPayload(newJob.GetPayload())

		jobName := newJob.GetPayload().Name
		log.Infof("Promote pipeline job triggered with the name %s\n", *jobName)
		if !follow {
			return nil
		}

		return log1.New(
			cmd.ErrOrStderr(),
			log1.GetReplicasForJob(apiClient, appName, *jobName),
			log1.GetLogsForJob(apiClient, appName, *jobName),
			time.Second, // not used
		).StreamLogs(cmd.Context(), true)
	},
}

func getActiveDeploymentName(apiClient *radixapi.Radixapi, appName, envName string) (string, error) {
	params := environment.NewGetEnvironmentParams()
	params.SetAppName(appName)
	params.SetEnvName(envName)

	resp, err := apiClient.Environment.GetEnvironment(params, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get environment details: %w", err)
	}

	if resp.Payload.ActiveDeployment == nil || resp.Payload.ActiveDeployment.Name == nil || *resp.Payload.ActiveDeployment.Name == "" {
		return "", fmt.Errorf("environment '%s' does not have any active deployments", envName)
	}

	return *resp.Payload.ActiveDeployment.Name, nil
}

func getLastDeploymentNameByCommitId(apiClient *radixapi.Radixapi, appName, envName, commitId string) (string, error) {
	params := environment.NewGetEnvironmentParams()
	params.SetAppName(appName)
	params.SetEnvName(envName)

	resp, err := apiClient.Environment.GetEnvironment(params, nil)
	if err != nil || resp.Payload == nil {
		return "", fmt.Errorf("failed to get environment details: %w", err)
	}
	deploymentSummaries := slice.FindAll(resp.Payload.Deployments, func(item *models.DeploymentSummary) bool { return item.GitCommitHash == commitId })
	if len(deploymentSummaries) == 0 {
		return "", fmt.Errorf("no deployments found with commitID '%s'", commitId)
	}
	if len(deploymentSummaries) == 1 {
		return *deploymentSummaries[0].Name, nil
	}

	sort.Slice(deploymentSummaries, func(i, j int) bool {
		if deploymentSummaries[i].ActiveFrom == nil && deploymentSummaries[j].ActiveFrom == nil {
			return false
		}
		if deploymentSummaries[i].ActiveFrom == nil {
			return true
		}
		if deploymentSummaries[j].ActiveFrom == nil {
			return false
		}
		return time.Time(*deploymentSummaries[i].ActiveFrom).Before(time.Time(*deploymentSummaries[j].ActiveFrom))
	})
	return *deploymentSummaries[len(deploymentSummaries)-1].Name, nil
}

func init() {
	createJobCmd.AddCommand(createPromotePipelineJobCmd)
	createPromotePipelineJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to be promoted")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.Deployment, "d", "", "(Optional) Name of a deployment to be promoted. This cannot be used together with the option commitID")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.CommitID, "i", "", "(Optional) An optional 40 character commitID of the deployment to promote. This cannot be used together with an option deployment. The latest deployment is promoted if there are multiple deployments with the same commitID")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.FromEnvironment, "", "", "The deployment source environment")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.ToEnvironment, "", "", "The deployment target environment")
	createPromotePipelineJobCmd.Flags().StringP(flagnames.User, "u", "", "The user who triggered the promote pipeline job")
	createPromotePipelineJobCmd.Flags().BoolP(flagnames.Follow, "f", false, "Follow the promote pipeline job log")
	createPromotePipelineJobCmd.Flags().BoolP(flagnames.UseActiveDeployment, "", false, "(Optional) Promote the active deployment")
	_ = createPromotePipelineJobCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = createPromotePipelineJobCmd.RegisterFlagCompletionFunc(flagnames.FromEnvironment, completion.EnvironmentCompletion)
	_ = createPromotePipelineJobCmd.RegisterFlagCompletionFunc(flagnames.ToEnvironment, completion.EnvironmentCompletion)
	_ = createPromotePipelineJobCmd.RegisterFlagCompletionFunc(flagnames.Deployment, completion.CreateDeploymentCompletion(flagnames.FromEnvironment, true))
	_ = createPromotePipelineJobCmd.RegisterFlagCompletionFunc(flagnames.CommitID, completion.CreateDeploymentCommitIDCompletion(flagnames.FromEnvironment, true))
	setContextSpecificPersistentFlags(createPromotePipelineJobCmd)
}
