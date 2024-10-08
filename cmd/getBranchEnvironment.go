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

	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/spf13/cobra"
)

// getBranchEnvironmentCmd represents the getBranchEnvironmentCmd command
var getBranchEnvironmentCmd = &cobra.Command{
	Use:   "branch-environment",
	Short: "Gets the environment for a given branch",
	Long:  `Gets the environment for a given branch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fromConfig, _ := cmd.Flags().GetBool(flagnames.FromConfig)
		if !fromConfig {
			return errors.New("config can only come from radixconfig file in current folder")
		}

		_, err := config.GetRadixApplicationFromFile()
		if err != nil {
			return err
		}

		branch, _ := cmd.Flags().GetString(flagnames.Branch)

		if branch == "" {
			return errors.New("`branch` is required")
		}

		environment, err := config.GetEnvironmentFromConfig(cmd, branch)
		if err != nil {
			return err
		}

		fmt.Print(environment)
		return nil
	},
}

func init() {
	getConfigCmd.AddCommand(getBranchEnvironmentCmd)
	getBranchEnvironmentCmd.Flags().StringP(flagnames.Branch, "b", "", "Branch of the repository. Should be used together with --from-config to get the environment")
	setContextSpecificPersistentFlags(getBranchEnvironmentCmd)
}
