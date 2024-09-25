package completion

import (
	"strings"

	"github.com/equinor/radix-cli/pkg/config"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

func ConfigContext(cmd *cobra.Command, args []string, toComplete string) (applications []string, shell cobra.ShellCompDirective) {
	shell = cobra.ShellCompDirectiveNoFileComp

	applications = slices.Filter(nil, config.ValidContexts, func(appName string) bool {
		return strings.HasPrefix(appName, toComplete)
	})
	return
}
