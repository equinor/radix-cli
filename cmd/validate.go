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

	radixv1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	"github.com/equinor/radix-operator/pkg/apis/radixvalidators"
	"github.com/equinor/radix-operator/pkg/apis/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

// logoutCmd represents the logout command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate radixconfig.yaml",
	Long:  `Check radixconfig.yaml for structural and logical errors`,
	RunE: func(cmd *cobra.Command, args []string) error {

		cmd.SilenceUsage = true

		radixconfig, err := cmd.Flags().GetString("config-file")
		if err != nil {
			return err
		}

		printfile, err := cmd.Flags().GetBool("print")
		if err != nil {
			return err
		}

		if _, err := os.Stat(radixconfig); errors.Is(err, os.ErrNotExist) {
			return errors.New(fmt.Sprintf("RadixConfig file note found: %s", radixconfig))
		}

		ra, err := utils.GetRadixApplicationFromFile(radixconfig)
		if err != nil {
			fmt.Printf("Syntax error in RadixConfig file:\n\n")
			fmt.Println(err)
			return errors.New("RadixConfig is invalid")
		}

		if printfile {
			printRA(ra)
		}

		err = radixvalidators.IsRadixApplicationValid(ra)
		if err == nil {
			fmt.Println("RadixConfig is valid")
			return nil
		}

		fmt.Println(err)

		return errors.New("RadixConfig is invalid")
	},
}

func printRA(ra *radixv1.RadixApplication) {
	b, err := yaml.Marshal(ra)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", b)
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("config-file", "f", "radixconfig.yaml", "Name of the radixconfig file. Defaults to radixconfig.yaml in current directory")
	validateCmd.Flags().BoolP("print", "p", false, "Print parsed config file")
	setVerbosePersistentFlag(validateCmd)
}
