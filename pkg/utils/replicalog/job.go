package replicalog

import (
	"context"
	"errors"
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

var ErrJobFailed = errors.New("job has failed")
var completedJobStatuses = []string{jobStatusSucceeded, jobStatusStopped, jobStatusFailed, jobStoppedNoChanges}

type JobStep struct {
	stepName      string
	containerName string
	componentName string

	isSubPipeline       bool
	pipelineRunName     string
	pipelineRunKubeName string
	pipelineStepName    string
}

func (j JobStep) Identifier() string {
	if j.isSubPipeline {
		return fmt.Sprintf("%s/%s", j.stepName, j.pipelineStepName)
	}

	if j.componentName != "" {
		return fmt.Sprintf("%s/%s", j.componentName, j.stepName)
	}

	return j.stepName
}
func (c JobStep) Created() time.Time {
	return time.Time{} // We dont care about time for pipeline jobs
}

func GetReplicasForJob(apiClient *radixapi.Radixapi, appName, jobName string) GetReplicasFunc[JobStep] {
	return func() ([]JobStep, bool, error) {
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

		replicas := make([]JobStep, 0)
		for _, step := range respJob.Payload.Steps {
			if step.Status == stepStatusWaiting {
				continue
			}

			component := strings.Join(step.Components, ",")

			if step.SubPipelineTaskStep == nil {
				replicas = append(replicas, JobStep{
					stepName:      step.Name,
					componentName: component,
					isSubPipeline: false,
				})
				continue
			}

			replicas = append(replicas, JobStep{
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
		if respJob.Payload.Status == jobStatusFailed {
			err = ErrJobFailed
		}
		return replicas, jobCompleted, err
	}
}

func GetLogsForJob(apiClient *radixapi.Radixapi, appName, jobName string) GetLogFunc[JobStep] {
	return func(ctx context.Context, item JobStep, _ time.Time, print func(text string)) error {
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
