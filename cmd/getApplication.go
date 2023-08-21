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
	"strings"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/client/platform"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/utils/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getApplicationCmd represents the getApplicationCmd command
var getApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Gets Radix application",
	Long:  `Gets a list of Radix applications or a single application if provided`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		if appName == nil || strings.EqualFold(*appName, "") {
			// List applications
			showApplicationParams := platform.NewShowApplicationsParams()
			resp, err := apiClient.Platform.ShowApplications(showApplicationParams, nil)

			if err == nil {
				for _, application := range resp.Payload {
					log.Infof("App: %s", application.Name)
				}
			}
			return err
		}
		getApplicationParams := application.NewGetApplicationParams()
		getApplicationParams.SetAppName(*appName)
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
	getApplicationCmd.Flags().StringP("application", "a", "", "Name of the application")
	setContextSpecificPersistentFlags(getApplicationCmd)
}
