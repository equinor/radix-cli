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
	"time"

	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/replicalog"
	"github.com/spf13/cobra"
)

// logsJobCmd represents the logsJobCmd command
var logsJobCmd = &cobra.Command{
	Use:     "pipeline-job",
	Aliases: []string{"job"},
	Short:   "Gets logs for a pipeline job",
	Long: `Gets and follows logs for a pipeline job.

It may take few seconds to get the log.`,
	Example: `# Get logs for a pipeline job 
rx get logs pipeline-job --application radix-test --job radix-pipeline-20230323185013-ehvnz`,

	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == "" {
			return errors.New("application name is required")
		}

		jobName, _ := cmd.Flags().GetString(flagnames.Job)

		if jobName == "" {
			return errors.New("`job` is required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		return replicalog.New(
			cmd.ErrOrStderr(),
			replicalog.GetReplicasForJob(apiClient, appName, jobName),
			replicalog.GetLogsForJob(apiClient, appName, jobName),
			time.Second, // not used
		).StreamLogs(cmd.Context(), true)
	},
}

func init() {
	logsCmd.AddCommand(logsJobCmd)

	logsJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application for the job")
	logsJobCmd.Flags().StringP(flagnames.Job, "j", "", "The job to get logs for")

	_ = logsJobCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = logsJobCmd.RegisterFlagCompletionFunc(flagnames.Job, completion.JobCompletion)
	setContextSpecificPersistentFlags(logsJobCmd)
}
