package completion

import (
	"strings"

	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

func ComponentCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
	if err != nil || appName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	envName, err := cmd.Flags().GetString(flagnames.Environment)
	if err != nil || envName == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	apiClient, err := client.GetForCommand(cmd)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	GetEnvironmentParams := environment.NewGetEnvironmentParams().WithEnvName(envName).WithAppName(appName)
	resp, err := apiClient.Environment.GetEnvironment(GetEnvironmentParams, nil)
	if err != nil || resp.Payload == nil || resp.Payload.ActiveDeployment == nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	names := slice.Map(resp.Payload.ActiveDeployment.Components, func(component *models.Component) string {
		return pointers.Val(component.Name)
	})

	filteredNames := slice.FindAll(names, func(appName string) bool {
		return strings.HasPrefix(appName, toComplete)
	})

	return filteredNames, cobra.ShellCompDirectiveNoFileComp
}
