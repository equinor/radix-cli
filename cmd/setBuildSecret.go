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

	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/generated/radixapi/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// setBuildSecretCmd represents the setBuildSecretCmd command
var setBuildSecretCmd = &cobra.Command{
	Use:   "build-secret",
	Short: "Will set a build secret",
	Long:  `Will set a build secret`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == "" {
			return errors.New("application name is required")
		}

		secretName, _ := cmd.Flags().GetString(flagnames.Secret)
		secretValue, _ := cmd.Flags().GetString(flagnames.Value)

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		componentSecret := models.SecretParameters{}
		componentSecret.SecretValue = &secretValue

		params := application.NewUpdateBuildSecretsSecretValueParams()
		params.SetAppName(appName)
		params.SetSecretName(secretName)
		params.SetSecretValue(&componentSecret)

		if _, err = apiClient.Application.UpdateBuildSecretsSecretValue(params, nil); err != nil {
			return err
		}
		fmt.Printf("Successfully set build secret %s for app %s\n", secretName, appName)
		return nil
	},
}

func init() {
	setCmd.AddCommand(setBuildSecretCmd)
	setBuildSecretCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to set secret for")
	setBuildSecretCmd.Flags().StringP(flagnames.Secret, "s", "", "Name of the secret to set")
	setBuildSecretCmd.Flags().StringP(flagnames.Value, "v", "", "Value of the secret to set")

	_ = setBuildSecretCmd.MarkFlagRequired(flagnames.Application)
	_ = setBuildSecretCmd.MarkFlagRequired(flagnames.Secret)
	_ = setBuildSecretCmd.MarkFlagRequired(flagnames.Value)

	_ = setBuildSecretCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = setBuildSecretCmd.RegisterFlagCompletionFunc(flagnames.Secret, completion.BuildSecretCompletion)
	setContextSpecificPersistentFlags(setBuildSecretCmd)
}
