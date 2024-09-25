package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/equinor/radix-cli/pkg/flagnames"
	v1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	"github.com/equinor/radix-operator/pkg/apis/utils"
	"github.com/spf13/cobra"
)

func GetAppNameFromConfigOrFromParameter(cmd *cobra.Command, appNameField string) (*string, error) {
	var appName string
	var err error

	fromConfig, _ := cmd.Flags().GetBool(flagnames.FromConfig)
	if fromConfig {
		radicConfig, err := GetRadixApplicationFromFile()
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

func GetEnvironmentFromConfig(cmd *cobra.Command, branchName string) (*string, error) {
	var err error

	fromConfig, _ := cmd.Flags().GetBool(flagnames.FromConfig)
	if !fromConfig {
		return nil, errors.New("--from-config is required when getting environment from branch")
	}

	cmd.SilenceUsage = true

	radicConfig, err := GetRadixApplicationFromFile()
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

func GetRadixApplicationFromFile() (*v1.RadixApplication, error) {
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

func GetStringFromFlagValueOrFlagFile(cmd *cobra.Command, valueFlag, fileNameFlag string) (string, error) {
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
