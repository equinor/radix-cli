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
	"context"
	"fmt"

	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/spf13/cobra"
)

var (
	errInvalidAzureClientFlags   = fmt.Errorf("%s must be used together with %s, %s or %s", flagnames.AzureClientId, flagnames.AzureClientSecret, flagnames.FederatedTokenFile, flagnames.UseGithubCredentials)
	errMissingAzureClientIdFlags = fmt.Errorf("%s, %s or %s must be used together with %s", flagnames.AzureClientSecret, flagnames.FederatedTokenFile, flagnames.UseGithubCredentials, flagnames.AzureClientId)
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Radix",
	Long:  `Login to Radix.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		useGithubCredentials, _ := cmd.Flags().GetBool(flagnames.UseGithubCredentials)
		useInteractiveLogin, _ := cmd.Flags().GetBool(flagnames.UseInteractiveLogin)
		useDeviceCode, _ := cmd.Flags().GetBool(flagnames.UseDeviceCode)
		federatedTokenFile, _ := cmd.Flags().GetString(flagnames.FederatedTokenFile)
		azureClientId, _ := cmd.Flags().GetString(flagnames.AzureClientId)
		azureClientSecret, _ := cmd.Flags().GetString(flagnames.AzureClientSecret)

		if azureClientId != "" && !useGithubCredentials && azureClientSecret == "" && federatedTokenFile == "" {
			return errInvalidAzureClientFlags
		}
		if (useGithubCredentials || federatedTokenFile != "" || azureClientSecret != "") && azureClientId == "" {
			return errMissingAzureClientIdFlags
		}
		if !useInteractiveLogin && !useDeviceCode && !useGithubCredentials && azureClientId == "" {
			useInteractiveLogin = true
		}

		err := client.LoginCommand(context.Background(), useInteractiveLogin, useDeviceCode, useGithubCredentials, azureClientId, federatedTokenFile, azureClientSecret)
		if err != nil {
			return err
		}
		println("Logged in to Radix")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().Bool(flagnames.UseInteractiveLogin, false, "Authenticate with Azure Interactive Login. Default if no other option is specified")
	loginCmd.Flags().Bool(flagnames.UseDeviceCode, false, "Authenticate with Azure Device Code")
	loginCmd.Flags().Bool(flagnames.UseGithubCredentials, false, "Authenticate with Github Workload Identity")
	loginCmd.Flags().String(flagnames.AzureClientId, "", "Authenticate with Azure Client Id and federated token or client secret")
	loginCmd.Flags().String(flagnames.FederatedTokenFile, "", "Authenticate with Federated Credentials and Azure Client Id")
	loginCmd.Flags().String(flagnames.AzureClientSecret, "", "Authenticate with Azure Client Secret and Azure Client Id")

	loginCmd.MarkFlagsMutuallyExclusive(flagnames.UseInteractiveLogin, flagnames.AzureClientSecret, flagnames.UseGithubCredentials, flagnames.FederatedTokenFile, flagnames.UseDeviceCode)
	setVerbosePersistentFlag(loginCmd)
}
