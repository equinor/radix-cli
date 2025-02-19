package completion

import (
	"strings"

	radixapi "github.com/equinor/radix-cli/generated/radixapi/client"
	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/client/environment"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

func CreateDeploymentCompletion(environmentFlagName string, envRequired bool) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil || appName == "" {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		envName, err := cmd.Flags().GetString(environmentFlagName)
		if err != nil || (envRequired && envName == "") {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var names []string
		if envName == "" {
			names = getAllDeployments(appName, apiClient)
		} else {
			names = getEnvironmentDeployments(appName, envName, apiClient)
		}

		filteredNames := slice.FindAll(names, func(name string) bool {
			return strings.HasPrefix(name, toComplete)
		})

		return filteredNames, cobra.ShellCompDirectiveNoFileComp
	}
}

func getEnvironmentDeployments(appName, envName string, apiClient *radixapi.Radixapi) []string {
	params := environment.NewGetEnvironmentParams().WithEnvName(envName).WithAppName(appName)
	resp, err := apiClient.Environment.GetEnvironment(params, nil)
	if err != nil || resp.Payload == nil {
		return nil
	}

	return slice.Map(resp.Payload.Deployments, func(item *models.DeploymentSummary) string {
		return pointers.Val(item.Name)
	})
}

func getAllDeployments(appName string, apiClient *radixapi.Radixapi) []string {
	params := application.NewGetDeploymentsParams().WithAppName(appName)
	resp, err := apiClient.Application.GetDeployments(params, nil)
	if err != nil {
		return nil
	}

	return slice.Map(resp.Payload, func(item *models.DeploymentSummary) string {
		return pointers.Val(item.Name)
	})
}
