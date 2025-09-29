package streaminglog

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/pipeline_job"
	"github.com/equinor/radix-common/utils/pointers"
)

const (
	stepStatusWaiting   = "Waiting"
	jobStatusFailed     = "Failed"
	jobStatusSucceeded  = "Succeeded"
	jobStatusStopped    = "Stopped"
	jobStoppedNoChanges = "StoppedNoChanges"
)

var completedJobStatuses = []string{jobStatusSucceeded, jobStatusStopped, jobStatusFailed, jobStoppedNoChanges}

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

func GetReplicasForJob(apiClient *radixapi.Radixapi, appName, jobName string) GetReplicasFunc[Step] {
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

func GetLogsForJob(apiClient *radixapi.Radixapi, appName, jobName string) GetLogFunc[Step] {
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

			jobStepLog, err := apiClient.PipelineJob.GetTektonPipelineRunTaskStepLogs(logParameters, nil, CreateLogStreamer(print))
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

		jobStepLog, err := apiClient.PipelineJob.GetPipelineJobStepLogs(stepLogsParams, nil, CreateLogStreamer(print))
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

func isCompletedJob(status string) bool {
	return slices.Contains(completedJobStatuses, status)
}
