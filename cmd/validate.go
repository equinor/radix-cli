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
	"os"

	"github.com/equinor/radix-operator/pkg/apis/radixvalidators"
	"github.com/equinor/radix-operator/pkg/apis/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate radixconfig.yaml",
	Long:  `Make sure the radixconfig.yaml file has correct structure, but doe not check for logic errors in the configuration`,
	RunE: func(cmd *cobra.Command, args []string) error {

		cmd.SilenceUsage = true

		radixconfig, err := cmd.Flags().GetString("radixconfig")
		if err != nil {
			return err
		}

		if _, err := os.Stat(radixconfig); errors.Is(err, os.ErrNotExist) {
			return errors.New(fmt.Sprintf("Config file note found: %s", radixconfig))
		}

		ra, err := utils.GetRadixApplicationFromFile(radixconfig)
		if err != nil {
			fmt.Println(err)
			return errors.New("Radix Config is invalid")
		}

		err = radixvalidators.IsRadixApplicationValid(ra)
		if err == nil {
			fmt.Println("Radixconfig is valid")
			return nil
		}

		fmt.Println(err)

		return errors.New("Radix Config is invalid")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("radixconfig", "c", "radixconfig.yaml", "Path to radixconfig.yaml")
	setVerbosePersistentFlag(validateCmd)
}
