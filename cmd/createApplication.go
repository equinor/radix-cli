// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
		owner, _ := cmd.Flags().GetString("owner")
		sharedSecret, _ := cmd.Flags().GetString("shared-secret")

		if appName == nil || *appName == "" || repository == "" || owner == "" {
			return errors.New("Application name, repository and owner are required fields")
		}

		adGroups, _ := cmd.Flags().GetStringSlice("ad-groups")

		registerApplicationParams := platform.NewRegisterApplicationParams()
		registerApplicationParams.SetApplicationRegistration(&models.ApplicationRegistration{
			Name:         appName,
			Repository:   &repository,
			Owner:        &owner,
			SharedSecret: &sharedSecret,
			AdGroups:     adGroups,
		})

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		resp, err := apiClient.Platform.RegisterApplication(registerApplicationParams, nil)

		if err == nil {
			print(strings.TrimRight(resp.Payload.PublicKey, "\t \n"))
		} else {
			println(fmt.Sprintf("%v", err))
		}

		return nil
	},
}

func init() {
	if createApplicationEnabled {
		createCmd.AddCommand(createApplicationCmd)
		createApplicationCmd.Flags().StringP("application", "a", "", "Name of the application to create")
		createApplicationCmd.Flags().StringP("repository", "", "", "Repository path")
		createApplicationCmd.Flags().StringP("owner", "", "", "Email adress of owner")
		createApplicationCmd.Flags().StringP("shared-secret", "", "", "Shared secret for the webhook")
		createApplicationCmd.Flags().StringSliceP("ad-groups", "", []string{}, "Admin groups")
	}
}
