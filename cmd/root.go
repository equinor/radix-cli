package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/equinor/radix-cli/pkg/client"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Version = "dev"
var rootLongHelp = strings.TrimSpace(`
A command line interface which allows you to interact with the Radix platform through automation.
`)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "rx",
	Short:   "Command line interface for Radix platform",
	Long:    rootLongHelp,
	Version: Version,
}

// Execute the top level command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	if Version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok {
			Version = info.Main.Version
			rootCmd.Version = info.Main.Version
		}
	}
}

func setVerbosePersistentFlag(cmd *cobra.Command) *bool {
	return cmd.PersistentFlags().Bool(flagnames.Verbose, false, "(Optional) Verbose output")
}

func setContextSpecificPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool(flagnames.FromConfig, false, "(Optional) Read and use radix config from file as context")
	cmd.PersistentFlags().Bool(flagnames.TokenEnvironment, false, fmt.Sprintf("(Optional) Take the token from environment variable %s", client.TokenEnvironmentName))
	cmd.PersistentFlags().Bool(flagnames.TokenStdin, false, "(Optional) Take the token from stdin")
	setContextPersistentFlags(cmd)
	cmd.PersistentFlags().String(flagnames.Cluster, "", "(Optional) Set cluster to override context")
	cmd.PersistentFlags().String(flagnames.ApiEnvironment, "prod", "(Optional) The API api-environment to run with (default prod)")
	setVerbosePersistentFlag(cmd)
}

func setContextPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(flagnames.Context, "c", "", fmt.Sprintf("(Optional) Use context %s regardless of current context (default production) ", strings.Join(radixconfig.ValidContexts, "|")))
	_ = cmd.RegisterFlagCompletionFunc(flagnames.Context, completion.ConfigContext)
}

// CheckFnNew The prototype function for any check function
type checkFn func() bool

func awaitReconciliation(checkFunc checkFn) bool {
	timeout := time.NewTimer(settings.DeltaTimeout)
	checkReconciliation := time.Tick(settings.DeltaRefreshApplication)

	for {
		select {
		case <-checkReconciliation:
			success := checkFunc()
			if success {
				return true
			}

			log.Info("Radix still not appear to be reconciled")

		case <-timeout.C:
			log.Info("Time out checking reconciled state")
			return false
		}
	}
}
