// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/pipeline_job"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/log"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

const logsJobEnabled = true

// logsJobCmd represents the logsJobCmd command
var logsJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Get logs of job",
	Long:  `Will get and follow logs of job`,
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

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		getLogsJob(cmd, apiClient, *appName, jobName)
		return nil
	},
}

func getLogsJob(cmd *cobra.Command, apiClient *apiclient.Radixapi, appName, jobName string) {
	timeout := time.NewTimer(settings.DeltaTimeout)
	refreshLog := time.Tick(settings.DeltaRefreshApplication)

	// Somtimes, even though we get delta, the log is the same as previous
	previousLogForStep := make(map[string][]string)

	for {
		select {
		case <-refreshLog:

			now := time.Now()
			sinceTime := now.Add(-settings.DeltaRefreshApplication)

			loggedForJob := false
			steps := getSteps(apiClient, appName, jobName, sinceTime)

			for i, step := range steps {
				// Somtimes, even though we get delta, the log is the same as previous
				previousLogLines := previousLogForStep[*step.Name]
				logLines := strings.Split(strings.Replace(step.Log, "\r\n", "\n", -1), "\n")
				if len(logLines) > 0 && !strings.EqualFold(logLines[0], "") {
					log.PrintLines(cmd, *step.Name, previousLogLines, logLines, log.GetColor(i))
					loggedForJob = true
					previousLogForStep[*step.Name] = logLines
				}
			}

			if loggedForJob {
				// Reset timeout
				timeout = time.NewTimer(settings.DeltaTimeout)
			}
		case <-timeout.C:
			jobParameters := pipeline_job.NewGetApplicationJobParams()
			jobParameters.SetAppName(appName)
			jobParameters.SetJobName(jobName)

			respJob, _ := apiClient.PipelineJob.GetApplicationJob(jobParameters, nil)
			if respJob != nil {
				jobSummary := respJob.Payload
				if jobSummary.Status == "Succeeded" {
					log.Print(cmd, "radix-cli", "Build complete", log.Green)
				} else if jobSummary.Status == "Failed" {
					log.Print(cmd, "radix-cli", "Build failed", log.Red)
				} else if jobSummary.Status == "Running" {
					// Reset timeout
					timeout = time.NewTimer(settings.DeltaTimeout)
					break
				} else {
					log.Print(cmd, "radix-cli", fmt.Sprintf("Nothing logged the last %s. Job summary: %v. Status: %s. Timeout", settings.DeltaTimeout, jobSummary, jobSummary.Status), log.GetColor(0))
				}
			}

			return
		}
	}
}

func getSteps(apiClient *apiclient.Radixapi, appName, jobName string, sinceTime time.Time) []*models.StepLog {
	since := strfmt.DateTime(sinceTime)
	jobLogParameters := pipeline_job.NewGetApplicationJobLogsParams()
	jobLogParameters.SetAppName(appName)
	jobLogParameters.SetJobName(jobName)
	jobLogParameters.SetSinceTime(&since)

	respJobLog, err := apiClient.PipelineJob.GetApplicationJobLogs(jobLogParameters, nil)
	if err == nil {
		return respJobLog.Payload
	}

	return nil
}

func init() {
	if logsJobEnabled {
		logsCmd.AddCommand(logsJobCmd)

		logsJobCmd.Flags().StringP("application", "a", "", "Name of the application for the job")
		logsJobCmd.Flags().StringP("job", "j", "", "The job to get logs for")
	}
}
