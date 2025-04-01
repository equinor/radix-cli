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
	"fmt"

	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// setBuildSecretCmd represents the setBuildSecretCmd command
var setPrivateImageHubSecretCmd = &cobra.Command{
	Use:   "private-image-hub-secret",
	Short: "Will set a Private Image Hub secret",
	Long:  `Will set a Private Image Hub secret`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		domain, _ := cmd.Flags().GetString(flagnames.Server)
		secretValue, _ := cmd.Flags().GetString(flagnames.Value)

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		secretParams := &models.SecretParameters{}
		secretParams.SecretValue = &secretValue

		params := application.NewUpdatePrivateImageHubsSecretValueParams()
		params.SetAppName(appName)
		params.SetServerName(domain)
		params.SetImageHubSecret(secretParams)

		if _, err = apiClient.Application.UpdatePrivateImageHubsSecretValue(params, nil); err != nil {
			return err
		}
		fmt.Printf("Successfully set private image hub secret for domain %s in app %s\n", domain, appName)
		return nil
	},
}

func init() {
	setCmd.AddCommand(setPrivateImageHubSecretCmd)
	setPrivateImageHubSecretCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to set secret for")
	setPrivateImageHubSecretCmd.Flags().StringP(flagnames.Server, "s", "", "Server of private image hub")
	setPrivateImageHubSecretCmd.Flags().StringP(flagnames.Value, "v", "", "Value of the secret to set")

	_ = setPrivateImageHubSecretCmd.MarkFlagRequired(flagnames.Application)
	_ = setPrivateImageHubSecretCmd.MarkFlagRequired(flagnames.Server)
	_ = setPrivateImageHubSecretCmd.MarkFlagRequired(flagnames.Value)

	_ = setPrivateImageHubSecretCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = setPrivateImageHubSecretCmd.RegisterFlagCompletionFunc(flagnames.Server, completion.PrivateImageHubServerCompletion)
	setContextSpecificPersistentFlags(setPrivateImageHubSecretCmd)
}
