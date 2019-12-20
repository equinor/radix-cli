package cmd

import (
	"fmt"
	"os"

	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rx",
	Short: "Command line interface for Radix platform",
	Long:  `....`,
}

// Execute the top level command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("token-stdin", "", false, "Take the token from stdin")
	rootCmd.PersistentFlags().StringP("context", "c", "", fmt.Sprintf("Use context %s|%s|%s regardless of current context",
		radixconfig.ContextProdction, radixconfig.ContextPlayground, radixconfig.ContextDevelopment))
	rootCmd.PersistentFlags().StringP("cluster", "k", "", "Set cluster to override context")
	rootCmd.PersistentFlags().StringP("environment", "e", "prod", "The API environment to run with")
}
