package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	radixconfig "github.com/equinor/radix-cli/pkg/config"
	v1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	"github.com/equinor/radix-operator/pkg/apis/utils"
	log "github.com/sirupsen/logrus"
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
	rootCmd.PersistentFlags().BoolP("from-config", "", false, "Read and use radix config from file as context")
	rootCmd.PersistentFlags().BoolP("token-stdin", "", false, "Take the token from stdin")
	rootCmd.PersistentFlags().StringP("context", "c", "", fmt.Sprintf("Use context %s|%s|%s regardless of current context",
		radixconfig.ContextProdction, radixconfig.ContextPlayground, radixconfig.ContextDevelopment))
	rootCmd.PersistentFlags().StringP("cluster", "k", "", "Set cluster to override context")
	rootCmd.PersistentFlags().StringP("environment", "e", "prod", "The API environment to run with")
}

func getRadixApplicationFromFile() (*v1.RadixApplication, error) {
	filePath, _ := filepath.Abs("./radixconfig.yaml")
	return loadConfigFromFile(filePath)
}

// LoadConfigFromFile loads radix config from appFileName
func loadConfigFromFile(appFileName string) (*v1.RadixApplication, error) {
	radixApplication, err := utils.GetRadixApplication(appFileName)
	if err != nil {
		log.Errorf("Failed to get ra from file (%s) for app Error: %v", appFileName, err)
		return nil, err
	}

	return radixApplication, nil
}
