package completion

import (
	"strings"

	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

func ConfigContext(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	contexts := slice.FindAll(config.ValidContexts, func(context string) bool {
		return strings.HasPrefix(context, toComplete)
	})
	return contexts, cobra.ShellCompDirectiveNoFileComp
}
