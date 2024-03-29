package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	"github.com/equinor/radix-cli/pkg/client"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/settings"
	v1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	"github.com/equinor/radix-operator/pkg/apis/utils"
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
	return cmd.PersistentFlags().Bool(flagnames.Verbose, false, "Verbose output")
}

func setContextSpecificPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool(flagnames.FromConfig, false, "Read and use radix config from file as context")
	cmd.PersistentFlags().Bool(flagnames.TokenEnvironment, false, fmt.Sprintf("Take the token from environment variable %s", client.TokenEnvironmentName))
	cmd.PersistentFlags().Bool(flagnames.TokenStdin, false, "Take the token from stdin")
	setContextPersistentFlags(cmd)
	cmd.PersistentFlags().String(flagnames.Cluster, "", "Set cluster to override context")
	cmd.PersistentFlags().String(flagnames.ApiEnvironment, "prod", "The API api-environment to run with (default prod)")
	setVerbosePersistentFlag(cmd)
}

func setContextPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(flagnames.Context, "c", "", fmt.Sprintf("Use context %s|%s|%s|%s regardless of current context (default production) ",
		radixconfig.ContextPlatform, radixconfig.ContextPlatform2, radixconfig.ContextPlayground, radixconfig.ContextDevelopment))
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

func getAppNameFromConfigOrFromParameter(cmd *cobra.Command, appNameField string) (*string, error) {
	var appName string
	var err error

	fromConfig, _ := cmd.Flags().GetBool(flagnames.FromConfig)
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

	fromConfig, _ := cmd.Flags().GetBool(flagnames.FromConfig)
	if !fromConfig {
		return nil, errors.New("--from-config is required when getting environment from branch")
	}

	cmd.SilenceUsage = true

	radicConfig, err := getRadixApplicationFromFile()
	if err != nil {
		return nil, err
	}

	for _, environment := range radicConfig.Spec.Environments {
		if environment.Build.From != "" && environment.Build.From == branchName {
			return &environment.Name, nil
		}
	}

	return nil, fmt.Errorf("no environment found which maps to branch name `%s`", branchName)
}

func getRadixApplicationFromFile() (*v1.RadixApplication, error) {
	filePath, _ := filepath.Abs("./radixconfig.yaml")
	return loadConfigFromFile(filePath)
}

// LoadConfigFromFile loads radix config from appFileName
func loadConfigFromFile(appFileName string) (*v1.RadixApplication, error) {
	radixApplication, err := utils.GetRadixApplicationFromFile(appFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to get ra from file (%s) for app Error: %v", appFileName, err)
	}

	return radixApplication, nil
}

func getStringFromFlagValueOrFlagFile(cmd *cobra.Command, valueFlag, fileNameFlag string) (string, error) {
	fileName, err := cmd.Flags().GetString(fileNameFlag)
	if err != nil {
		return "", err
	}
	if len(fileName) > 0 {
		fileContent, err := os.ReadFile(fileName)
		return string(fileContent), err
	}

	return cmd.Flags().GetString(valueFlag)
}
