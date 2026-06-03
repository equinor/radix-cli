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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/equinor/radix-cli/pkg/auth"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/workloadidentity"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var validateWorkloadIdentityCmd = &cobra.Command{
	Use:     "workload-identity",
	Short:   "Validate radixconfig.yaml",
	Long:    `Valida workload identity configuration`,
	Example: ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		// apiClient, err := client.GetRadixApiForCommand(cmd)
		// if err != nil {
		// 	return err
		// }

		// apps, err := apiClient.Platform.ShowApplications(&platform.ShowApplicationsParams{}, nil)
		// if err != nil {
		// 	return err
		// }

		// fmt.Println(apps.Payload)
		radixAuth, err := auth.New()
		if err != nil {
			return err
		}
		graphHelper, err := workloadidentity.NewAzureServicePrincipalService(&tokenCredentialAdapter{auth: radixAuth})
		if err != nil {
			return fmt.Errorf("Error initializing client: %w", err)
		}

		// id, err := graphHelper.GetManagedIdentity(cmd.Context(), "16ede44b-1f74-40a5-b428-46cca9a5741b", "test-resources", "id-radix-fed-test")
		// if err != nil {
		// 	return err
		// }
		// fmt.Println(*id.Name)
		// return nil

		resolveAndPrint := func(appId string) error {
			sp, err := graphHelper.GetServicePrincipal(context.Background(), appId)
			if err != nil {
				return err
			}

			fmt.Printf("Found %s with name %s\n", sp.Type, sp.DisplayName)
			fmt.Printf("Number of fed creds: %v\n", len(sp.FederatedCredentials))
			return nil
		}

		if err := resolveAndPrint("5e48ca1f-a2bf-4dec-b96d-bbf8ce69f9f6"); err != nil {
			return err
		}

		if err := resolveAndPrint("b96d264b-7053-4465-a4a7-32be5b0fec49"); err != nil {
			return err
		}

		return nil
	},
}

type tokenCredentialAdapter struct {
	auth *auth.Auth
}

func (a *tokenCredentialAdapter) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	if a.auth == nil {
		return azcore.AccessToken{}, fmt.Errorf("auth not set")
	}

	t, err := a.auth.GetAccessToken(ctx, options.Scopes)
	if err != nil {
		return azcore.AccessToken{}, err
	}

	return azcore.AccessToken{Token: t.Token, ExpiresOn: t.ExpiresOn}, nil
}

func init() {
	validateCmd.AddCommand(validateWorkloadIdentityCmd)
	err := validateRadixConfigCmd.Flags().MarkHidden(flagnames.TokenEnvironment)
	if err != nil {
		panic(err)
	}
}
