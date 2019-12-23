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

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/client/job"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const deltaRefreshApplication = 3 * time.Second
const deltaRefreshOutput = 50 * time.Millisecond

var (
	yellow = color.New(color.FgHiYellow, color.BgBlack, color.Bold).SprintFunc()
	green  = color.New(color.FgHiGreen, color.BgBlack, color.Bold).SprintFunc()
	blue   = color.New(color.FgHiBlue, color.BgBlack, color.Underline).SprintFunc()
	cyan   = color.New(color.FgCyan, color.BgBlack).SprintFunc()
	red    = color.New(color.FgHiRed, color.BgBlack).Add(color.Italic).SprintFunc()
)

// buildDeployApplicationCmd represents the buildApplication command
var buildDeployApplicationCmd = &cobra.Command{
	Use:   "build-deploy",
	Short: "Will trigger build-deploy of a Radix application",
	Long:  `Triggers build-deploy of Radix application, if branch to environment map exists for the branch in the Radix config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		branch, _ := cmd.Flags().GetString("branch")
		commitID, _ := cmd.Flags().GetString("commitID")
		follow, _ := cmd.Flags().GetBool("follow")

		if appName == nil || *appName == "" || branch == "" {
			return errors.New("Application name and branch are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineBuildDeployParams()
		triggerPipelineParams.SetAppName(*appName)
		triggerPipelineParams.SetPipelineParametersBuild(&models.PipelineParametersBuild{
			Branch:   branch,
			CommitID: commitID,
		})

		newJob, err := apiClient.Application.TriggerPipelineBuildDeploy(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		if follow {
			jobName := newJob.GetPayload().Name

			jobParameters := job.NewGetApplicationJobParams()
			jobParameters.SetAppName(*appName)
			jobParameters.SetJobName(jobName)

			fmt.Fprintf(cmd.OutOrStdout(), "\r%s", fmt.Sprintf("Building %s on branch %s with name %s", cyan(appName), yellow(branch), yellow(jobName)))

			buildComplete := false

			numLogLinesOutput := 0
			refreshApplication := time.Tick(deltaRefreshApplication)

			for {
				select {
				case <-refreshApplication:
					jobLogParameters := job.NewGetApplicationJobLogsParams()
					jobLogParameters.SetAppName(*appName)
					jobLogParameters.SetJobName(jobName)

					respJobLog, _ := apiClient.Job.GetApplicationJobLogs(jobLogParameters, nil)
					if respJobLog != nil {
						numLogLines := 0

						stepsLog := respJobLog.Payload
						for _, stepLog := range stepsLog {
							stepLogLines := strings.Split(strings.Replace(stepLog.Log, "\r\n", "\n", -1), "\n")

							for _, stepLogLine := range stepLogLines {
								if numLogLinesOutput <= numLogLines {
									fmt.Fprintf(cmd.OutOrStdout(), "\r\n%s", stepLogLine)
									numLogLinesOutput++
								}

								numLogLines++
							}
						}
					}

					respJob, _ := apiClient.Job.GetApplicationJob(jobParameters, nil)
					if respJob != nil {
						jobSummary := respJob.Payload
						if jobSummary.Status == "Succeeded" {
							fmt.Fprintf(cmd.OutOrStdout(), fmt.Sprintf("%s", green("\nBuild complete\n")))
							buildComplete = true
						} else if jobSummary.Status == "Failed" {
							fmt.Fprintf(cmd.OutOrStdout(), fmt.Sprintf("%s", red("\nBuild failed\n")))
							buildComplete = true
						}

						if buildComplete {
							return nil
						}
					}
				}

			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildDeployApplicationCmd)
	buildDeployApplicationCmd.Flags().StringP("application", "a", "", "Name of the application to build-deploy")
	buildDeployApplicationCmd.Flags().StringP("branch", "b", "", "Branch to build-deploy from")
	buildDeployApplicationCmd.Flags().StringP("commitID", "i", "", "Commit id")
	buildDeployApplicationCmd.Flags().BoolP("follow", "f", false, "Follow build-deploy")
}
