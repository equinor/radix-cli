// Copyright © 2023
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/client/platform"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/spf13/cobra"
)

// createApplicationCmd represents the create application command
var createApplicationCmd = &cobra.Command{
	Use:     "application",
	Short:   "Create application",
	Long:    "Creates a Radix application in the cluster",
	Example: `rx create application --application your-application-name --repository https://github.com/your-repository --config-branch main --ad-groups abcdef-1234-5678-9aaa-abcdefgf --reader-ad-groups=23456789--9123-4567-8901-23456701 --shared-secret someSecretPhrase12345 --acknowledge-warnings --configuration-item "YOUR PROJECT CONFIG ITEM" --context playground`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		repository, _ := cmd.Flags().GetString(flagnames.Repository)
		sharedSecret, _ := cmd.Flags().GetString(flagnames.SharedSecret)
		configBranch, _ := cmd.Flags().GetString(flagnames.ConfigBranch)
		configFile, _ := cmd.Flags().GetString(flagnames.ConfigFile)
		configurationItem, _ := cmd.Flags().GetString(flagnames.ConfigurationItem)
		acknowledgeWarnings, err := cmd.Flags().GetBool(flagnames.AcknowledgeWarnings)
		if err != nil {
			println(fmt.Sprintf("invalid argument %s: %v", "acknowledge-warnings", err))
			return err
		}

		if appName == "" || repository == "" || configBranch == "" || configurationItem == "" {
			return errors.New("application name, repository, configuration item and config branch are required fields")
		}

		adGroups, _ := cmd.Flags().GetStringSlice(flagnames.AdminADGroups)
		readerAdGroups, _ := cmd.Flags().GetStringSlice(flagnames.ReaderADGroups)

		cmd.SilenceUsage = true

		registerApplicationParams := platform.NewRegisterApplicationParams()
		registerApplicationParams.SetApplicationRegistration(&models.ApplicationRegistrationRequest{
			AcknowledgeWarnings: acknowledgeWarnings,
			ApplicationRegistration: &models.ApplicationRegistration{
				AdGroups:            adGroups,
				ConfigBranch:        &configBranch,
				ConfigurationItem:   configurationItem,
				Name:                &appName,
				RadixConfigFullName: configFile,
				ReaderAdGroups:      readerAdGroups,
				Repository:          &repository,
				SharedSecret:        &sharedSecret,
			},
		})

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		resp, err := apiClient.Platform.RegisterApplication(registerApplicationParams, nil)

		if err != nil {
			return err
		}
		registrationUpsertResponse := resp.Payload
		if len(registrationUpsertResponse.Warnings) > 0 {
			println("Warnings:")
			for _, warning := range registrationUpsertResponse.Warnings {
				println(fmt.Sprintf("- %s", warning))
			}
			return fmt.Errorf("if you agree to proceed with warnings - please add an option --acknowledge-warnings")
		}
		if registrationUpsertResponse.ApplicationRegistration == nil {
			return errors.New("unspecified error")

		}
		deployKeyAndSecretParams := application.NewGetDeployKeyAndSecretParams()
		deployKeyAndSecretParams.SetAppName(appName)
		getRadixRegistrationNoAccessErrorCount := 3
		getRadixRegistrationNoAccessErrorPause := 2 * time.Second
		for {
			deployKeyResp, err := apiClient.Application.GetDeployKeyAndSecret(deployKeyAndSecretParams, nil)
			if err != nil {
				getRadixRegistrationNoAccessErrorCount--
				if getRadixRegistrationNoAccessErrorCount == 0 {
					return fmt.Errorf("error getting public deploy key: %w", err)
				}
				time.Sleep(getRadixRegistrationNoAccessErrorPause) // Sleep before trying again
				continue
			}
			if deployKeyResp.Payload == nil || deployKeyResp.Payload.PublicDeployKey == nil || len(*deployKeyResp.Payload.PublicDeployKey) == 0 {
				time.Sleep(2 * time.Second) // Sleep before trying again
				continue
			}
			print(strings.TrimRight(*deployKeyResp.Payload.PublicDeployKey, "\t \n"))
			return nil
		}
	},
}

func init() {
	createCmd.AddCommand(createApplicationCmd)
	createApplicationCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to create")
	createApplicationCmd.Flags().StringP(flagnames.Repository, "", "", "The GitHub repository URL")
	createApplicationCmd.Flags().StringP(flagnames.SharedSecret, "", "", "(Optional) Shared secret for the webhook. It is needed when use a GitHub webhook.")
	createApplicationCmd.Flags().StringP(flagnames.ConfigBranch, "", "", "Name of the branch where Radix will read your radixconfig.yaml from")
	createApplicationCmd.Flags().StringP(flagnames.ConfigFile, "", "", "(Optional) Name of the radix config file. Defaults to radixconfig.yaml")
	createApplicationCmd.Flags().StringSliceP(flagnames.AdminADGroups, "", []string{}, "Admin groups (UUIDs of Microsoft Entra groups). Optional for the Playground context")
	createApplicationCmd.Flags().StringSliceP(flagnames.ReaderADGroups, "", []string{}, "(Optional) Reader groups (UUIDs of Microsoft Entra groups)")
	createApplicationCmd.Flags().StringP(flagnames.ConfigurationItem, "", "", "Configuration item (UUID from Business Applications Inventory). Optional for the Playground context")
	createApplicationCmd.Flags().Bool(flagnames.AcknowledgeWarnings, false, "(Optional) Acknowledge warnings and proceed when multiple application use the same GitHub repository")
	setContextSpecificPersistentFlags(createApplicationCmd)
}
