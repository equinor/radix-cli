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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rerunCmd represents the rerun command
var rerunCmd = &cobra.Command{
	Use:   "rerun",
	Short: "Rerun Radix pipeline job",
	Long: `Rerun failed of stopped Radix pipeline job .

	Example:
	# Get logs for a pipeline job
	rx rerun --application radix-test --job radix-pipeline-20230323185013-ehvnz
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("application name is required")
		}

		jobName, _ := cmd.Flags().GetString("job")

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

		job, err := apiClient.PipelineJob.RerunApplicationJob(&params, nil)
		log.Info(job) // TODO
		return err
	},
}

func init() {
	rootCmd.AddCommand(rerunCmd)
	rerunCmd.Flags().StringP("application", "a", "", "Name of the application for the job")
	rerunCmd.Flags().StringP("job", "j", "", "The job to get logs for")
	rerunCmd.Flags().StringP("user", "u", "", "The user who triggered the deploy")
	rerunCmd.Flags().BoolP("follow", "f", false, "Follow deploy")
	setContextSpecificPersistentFlags(rerunCmd)
}
