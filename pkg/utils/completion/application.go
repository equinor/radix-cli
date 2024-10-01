package completion

import (
	"strings"

	"github.com/equinor/radix-cli/generated-client/client/platform"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

const knownApps = "known_apps"

func ApplicationCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if content, ok := config.GetCache[[]string]("known_apps"); ok {
		return content, cobra.ShellCompDirectiveNoFileComp
	}

	apiClient, err := client.GetForCommand(cmd)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	showApplicationParams := platform.NewShowApplicationsParams()
	resp, err := apiClient.Platform.ShowApplications(showApplicationParams, nil)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	appNames := slice.Map(resp.Payload, func(app *models.ApplicationSummary) string {
		return app.Name
	})
	applications := slice.FindAll(appNames, func(appName string) bool {
		return strings.HasPrefix(appName, toComplete)
	})
	config.SetCache(knownApps, applications, config.DefaultCacheDuration)

	return applications, cobra.ShellCompDirectiveNoFileComp
}

func UpdateAppNamesCache(appNames []string) {
	config.SetCache(knownApps, appNames, config.DefaultCacheDuration)
}
