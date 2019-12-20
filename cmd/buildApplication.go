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

// buildApplicationCmd represents the buildApplication command
var buildApplicationCmd = &cobra.Command{
	Use:   "build",
	Short: "Will trigger build of a Radix application",
	Long:  `Triggers build of Radix application, if branch to environment map exists for the branch in the Radix config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, _ := cmd.Flags().GetString("application")
		branch, _ := cmd.Flags().GetString("branch")
		follow, _ := cmd.Flags().GetBool("follow")

		if appName == "" || branch == "" {
			return errors.New("Application name and branch are required")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		triggerPipelineParams := application.NewTriggerPipelineBuildDeployParams()
		triggerPipelineParams.SetAppName(appName)
		triggerPipelineParams.SetPipelineParametersBuild(&models.PipelineParametersBuild{
			Branch: branch,
		})

		newJob, err := apiClient.Application.TriggerPipelineBuildDeploy(triggerPipelineParams, nil)
		if err != nil {
			return err
		}

		if follow {
			jobParameters := job.NewGetApplicationJobParams()
			jobParameters.SetAppName(appName)
			jobParameters.SetJobName(newJob.GetPayload().Name)

			newJobName := newJob.GetPayload().Name
			m := fmt.Sprintf("Building %s on branch %s with name %s", cyan(appName), yellow(branch), yellow(newJobName))
			s := `-\|/-`
			i := 0
			buildComplete := false

			refreshApplication := time.After(deltaRefreshApplication)
			tick := time.Tick(deltaRefreshOutput)

			for {
				select {
				case <-refreshApplication:
					respJob, _ := apiClient.Job.GetApplicationJob(jobParameters, nil)
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

					// Reset timer
					refreshApplication = time.After(deltaRefreshApplication)

				case <-tick:
					fmt.Fprintf(cmd.OutOrStdout(), "\r%s %c", m, s[i%len(s)])
					i++
				}

			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildApplicationCmd)
	buildApplicationCmd.Flags().StringP("application", "a", "", "Name of the application to build")
	buildApplicationCmd.Flags().StringP("branch", "b", "", "Branch to build from")
	buildApplicationCmd.Flags().BoolP("follow", "f", true, "Follow build")
}
