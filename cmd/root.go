package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/equinor/radix-cli/pkg/client"
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
	rootCmd.PersistentFlags().Bool("from-config", false, "Read and use radix config from file as context")
	rootCmd.PersistentFlags().Bool("token-environment", false, fmt.Sprintf("Take the token from environment variable %s", client.TokenEnvironmentName))
	rootCmd.PersistentFlags().Bool("token-stdin", false, "Take the token from stdin")
	rootCmd.PersistentFlags().StringP("context", "c", "", fmt.Sprintf("Use context %s|%s|%s regardless of current context (default production) ",
		radixconfig.ContextProdction, radixconfig.ContextPlayground, radixconfig.ContextDevelopment))
	rootCmd.PersistentFlags().String("cluster", "", "Set cluster to override context")
	rootCmd.PersistentFlags().String("api-environment", "prod", "The API api-environment to run with (default prod)")
}

func getAppNameFromConfigOrFromParameter(cmd *cobra.Command, appNameField string) (*string, error) {
	var appName string
	var err error

	fromConfig, _ := cmd.Flags().GetBool("from-config")
	if fromConfig {
		radicConfig, err := getRadixApplicationFromFile()
		if err != nil {
			return nil, err
		}

		appName = radicConfig.GetName()
	} else {
		appName, err = cmd.Flags().GetString(appNameField)
		if err != nil {
			return nil, err
		}
	}

	return &appName, nil
}

func getEnvironmentFromConfig(cmd *cobra.Command, branchName string) (*string, error) {
	var err error

	fromConfig, _ := cmd.Flags().GetBool("from-config")
	if !fromConfig {
		return nil, errors.New("--from-config is required when getting environment from branch")
	}

	radicConfig, err := getRadixApplicationFromFile()
	if err != nil {
		return nil, err
	}

	for _, environment := range radicConfig.Spec.Environments {
		if environment.Build.From != "" && environment.Build.From == branchName {
			return &environment.Name, nil
		}
	}

	return nil, fmt.Errorf("No environment found which maps to branch name `%s`", branchName)
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
