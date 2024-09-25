package completion

import (
	"strings"

	"github.com/equinor/radix-cli/pkg/config"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

func ConfigContext(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	contexts := slices.Filter(nil, config.ValidContexts, func(appName string) bool {
		return strings.HasPrefix(appName, toComplete)
	})
	return contexts, cobra.ShellCompDirectiveNoFileComp
}
