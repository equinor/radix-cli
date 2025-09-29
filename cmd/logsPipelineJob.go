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
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/pipeline_job"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/streaminglog"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/spf13/cobra"
)

const (
	stepStatusWaiting = "Waiting"
)

type Step struct {
	stepName      string
	containerName string
	componentName string

	isSubPipeline       bool
	pipelineRunName     string
	pipelineRunKubeName string
	pipelineStepName    string
}

func (j Step) String() string {
	if j.isSubPipeline {
		return fmt.Sprintf("%s/%s", j.stepName, j.pipelineStepName)
	}

	if j.componentName != "" {
		return fmt.Sprintf("%s/%s", j.componentName, j.stepName)
	}

	return j.stepName
}

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

		return streaminglog.New(
			cmd.ErrOrStderr(),
			getReplicasForJob(apiClient, appName, jobName),
			getLogsForJob(apiClient, appName, jobName),
			time.Second, // not used
		).StreamLogs(cmd.Context())
	},
}

func getReplicasForJob(apiClient *radixapi.Radixapi, appName, jobName string) streaminglog.GetReplicasFunc[Step] {
	return func() ([]Step, bool, error) {
		jobParameters := pipeline_job.NewGetApplicationJobParams()
		jobParameters.SetAppName(appName)
		jobParameters.SetJobName(jobName)
		respJob, err := apiClient.PipelineJob.GetApplicationJob(jobParameters, nil)
		if err != nil {
			return nil, false, err
		}

		if respJob == nil {
			return nil, false, nil
		}

		replicas := make([]Step, 0)
		for _, step := range respJob.Payload.Steps {
			if step.Status == stepStatusWaiting {
				continue
			}

			component := strings.Join(step.Components, ",")

			if step.SubPipelineTaskStep == nil {
				replicas = append(replicas, Step{
					stepName:      step.Name,
					componentName: component,
					isSubPipeline: false,
				})
				continue
			}

			replicas = append(replicas, Step{
				stepName:            step.Name,
				componentName:       component,
				containerName:       *step.SubPipelineTaskStep.KubeName,
				isSubPipeline:       true,
				pipelineRunName:     *step.SubPipelineTaskStep.PipelineRunName,
				pipelineRunKubeName: *step.SubPipelineTaskStep.KubeName,
				pipelineStepName:    *step.SubPipelineTaskStep.Name,
			})

		}

		jobCompleted := isCompletedJob(respJob.Payload.Status)
		return replicas, jobCompleted, nil
	}
}

func getLogsForJob(apiClient *radixapi.Radixapi, appName, jobName string) streaminglog.GetLogFunc[Step] {
	return func(ctx context.Context, item Step, _ time.Time, print func(text string)) error {
		if item.isSubPipeline {
			logParameters := pipeline_job.NewGetTektonPipelineRunTaskStepLogsParamsWithContext(ctx)
			logParameters.SetAppName(appName)
			logParameters.SetJobName(jobName)
			logParameters.SetPipelineRunName(item.pipelineRunName)
			logParameters.SetTaskName(item.pipelineRunKubeName)
			logParameters.SetStepName(item.pipelineStepName)
			logParameters.WithFollow(pointers.Ptr("true"))
			logParameters.WithContext(ctx)

			jobStepLog, err := apiClient.PipelineJob.GetTektonPipelineRunTaskStepLogs(logParameters, nil, streaminglog.CreateLogStreamer(print))
			if err != nil {
				return err
			}
			logLines := strings.Split(jobStepLog.Payload, "\n")
			for _, line := range logLines {
				print(line)
			}
			return nil
		}

		stepLogsParams := pipeline_job.NewGetPipelineJobStepLogsParamsWithContext(ctx)
		stepLogsParams.SetAppName(appName)
		stepLogsParams.SetJobName(jobName)
		stepLogsParams.SetStepName(item.stepName)
		stepLogsParams.WithFollow(pointers.Ptr("true"))
		stepLogsParams.WithContext(ctx)

		jobStepLog, err := apiClient.PipelineJob.GetPipelineJobStepLogs(stepLogsParams, nil, streaminglog.CreateLogStreamer(print))
		if err != nil {
			return err
		}
		logLines := strings.Split(jobStepLog.Payload, "\n")
		for _, line := range logLines {
			print(line)
		}

		return nil
	}
}

func init() {
	logsCmd.AddCommand(logsJobCmd)

	logsJobCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application for the job")
	logsJobCmd.Flags().StringP(flagnames.Job, "j", "", "The job to get logs for")

	_ = logsJobCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = logsJobCmd.RegisterFlagCompletionFunc(flagnames.Job, completion.JobCompletion)
	setContextSpecificPersistentFlags(logsJobCmd)
}
