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
	"github.com/equinor/radix-cli/generated-client/client/job"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/utils/log"
	"github.com/spf13/cobra"
)

// followJobCmd represents the followJobCmd command
var followJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Will follow a job",
	Long:  `Will follow a job`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("Application name is required")
		}

		jobName, _ := cmd.Flags().GetString("job")

		if jobName == "" {
			return errors.New("`job` is required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		followJob(cmd, apiClient, *appName, jobName)
		return nil
	},
}

func followJob(cmd *cobra.Command, apiClient *apiclient.Radixapi, appName, jobName string) {
	timeout := time.NewTimer(deltaTimeout)
	refreshLog := time.Tick(deltaRefreshApplication)
	loggedForStep := make(map[string]int)

	for {
		select {
		case <-refreshLog:

			loggedForJob := false
			steps := getSteps(apiClient, appName, jobName)

			for i, step := range steps {
				totalLinesLogged := 0

				if _, contained := loggedForStep[*step.Name]; contained {
					totalLinesLogged = loggedForStep[*step.Name]
				}

				logLines := strings.Split(strings.Replace(step.Log, "\r\n", "\n", -1), "\n")
				logged := log.From(cmd, *step.Name, totalLinesLogged, logLines, log.GetColor(i))

				totalLinesLogged += logged
				loggedForStep[*step.Name] = totalLinesLogged

				if logged > 0 {
					loggedForJob = true
				}
			}

			if loggedForJob {
				// Reset timeout
				timeout = time.NewTimer(deltaTimeout)
			}
		case <-timeout.C:
			jobParameters := job.NewGetApplicationJobParams()
			jobParameters.SetAppName(appName)
			jobParameters.SetJobName(jobName)

			respJob, _ := apiClient.Job.GetApplicationJob(jobParameters, nil)
			if respJob != nil {
				jobSummary := respJob.Payload
				if jobSummary.Status == "Succeeded" {
					log.Print(cmd, "radix-cli", "Build complete", log.Green)
				} else if jobSummary.Status == "Failed" {
					log.Print(cmd, "radix-cli", "Build failed", log.Red)
				} else {
					log.Print(cmd, "radix-cli", fmt.Sprintf("Nothing logged the last %s. Timeout", deltaTimeout), log.GetColor(0))
				}
			}

			return
		}
	}
}

func getSteps(apiClient *apiclient.Radixapi, appName, jobName string) []*models.StepLog {
	jobLogParameters := job.NewGetApplicationJobLogsParams()
	jobLogParameters.SetAppName(appName)
	jobLogParameters.SetJobName(jobName)

	respJobLog, err := apiClient.Job.GetApplicationJobLogs(jobLogParameters, nil)
	if err == nil {
		return respJobLog.Payload
	}

	return nil
}

func init() {
	followJobCmd.Flags().StringP("application", "a", "", "Name of the application owning the component")
	followJobCmd.Flags().StringP("job", "j", "", "The job to follow")
}
