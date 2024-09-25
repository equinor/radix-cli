package completion

import (
	"strings"

	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/equinor/radix-common/utils/slice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

func ComponentCompletion(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	apiClient, err := client.GetForCommand(cmd)
	if err != nil {
		log.Warn(err)
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	params := component.NewComponentsParams()
	resp, err := apiClient.Component.Components(params, nil)

	componentNamesNames := slice.Map(resp.Payload, func(component *models.Component) string {
		return pointers.Val(component.Name)
	})
	components := slices.Filter(nil, componentNamesNames, func(appName string) bool {
		return strings.HasPrefix(appName, toComplete)
	})

	return components, cobra.ShellCompDirectiveNoFileComp
}
