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
	"strings"
	"time"

	"k8s.io/utils/strings/slices"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/pipeline_job"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/log"
	"github.com/spf13/cobra"
)

const (
	jobStatusRunning   = "Running"
	jobStatusFailed    = "Failed"
	jobStatusSucceeded = "Succeeded"
	jobStatusStopped   = "Stopped"

	stepStatusWaiting = "Waiting"
)

var completedJobStatuses = []string{jobStatusSucceeded, jobStatusStopped, jobStatusFailed}

// logsJobCmd represents the logsJobCmd command
var logsJobCmd = &cobra.Command{
	Use:     "pipeline-job",
	Aliases: []string{"job"},
	Short:   "Gets logs for a pipeline job",
	Long: `Gets and follows logs for a pipeline job.

It may take few seconds to get the log.

Example:
  # Get logs for a pipeline job 
  rx get logs pipeline-job --application radix-test --job radix-pipeline-20230323185013-ehvnz
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

		return getLogsJob(cmd, apiClient, *appName, jobName)
	},
}

func getLogsJob(cmd *cobra.Command, apiClient *apiclient.Radixapi, appName, jobName string) error {
	timeout := time.NewTimer(settings.DeltaTimeout)
	refreshLog := time.Tick(settings.DeltaRefreshApplication)

	// Sometimes, even though we get delta, the log is the same as previous
	previousLogForStep := make(map[string][]string)
	jobParameters := pipeline_job.NewGetApplicationJobParams()
	jobParameters.SetAppName(appName)
	jobParameters.SetJobName(jobName)
	getLogAttempts := 5
	getLogStartTime := time.Now()

	for {
		select {
		case <-refreshLog:
			respJob, _ := apiClient.PipelineJob.GetApplicationJob(jobParameters, nil)
			if respJob == nil {
				continue
			}
			if isCompletedJob(respJob.Payload.Status) {
				return errorAndLogCompletedJob(respJob.Payload.Status, cmd)
			}
			loggedForJob := false

			for i, step := range respJob.Payload.Steps {
				if step.Status == stepStatusWaiting {
					continue
				}

				// Sometimes, even though we get delta, the log is the same as previous
				previousLogLines := previousLogForStep[step.Name]
				stepLogsParams := pipeline_job.NewGetPipelineJobStepLogsParams()
				stepLogsParams.SetAppName(jobParameters.AppName)
				stepLogsParams.SetJobName(jobParameters.JobName)
				stepLogsParams.SetStepName(step.Name)

				jobStepLog, err := apiClient.PipelineJob.GetPipelineJobStepLogs(stepLogsParams, nil)
				if err != nil {
					log.Print(cmd, "radix-cli", fmt.Sprintf("Failed to get pipeline job logs. %v", err), log.Red)
					break
				}
				logLines := strings.Split(strings.Replace(jobStepLog.Payload, "\r\n", "\n", -1), "\n")
				if len(logLines) > 0 && !strings.EqualFold(logLines[0], "") {
					log.PrintLines(cmd, step.Name, previousLogLines, logLines, log.GetColor(i))
					loggedForJob = true
					previousLogForStep[step.Name] = logLines
				}
			}

			if loggedForJob {
				// Reset timeout
				timeout = time.NewTimer(settings.DeltaTimeout)
			}
		case <-timeout.C:
			respJob, err := apiClient.PipelineJob.GetApplicationJob(jobParameters, nil)
			if err != nil {
				return err
			}
			if respJob == nil {
				continue
			}
			jobSummary := respJob.Payload
			if isCompletedJob(jobSummary.Status) {
				return errorAndLogCompletedJob(jobSummary.Status, cmd)
			}
			if jobSummary.Status == "Running" {
				// Reset timeout
				timeout = time.NewTimer(settings.DeltaTimeout)
				break
			}
			getLogAttempts--
			if getLogAttempts > 0 {
				getLogAwaitingTime := int(time.Since(getLogStartTime))
				log.Print(cmd, "radix-cli", fmt.Sprintf("Nothing logged the last %d seconds. Job summary: %v. Status: %s. Contihue waiting", getLogAwaitingTime, jobSummary, jobSummary.Status), log.GetColor(0))
				break
			}
			log.Print(cmd, "radix-cli", fmt.Sprintf("Nothing logged the last %s. Job summary: %v. Status: %s. Timeout", settings.DeltaTimeout, jobSummary, jobSummary.Status), log.GetColor(0))
			return nil
		}
	}
}

func isCompletedJob(status string) bool {
	return slices.Contains(completedJobStatuses, status)
}

func errorAndLogCompletedJob(status string, cmd *cobra.Command) error {
	msg := fmt.Sprintf("job completed with status %s", status)
	if status == jobStatusFailed {
		fmt.Fprintln(cmd.OutOrStdout())
		return errors.New(msg)
	}
	log.Print(cmd, "radix-cli", msg, log.Red)
	return nil
}

func init() {
	logsCmd.AddCommand(logsJobCmd)

	logsJobCmd.Flags().StringP("application", "a", "", "Name of the application for the job")
	logsJobCmd.Flags().StringP("job", "j", "", "The job to get logs for")
	setContextSpecificPersistentFlags(logsJobCmd)
}
