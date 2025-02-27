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
	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/job"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/flagvalues"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/json"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// getComponentCmd represents the get component command
var getScheduledJobsCmd = &cobra.Command{
	Use:   "scheduled-job",
	Short: "Get jobs",
	Long: `Get one or all scheduled batches or jobs
  - Gets scheduled batches or jobs for a job component`,
	Example: `
	  # Get all scheduled jobs for a job component
	  rx get scheduled-job --application radix-test --environment dev --component my-job-component --batches

	  # Get a scheduled job for a job component
	  rx get scheduled-job --application radix-test --environment dev --component my-job-component --job an-unique-job-name

	  # Get all scheduled batches for a job component
	  rx get scheduled-job --application radix-test --environment dev --component my-job-component --jobs

	  # Get a scheduled batch for a job component
	  rx get scheduled-job --application radix-test --environment dev --component my-job-component --batch an-unique-batch-name
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		outputFormat, _ := cmd.Flags().GetString(flagnames.Output)
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
			return errors.New("component name is a required field")
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

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		if jobName == "" && batchName == "" {
			if allJobs {
				return getAllJobs(apiClient, appName, envName, cmpName, outputFormat)
			}
			if allBatches {
				return getAllBatches(apiClient, appName, envName, cmpName, outputFormat)
			}
			return errors.New("when options --batch and --job are not defined, options --jobs or --batches are required")
		}

		if allJobs || allBatches {
			return errors.New("options --batch and --job and options --jobs or --batches cannot be used at the same time")
		}
		if len(batchName) > 0 {
			return getBatch(apiClient, appName, envName, cmpName, batchName, outputFormat)

		}
		return getJob(apiClient, appName, envName, cmpName, jobName, outputFormat)
	},
}

func getBatch(apiClient *radixapi.Radixapi, appName, envName, cmpName, batchName, outputFormat string) error {
	parameters := job.NewGetBatchParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName).
		WithBatchName(batchName)
	resp, err := apiClient.Job.GetBatch(parameters, nil)
	if err != nil {
		return err
	}
	if outputFormat == flagvalues.OutputFormatJson {
		return json.PrettyPrintJson(resp.Payload)
	}
	if resp == nil || resp.Payload == nil {
		color.Green("Batch is not found")
		return nil
	}
	fmt.Printf("# %s environment:\n", envName)
	fmt.Println("Scheduled batch:")
	prettyPrintTextScheduledBatch(resp.Payload, "")
	return nil
}

func getJob(apiClient *radixapi.Radixapi, appName, envName, cmpName, jobName, outputFormat string) error {
	parameters := job.NewGetJobParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName).
		WithJobName(jobName)
	resp, err := apiClient.Job.GetJob(parameters, nil)
	if err != nil {
		return err
	}
	if outputFormat == flagvalues.OutputFormatJson {
		return json.PrettyPrintJson(resp.Payload)
	}
	if resp == nil || resp.Payload == nil {
		color.Green("Job is not found")
		return nil
	}
	fmt.Printf("# %s environment:\n", envName)
	jobSummary := *resp.Payload
	if len(jobSummary.BatchName) > 0 {
		fmt.Println("Scheduled batch job:")
	} else {
		fmt.Println("Scheduled single job:")
	}
	prettyPrintTextScheduledJobSummary(jobSummary, "", false)
	return nil
}

func getAllBatches(apiClient *radixapi.Radixapi, appName, envName, cmpName, outputFormat string) error {
	parameters := job.NewGetBatchesParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName)
	resp, err := apiClient.Job.GetBatches(parameters, nil)
	if err != nil {
		return err
	}
	if outputFormat == flagvalues.OutputFormatJson {
		if resp == nil {
			return json.PrettyPrintJson([]*models.ScheduledBatchSummary{})
		}
		return json.PrettyPrintJson(resp.Payload)
	}
	if resp == nil || len(resp.Payload) == 0 {
		color.Green("No batches found")
		return nil
	}
	prettyPrintTextScheduledBatches(envName, resp.Payload)
	return nil
}

func getAllJobs(apiClient *radixapi.Radixapi, appName, envName, cmpName, outputFormat string) error {
	parameters := job.NewGetJobsParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithJobComponentName(cmpName)
	resp, err := apiClient.Job.GetJobs(parameters, nil)
	if err != nil {
		return err
	}
	if outputFormat == flagvalues.OutputFormatJson {
		if resp == nil {
			return json.PrettyPrintJson([]*models.ScheduledJobSummary{})
		}
		return json.PrettyPrintJson(resp.Payload)
	}
	if resp == nil || len(resp.Payload) == 0 {
		color.Green("No jobs found")
		return nil
	}
	fmt.Printf("# %s environment:\n", envName)
	fmt.Println("Scheduled jobs:")
	prettyPrintTextScheduledJobs(resp.Payload, "    ")
	return nil
}

func prettyPrintTextScheduledBatches(envName string, batchSummaries []*models.ScheduledBatchSummary) {
	fmt.Printf("# %s environment:\n", envName)
	fmt.Println("Scheduled batches:")
	for _, batch := range batchSummaries {
		prettyPrintTextScheduledBatch(batch, "  ")
	}
	fmt.Println()
}

func prettyPrintTextScheduledBatch(batchSummary *models.ScheduledBatchSummary, indent string) {
	if batchSummary == nil {
		return
	}
	prettyPrintTextScheduledBatchSummary(*batchSummary, indent)
	fmt.Printf("%s- Jobs:\n", indent)
	prettyPrintTextScheduledJobs((*batchSummary).JobList, indent+"    ")
}

func prettyPrintTextScheduledJobs(jobSummaries []*models.ScheduledJobSummary, indent string) {
	for _, jobSummary := range jobSummaries {
		if jobSummary == nil {
			continue
		}
		prettyPrintTextScheduledJobSummary(*jobSummary, indent, true)
	}
}

func prettyPrintTextScheduledBatchSummary(batch models.ScheduledBatchSummary, indent string) {
	fmt.Printf("%sName: %s\n", indent, pointers.Val(batch.Name))
}

func prettyPrintTextScheduledJobSummary(job models.ScheduledJobSummary, indent string, itemSeparator bool) {
	itemSeparatorValue := " "
	if itemSeparator {
		itemSeparatorValue = "-"
	}
	fmt.Printf("%s%s Name: %s\n", indent, itemSeparatorValue, pointers.Val(job.Name))
	if len(job.JobID) > 0 {
		fmt.Printf("%s  Job ID: %s\n", indent, job.JobID)
	}
	if len(job.BatchName) > 0 {
		fmt.Printf("%s  Batch: %s\n", indent, job.BatchName)
	}
	fmt.Printf("%s  Deployment name: %s\n", indent, pointers.Val(job.DeploymentName))
	fmt.Printf("%s  Created: %s\n", indent, job.Created.String())
	if !job.Started.IsZero() {
		fmt.Printf("%s  Started: %s\n", indent, job.Started.String())
	}
	if !job.Ended.IsZero() {
		fmt.Printf("%s  Ended: %s\n", indent, job.Ended.String())
	}
}

func init() {
	getCmd.AddCommand(getScheduledJobsCmd)
	getScheduledJobsCmd.Flags().StringP(flagnames.Application, "a", "", "Name of an application.")
	getScheduledJobsCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of an application environment.")
	getScheduledJobsCmd.Flags().StringP(flagnames.Component, "n", "", "Name of a job component.")
	getScheduledJobsCmd.Flags().StringP(flagnames.Batch, "", "", "(Optional) The name of a scheduled batch.")
	getScheduledJobsCmd.Flags().StringP(flagnames.Job, "j", "", "(Optional) The name of a scheduled job.")
	getScheduledJobsCmd.Flags().BoolP(flagnames.Batches, "", false, "(Optional) Get all scheduled batches.")
	getScheduledJobsCmd.Flags().BoolP(flagnames.Jobs, "", false, "(Optional) Get all scheduled jobs.")
	getScheduledJobsCmd.Flags().StringP(flagnames.Output, "o", "text", "(Optional) Output format. json or not set (plain text)")
	_ = getScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = getScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = getScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	_ = getScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Batch, completion.ScheduledBatchCompletion)
	_ = getScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Job, completion.ScheduledJobCompletion)
	_ = getScheduledJobsCmd.RegisterFlagCompletionFunc(flagnames.Output, completion.Output)
	setContextSpecificPersistentFlags(getScheduledJobsCmd)
}
