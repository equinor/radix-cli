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
	"github.com/equinor/radix-cli/generated/radixapi/client/configuration"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// getClusterConfigCmd represents the get-cluster-config command
var getClusterConfigCmd = &cobra.Command{
	Use:   "cluster-config",
	Short: "Gets setting from Radix cluster config",
	Long:  `Helper functionality to get data from radix cluster config.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		payload, err := apiClient.Configuration.GetConfiguration(configuration.NewGetConfigurationParams(), nil)
		if err != nil {
			return err
		}

		printPayload(payload.Payload)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getClusterConfigCmd)
}
