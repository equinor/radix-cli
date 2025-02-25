package completion

import (
	"github.com/equinor/radix-cli/generated/radixapi/client/job"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
	"strings"
)

func ScheduledJobCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
	if err != nil || appName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	envName, err := cmd.Flags().GetString(flagnames.Environment)
	if err != nil || envName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	componentName, err := cmd.Flags().GetString(flagnames.Component)
	if err != nil || componentName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	apiClient, err := client.GetRadixApiForCommand(cmd)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	getJobsParams := job.NewGetJobsParams().WithAppName(appName).WithEnvName(envName).WithJobComponentName(componentName)
	resp, err := apiClient.Job.GetJobs(getJobsParams, nil)
	if err != nil || resp.Payload == nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	jobNames := slice.Map(resp.Payload, func(jobSummary *models.ScheduledJobSummary) string {
		if jobSummary == nil {
			return ""
		}
		return pointers.Val(jobSummary.Name)
	})

	filteredNames := slice.FindAll(jobNames, func(jobName string) bool {
		return strings.HasPrefix(jobName, toComplete)
	})

	return filteredNames, cobra.ShellCompDirectiveNoFileComp
}
