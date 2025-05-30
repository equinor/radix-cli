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
	"github.com/equinor/radix-cli/generated/radixapi/client/platform"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/json"
	"github.com/spf13/cobra"
)

// getApplicationCmd represents the getApplicationCmd command
var getApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Gets Radix application",
	Long:  `Gets a list of Radix applications or a single application if provided`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		if appName == "" {
			// List applications
			showApplicationParams := platform.NewShowApplicationsParams()
			resp, err := apiClient.Platform.ShowApplications(showApplicationParams, nil)

			var appNames []string

			if err == nil {
				for _, application := range resp.Payload {
					fmt.Println(*application.Name)
					appNames = append(appNames, *application.Name)
				}
				completion.UpdateAppNamesCache(appNames)

				return nil
			}

			return err
		}
		getApplicationParams := application.NewGetApplicationParams()
		getApplicationParams.SetAppName(appName)
		resp, err := apiClient.Application.GetApplication(getApplicationParams, nil)
		if err != nil {
			return err
		}
		prettyJSON, err := json.Pretty(resp.Payload)
		if err != nil {
			return err
		}
		fmt.Println(*prettyJSON)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getApplicationCmd)
	getApplicationCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	_ = getApplicationCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(getApplicationCmd)
}
