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

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/model"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deployExternalDNSAlias model.BoolPtr

var createApplyConfigPipelineJobCmd = &cobra.Command{
	Use:   "apply-config",
	Short: "Will trigger apply-config of a Radix application",
	Long:  "Triggers applyConfig of a Radix application according to the radix config in its repository's master branch.",
	Example: `  # Create a Radix pipeline apply-config job to apply the radixconfig properties without re-building or re-deploying components.
By default it is applied changes in properties DNS alias, build secrets, create new or soft-delete existing environments.
  rx create job apply-config --application radix-test
  
  # Create a Radix pipeline applyConfig-only job, short option versions 
  rx create job apply-config -a radix-test

  # Create a Radix pipeline apply-config job with external DNS aliases
  rx create job apply-config -a radix-test --deploy-external-dns-alias true`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var errs []error
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			errs = append(errs, err)
		}
		triggeredByUser, err := cmd.Flags().GetString(flagnames.User)
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
		if appName == "" {
			return errors.New("application name is required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineApplyConfigParams()
		triggerPipelineParams.SetAppName(appName)
		parametersApplyConfig := models.PipelineParametersApplyConfig{
			TriggeredBy:       triggeredByUser,
			DeployExternalDNS: deployExternalDNSAlias.Get(),
		}
		triggerPipelineParams.SetPipelineParametersApplyConfig(&parametersApplyConfig)

		newJob, err := apiClient.Application.TriggerPipelineApplyConfig(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		jobName := newJob.GetPayload().Name
		log.Infof("Apply-config pipeline job triggered with the name %s\n", jobName)
		if !follow {
			return nil
		}
		return getLogsJob(cmd, apiClient, appName, jobName)
	},
}

func init() {
	createJobCmd.AddCommand(createApplyConfigPipelineJobCmd)
	createApplyConfigPipelineJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to apply-config")
	createApplyConfigPipelineJobCmd.Flags().Var(&deployExternalDNSAlias, flagnames.DeployExternalDNSAlias, "Optional. Deploy changes in External DNS aliases. Default is false.")
	createApplyConfigPipelineJobCmd.Flags().StringP(flagnames.User, "u", "", "The user who triggered the apply-config")
	createApplyConfigPipelineJobCmd.Flags().BoolP(flagnames.Follow, "f", false, "Follow applyConfig")
	_ = createApplyConfigPipelineJobCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(createApplyConfigPipelineJobCmd)
}
