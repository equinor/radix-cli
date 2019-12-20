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
	"log"

	"github.com/equinor/radix-cli/generated-client/client/platform"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// listApplicationsCmd represents the listApplications command
var listApplicationsCmd = &cobra.Command{
	Use:   "applications",
	Short: "Lists applications",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		showApplicationParams := platform.NewShowApplicationsParams()
		resp, err := apiClient.Platform.ShowApplications(showApplicationParams, nil)

		if err == nil {
			for _, application := range resp.Payload {
				log.Printf("App: %s", application.Name)
			}
		}

		return nil
	},
}
