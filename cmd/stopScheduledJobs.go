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

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/job"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// stopComponentCmd represents the stop component command
var stopScheduledJobsCmd = &cobra.Command{
	Use:   "scheduled-job",
	Short: "Stop jobs",
	Long: `Stop one or all scheduled batches or jobs
  - Stops scheduled batches or jobs for a job component`,
	Example: `
	  # Stop all scheduled jobs for a job component
	  rx stop scheduled-job --application radix-test --environment dev --component my-job-component --jobs

	  # Stop a scheduled job for a job component
	  rx stop scheduled-job --application radix-test --environment dev --component my-job-component --job an-unique-job-name

	  # Stop all scheduled batches for a job component
	  rx stop scheduled-job --application radix-test --environment dev --component my-job-component --batches

	  # Stop a scheduled batch for a job component
	  rx stop scheduled-job --application radix-test --environment dev --component my-job-component --batch an-unique-batch-name

	  # Stop all scheduled batches and jobs for a job component
	  rx stop scheduled-job --application radix-test --environment dev --component my-job-component --all

	  # Stop all scheduled batches and jobs for an environment
	  rx stop scheduled-job --application radix-test --environment dev --all
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}
		envName, err := cmd.Flags().GetString(flagnames.Environment)
		if err != nil || appName == "" || envName == "" {
			return errors.New("environment name and application name are required fields")
		}
		cmpName, err := cmd.Flags().GetString(flagnames.Component)
		if err != nil {
			return err
		}
		batchName, err := cmd.Flags().GetString(flagnames.Batch)
		if err != nil {
			return err
		}
		jobName, err := cmd.Flags().GetString(flagnames.Job)
		if err != nil {
			return err
		}
		if len(batchName) > 0 && len(jobName) > 0 {
			return errors.New("options --batch and --job cannot be used at the same time")
		}

		allJobs, err := cmd.Flags().GetBool(flagnames.Jobs)
		if err != nil {
			return err
		}
		allBatches, err := cmd.Flags().GetBool(flagnames.Batches)
		if err != nil {
			return err
		}
		allBatchesAndJobs, err := cmd.Flags().GetBool(flagnames.All)
		if err != nil {
			return err
		}
		if allBatchesAndJobs && (allJobs || allBatches) {
			return errors.New("an option --all cannot be used with options --batches and --jobs at the same time")
		}

		if len(cmpName) == 0 && !allBatchesAndJobs {
			return errors.New("when a component name is not defined, an option --all is required to stop all scheduled batches or jobs in an environment")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		if jobName == "" && batchName == "" {
			if allJobs || allBatchesAndJobs {
				return stopAllJobs(apiClient, appName, envName, cmpName)
			}
			if allBatches || allBatchesAndJobs {
				return stopAllBatches(apiClient, appName, envName, cmpName)
			}
			return errors.New("when options --batch and --job are not defined, options --all, --jobs or --batches are required")
		}

		if allJobs || allBatches || allBatchesAndJobs {
			return errors.New("options --batch and --job and options --all, --jobs or --batches cannot be used at the same time")
		}
		if len(batchName) > 0 {
			return stopBatch(apiClient, appName, envName, cmpName, batchName)

		}
		return stopJob(apiClient, appName, envName, cmpName, jobName)
	},
}

func stopBatch(apiClient *radixapi.Radixapi, appName, envName, cmpName, batchName string) error {
	parameters := job.NewStopBatchParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName).
		WithBatchName(batchName)
	_, err := apiClient.Job.StopBatch(parameters, nil)
	return err
}

func stopJob(apiClient *radixapi.Radixapi, appName, envName, cmpName, jobName string) error {
	parameters := job.NewStopJobParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName).
		WithJobName(jobName)
	_, err := apiClient.Job.StopJob(parameters, nil)
	return err
}

func stopAllBatches(apiClient *radixapi.Radixapi, appName, envName, cmpName string) error {
	parameters := job.NewStopAllBatchesParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName)
	_, err := apiClient.Job.StopAllBatches(parameters, nil)
	return err
}

func stopAllJobs(apiClient *radixapi.Radixapi, appName, envName, cmpName string) error {
	parameters := job.NewStopAllJobsParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName)
	_, err := apiClient.Job.StopAllJobs(parameters, nil)
	return err
}

func init() {
	stopCmd.AddCommand(stopScheduledJobsCmd)
	stopScheduledJobsCmd.Flags().StringP(flagnames.Application, "a", "", "Name of an application.")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of an application environment.")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Component, "n", "", "Name of a job component.")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Batch, "", "", "The name of a scheduled batch.")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Job, "j", "", "The name of a scheduled job.")
	stopScheduledJobsCmd.Flags().BoolP(flagnames.Batches, "", false, "Stop all scheduled batches.")
	stopScheduledJobsCmd.Flags().BoolP(flagnames.Jobs, "", false, "Stop all scheduled jobs.")
	stopScheduledJobsCmd.Flags().BoolP(flagnames.All, "", false, "Stop all scheduled batches and jobs.")
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Batch, completion.ScheduledBatchCompletion)
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Job, completion.ScheduledJobCompletion)
	setContextSpecificPersistentFlags(stopScheduledJobsCmd)
}
