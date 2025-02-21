package completion

import (
	"strings"

	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/spf13/cobra"
)

func Output(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	formats := slice.FindAll(config.ValidOutputs, func(format string) bool {
		return strings.HasPrefix(format, toComplete)
	})
	return formats, cobra.ShellCompDirectiveNoFileComp
}
