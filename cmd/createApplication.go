// Copyright © 2022
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

	"github.com/equinor/radix-cli/generated-client/client/platform"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

const createApplicationEnabled = true

// createApplicationCmd represents the create application command
var createApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Create application",
	Long:  `Creates a Radix application in the cluster`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		repository, _ := cmd.Flags().GetString("repository")
		sharedSecret, _ := cmd.Flags().GetString("shared-secret")
		configBranch, _ := cmd.Flags().GetString("config-branch")
		configFile, _ := cmd.Flags().GetString("config-file")
		configurationItem, _ := cmd.Flags().GetString("configuration-item")
		acknowledgeWarnings, err := cmd.Flags().GetBool("acknowledge-warnings")
		if err != nil {
			println(fmt.Sprintf("invalid argument %s: %v", "acknowledge-warnings", err))
			return err
		}

		if appName == nil || *appName == "" || repository == "" || configBranch == "" || configurationItem == "" {
			return errors.New("application name, repository, configuration item and config branch are required fields")
		}

		adGroups, _ := cmd.Flags().GetStringSlice("ad-groups")

		registerApplicationParams := platform.NewRegisterApplicationParams()
		registerApplicationParams.SetApplicationRegistration(&models.ApplicationRegistrationRequest{
			AcknowledgeWarnings: acknowledgeWarnings,
			ApplicationRegistration: &models.ApplicationRegistration{
				Name:                appName,
				Repository:          &repository,
				SharedSecret:        &sharedSecret,
				AdGroups:            adGroups,
				ConfigBranch:        &configBranch,
				RadixConfigFullName: configFile,
				ConfigurationItem:   configurationItem,
			},
		})

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		resp, err := apiClient.Platform.RegisterApplication(registerApplicationParams, nil)

		if err != nil {
			println(fmt.Sprintf("%v", err))
			return err
		}
		registrationUpsertResponse := resp.Payload
		if len(registrationUpsertResponse.Warnings) == 0 {
			if registrationUpsertResponse.ApplicationRegistration != nil {
				print(strings.TrimRight(registrationUpsertResponse.ApplicationRegistration.PublicKey, "\t \n"))
				return nil
			}
			return errors.New("unspecified error")
		}
		println("Warnings:")
		for _, warning := range registrationUpsertResponse.Warnings {
			println(fmt.Sprintf("- %s", warning))
		}
		println("If you agree to proceed with warnings - please add an option --acknowledge-warnings")
		return nil
	},
}

func init() {
	if createApplicationEnabled {
		createCmd.AddCommand(createApplicationCmd)
		createApplicationCmd.Flags().StringP("application", "a", "", "Name of the application to create")
		createApplicationCmd.Flags().StringP("repository", "", "", "Repository path")
		createApplicationCmd.Flags().StringP("shared-secret", "", "", "Shared secret for the webhook")
		createApplicationCmd.Flags().StringP("config-branch", "", "", "Name of the branch where Radix will read your radixconfig.yaml from")
		createApplicationCmd.Flags().StringP("config-file", "", "", "Name of the radix config file. Optional, defaults to radixconfig.yaml")
		createApplicationCmd.Flags().StringSliceP("ad-groups", "", []string{}, "Admin groups")
		createApplicationCmd.Flags().StringP("configuration-item", "", "", "Configuration item")
		createApplicationCmd.Flags().Bool("acknowledge-warnings", false, "Acknowledge warnings and proceed")
	}
}
