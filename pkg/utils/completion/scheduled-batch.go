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

func ScheduledBatchCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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

	getBatchesParams := job.NewGetBatchesParams().WithAppName(appName).WithEnvName(envName).WithJobComponentName(componentName)
	resp, err := apiClient.Job.GetBatches(getBatchesParams, nil)
	if err != nil || resp.Payload == nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	batchNames := slice.Map(resp.Payload, func(batchSummary *models.ScheduledBatchSummary) string {
		if batchSummary == nil {
			return ""
		}
		return pointers.Val(batchSummary.Name)
	})

	filteredNames := slice.FindAll(batchNames, func(batchName string) bool {
		return strings.HasPrefix(batchName, toComplete)
	})

	return filteredNames, cobra.ShellCompDirectiveNoFileComp
}
