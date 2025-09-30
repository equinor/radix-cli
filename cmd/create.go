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
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/flagvalues"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

var outputFormat = "text"

// createCmd represents the list command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Radix resources",
	Long:  `A longer description .`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("please specify the resource you want to create")
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&outputFormat, flagnames.Output, "o", "text", "(Optional) Output format. json or not set (plain text)")
	_ = createCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.Output)

	rootCmd.AddCommand(createCmd)
}

func printPayload(payload any) {
	if outputFormat == flagvalues.OutputFormatJson {
		jsonData, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			log.Fatalf("failed to print payload as json: %v", err)
		}
		fmt.Println(string(jsonData))
		return
	}

	yamlData, err := yaml.Marshal(payload)
	if err != nil {
		log.Fatalf("failed to print payload as yaml: %v", err)
	}
	fmt.Println(string(yamlData))
}
