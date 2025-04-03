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

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var names []string
		if envName == "" {
			names = getAllDeploymentNames(appName, apiClient)
		} else {
			names = getEnvironmentDeploymentNames(appName, envName, apiClient)
		}

		filteredItems := slice.FindAll(names, func(name string) bool {
			return strings.HasPrefix(name, toComplete)
		})

		return filteredItems, cobra.ShellCompDirectiveNoFileComp
	}
}

func CreateDeploymentCommitIDCompletion(environmentFlagName string, envRequired bool) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil || appName == "" {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		envName, err := cmd.Flags().GetString(environmentFlagName)
		if err != nil || (envRequired && envName == "") {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var names []string
		if envName == "" {
			names = getAllDeploymentCommitIDs(appName, apiClient)
		} else {
			names = getEnvironmentDeploymentCommitIds(appName, envName, apiClient)
		}

		filteredItems := slice.FindAll(names, func(name string) bool {
			return strings.HasPrefix(name, toComplete)
		})

		return filteredItems, cobra.ShellCompDirectiveNoFileComp
	}
}

func getEnvironmentDeploymentCommitIds(appName, envName string, apiClient *radixapi.Radixapi) []string {
	deploymentSummaries := getEnvironmentDeployments(appName, envName, apiClient)
	return getDeploymentSummariesWithCommitIds(deploymentSummaries)
}

func getEnvironmentDeploymentNames(appName, envName string, apiClient *radixapi.Radixapi) []string {
	deploymentSummaries := getEnvironmentDeployments(appName, envName, apiClient)
	return slice.Map(deploymentSummaries, func(item *models.DeploymentSummary) string {
		return pointers.Val(item.Name)
	})
}

func getEnvironmentDeployments(appName string, envName string, apiClient *radixapi.Radixapi) []*models.DeploymentSummary {
	params := environment.NewGetEnvironmentParams().WithEnvName(envName).WithAppName(appName)
	resp, err := apiClient.Environment.GetEnvironment(params, nil)
	if err != nil || resp.Payload == nil {
		return nil
	}
	return resp.Payload.Deployments
}

func getAllDeploymentCommitIDs(appName string, apiClient *radixapi.Radixapi) []string {
	deploymentSummaries, err := getAllDeployments(appName, apiClient)
	if err != nil {
		return nil
	}
	return getDeploymentSummariesWithCommitIds(deploymentSummaries)
}

func getDeploymentSummariesWithCommitIds(deploymentSummaries []*models.DeploymentSummary) []string {
	deploymentSummariesWithCommitId := slice.FindAll(deploymentSummaries, func(summary *models.DeploymentSummary) bool {
		return summary.GitCommitHash != ""
	})
	return slice.Map(deploymentSummariesWithCommitId, func(item *models.DeploymentSummary) string {
		return item.GitCommitHash
	})
}

func getAllDeploymentNames(appName string, apiClient *radixapi.Radixapi) []string {
	deploymentSummaries, err := getAllDeployments(appName, apiClient)
	if err != nil {
		return nil
	}

	return slice.Map(deploymentSummaries, func(item *models.DeploymentSummary) string {
		return pointers.Val(item.Name)
	})
}

func getAllDeployments(appName string, apiClient *radixapi.Radixapi) ([]*models.DeploymentSummary, error) {
	params := application.NewGetDeploymentsParams().WithAppName(appName)
	resp, err := apiClient.Application.GetDeployments(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
