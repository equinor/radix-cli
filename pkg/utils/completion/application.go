package completion

import (
	"strings"

	"github.com/equinor/radix-cli/generated-client/client/platform"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-common/utils/slice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

const KnownApps = "known_apps"

func ApplicationCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if content, ok := config.GetCache[[]string]("known_apps"); ok {
		return content, cobra.ShellCompDirectiveNoFileComp
	}

	apiClient, err := client.GetForCommand(cmd)
	if err != nil {
		log.Warn(err)
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	showApplicationParams := platform.NewShowApplicationsParams()
	resp, err := apiClient.Platform.ShowApplications(showApplicationParams, nil)
	if err != nil {
		log.Warn(err)
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	appNames := slice.Map(resp.Payload, func(app *models.ApplicationSummary) string {
		return app.Name
	})
	applications := slices.Filter(nil, appNames, func(appName string) bool {
		return strings.HasPrefix(appName, toComplete)
	})
	config.SetCache(KnownApps, applications, config.DefaultCacheDuration)

	return applications, cobra.ShellCompDirectiveNoFileComp
}
