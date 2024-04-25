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
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Radix",
	Long:  `Login to Radix.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		useDeviceCode, _ := cmd.Flags().GetBool(flagnames.UseDeviceCode)
		err := client.LoginCommand(cmd, useDeviceCode)
		if err != nil {
			return err
		}
		println("Logged in to Radix")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().Bool(flagnames.UseDeviceCode, false, "Use CLI's old authentication flow based on device code. The device code flow does not work for compliant device policy enabled accounts.")
	setVerbosePersistentFlag(loginCmd)
}
