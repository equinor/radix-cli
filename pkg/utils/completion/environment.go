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

func EnvironmentCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
	if err != nil || appName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	apiClient, err := client.GetRadixApiForCommand(cmd)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	getApplicationParams := application.NewGetApplicationParams()
	getApplicationParams.SetAppName(appName)
	resp, err := apiClient.Application.GetApplication(getApplicationParams, nil)
	if err != nil || resp.Payload == nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	names := slice.Map(resp.Payload.Environments, func(item *models.EnvironmentSummary) string {
		if item == nil {
			return ""
		}
		return *item.Name
	})
	filteredNames := slice.FindAll(names, func(name string) bool {
		return strings.HasPrefix(name, toComplete)
	})

	return filteredNames, cobra.ShellCompDirectiveNoFileComp
}
