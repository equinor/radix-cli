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

	"github.com/equinor/radix-cli/generated-client/client/pipeline_job"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// restartPipelineJobCmd represents the rerun command
var restartPipelineJobCmd = &cobra.Command{
	Use:     "pipeline-job",
	Aliases: []string{"job"},
	Short:   "Restart Radix pipeline job",
	Long:    "Restart failed of stopped Radix pipeline job.",
	Example: `rx restart pipeline-job --application radix-test --job radix-pipeline-20230323185013-ehvnz`,

	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}

		jobName, _ := cmd.Flags().GetString(flagnames.Job)

		if jobName == "" {
			return errors.New("`job` is required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		params := pipeline_job.NewRerunApplicationJobParams()
		params.AppName = *appName
		params.JobName = jobName

		_, err = apiClient.PipelineJob.RerunApplicationJob(params, nil)
		return err
	},
}

func init() {
	restartCmd.AddCommand(restartPipelineJobCmd)
	restartPipelineJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application for the job")
	restartPipelineJobCmd.Flags().StringP(flagnames.Job, "j", "", "The job to restart")

	_ = getApplicationCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(restartPipelineJobCmd)
}
