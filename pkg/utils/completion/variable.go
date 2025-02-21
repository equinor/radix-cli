package completion

import (
	"strings"

	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

func VariableCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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

	params := environment.NewGetEnvironmentParams().WithEnvName(envName).WithAppName(appName)
	resp, err := apiClient.Environment.GetEnvironment(params, nil)
	if err != nil || resp.Payload == nil || resp.Payload.ActiveDeployment == nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	var names []string
	for _, comp := range resp.Payload.ActiveDeployment.Components {
		if comp.Name != nil && *comp.Name == componentName {
			for variable := range comp.Variables {
				names = append(names, variable)
			}
		}
	}

	filteredNames := slice.FindAll(names, func(name string) bool {
		return strings.HasPrefix(name, toComplete) && !strings.HasPrefix(name, "RADIX_")
	})

	return filteredNames, cobra.ShellCompDirectiveNoFileComp
}
