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
	Long: `Stop one or all scheduled jobs
  - Stops scheduled jobs for a job component`,
	Example: `
	  # Stop all scheduled jobs for a job component
	  rx stop scheduled-job --application radix-test --environment dev --component jobcomponent --all
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
			return errors.New("job component name is a required field")
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
			return errors.New("batch and job names cannot be used at the same time")
		}
		allJobs, err := cmd.Flags().GetBool(flagnames.AllJobs)
		if err != nil {
			return err
		}
		allBatches, err := cmd.Flags().GetBool(flagnames.AllBatches)
		if err != nil {
			return err
		}
		if allJobs && allBatches {
			return errors.New("options --all-jobs and --all-batches cannot be used at the same time")
		}

		cmd.SilenceUsage = true

		if jobName == "" && batchName == "" {
			if allJobs {
				return stopAllJobs(cmd, appName, envName, cmpName)
			} else if allBatches {
				return stopAllBatches(cmd, appName, envName, cmpName)
			}
			return errors.New("when batch and/or job name are not specified, options --all-jobs or --all-batches are expected")
		}
		if allJobs || allBatches {
			return errors.New("batch and job names and the option --all-jobs or --all-batches cannot be used at the same time")
		}
		if len(batchName) > 0 {
			return stopBatchJob(cmd, appName, envName, cmpName, batchName)
		}
		return stopSingleJob(cmd, appName, envName, cmpName, jobName)
	},
}

func stopBatchJob(cmd *cobra.Command, appName, envName, cmpName, batchName string) error {
	parameters := job.NewStopBatchParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName).
		WithBatchName(batchName)

	apiClient, err := client.GetRadixApiForCommand(cmd)
	if err != nil {
		return err
	}
	_, err = apiClient.Job.StopBatch(parameters, nil)
	return err
}

func stopSingleJob(cmd *cobra.Command, appName, envName, cmpName, jobName string) error {
	parameters := job.NewStopJobParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName).
		WithJobName(jobName)

	apiClient, err := client.GetRadixApiForCommand(cmd)
	if err != nil {
		return err
	}
	_, err = apiClient.Job.StopJob(parameters, nil)
	return err
}

func stopAllBatches(cmd *cobra.Command, appName, envName, cmpName string) error {
	parameters := job.NewStopAllBatchesParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName)
	apiClient, err := client.GetRadixApiForCommand(cmd)
	if err != nil {
		return err
	}
	_, err = apiClient.Job.StopAllBatches(parameters, nil)
	return err
}

func stopAllJobs(cmd *cobra.Command, appName, envName, cmpName string) error {
	parameters := job.NewStopAllJobsParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName)

	apiClient, err := client.GetRadixApiForCommand(cmd)
	if err != nil {
		return err
	}

	_, err = apiClient.Job.StopAllJobs(parameters, nil)
	return err
}

func init() {
	stopCmd.AddCommand(stopScheduledJobsCmd)
	stopScheduledJobsCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Component, "n", "", "Name of the job component")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Batch, "", "", "The name of the scheduled batch")
	stopScheduledJobsCmd.Flags().StringP(flagnames.Job, "j", "", "The name of the scheduled single job")
	stopScheduledJobsCmd.Flags().BoolP(flagnames.AllJobs, "", false, "Stop all jobs for the job component")
	stopScheduledJobsCmd.Flags().BoolP(flagnames.AllBatches, "", false, "Stop all batches for the job component")
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = stopScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	setContextSpecificPersistentFlags(stopScheduledJobsCmd)
}
