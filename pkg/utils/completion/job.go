package completion

import (
	"strings"

	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

func JobCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
	if err != nil || appName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	apiClient, err := client.GetRadixApiForCommand(cmd)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	params := application.NewGetApplicationParams().WithAppName(appName)
	resp, err := apiClient.Application.GetApplication(params, nil)
	if err != nil || resp.Payload == nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	names := slice.Map(resp.Payload.Jobs, func(component *models.JobSummary) string {
		if component == nil {
			return ""
		}
		return *component.Name
	})

	filteredNames := slice.FindAll(names, func(appName string) bool {
		return strings.HasPrefix(appName, toComplete)
	})

	return filteredNames, cobra.ShellCompDirectiveNoFileComp
}
