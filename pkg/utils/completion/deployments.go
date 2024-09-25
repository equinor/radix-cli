package completion

import (
	"strings"

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

func CreateDeploymentCompletion(environmentFlagName string, envRequired bool) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil || appName == "" {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		envName, err := cmd.Flags().GetString(environmentFlagName)
		if err != nil || (envRequired && envName == "") {
			return []string{"missing-env"}, cobra.ShellCompDirectiveNoFileComp
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return []string{err.Error()}, cobra.ShellCompDirectiveNoFileComp
		}

		var names []string
		if envName == "" {
			names = getAllDeployments(appName, apiClient)
		} else {
			names = getEnvironmentDeployments(appName, envName, apiClient)
		}

		filteredNames := slices.Filter(nil, names, func(name string) bool {
			return strings.HasPrefix(name, toComplete)
		})

		return filteredNames, cobra.ShellCompDirectiveNoFileComp
	}
}

func getEnvironmentDeployments(appName, envName string, apiClient *apiclient.Radixapi) []string {
	params := environment.NewGetEnvironmentParams().WithEnvName(envName).WithAppName(appName)
	resp, err := apiClient.Environment.GetEnvironment(params, nil)
	if err != nil {
		return nil
	}

	return slice.Map(resp.Payload.Deployments, func(item *models.DeploymentSummary) string {
		return *item.Name
	})
}

func getAllDeployments(appName string, apiClient *apiclient.Radixapi) []string {
	params := application.NewGetDeploymentsParams().WithAppName(appName)
	resp, err := apiClient.Application.GetDeployments(params, nil)
	if err != nil {
		return nil
	}

	return slice.Map(resp.Payload, func(item *models.DeploymentSummary) string {
		return *item.Name
	})
}
