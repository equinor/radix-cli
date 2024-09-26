package completion

import (
	"strings"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

func EnvironmentCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
	if err != nil || appName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	apiClient, err := client.GetForCommand(cmd)
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
		return item.Name
	})
	filteredNames := slices.Filter(nil, names, func(name string) bool {
		return strings.HasPrefix(name, toComplete)
	})

	return filteredNames, cobra.ShellCompDirectiveNoFileComp
}
