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
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/spf13/cobra"
	"log"
)

var getContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Gets current context",
	Long: fmt.Sprintf("Gets the current context. It can be one of %s, %s, %s or %s",
		radixconfig.ContextPlatform, radixconfig.ContextPlatform2, radixconfig.ContextPlayground, radixconfig.ContextDevelopment),
	RunE: func(cmd *cobra.Command, args []string) error {

		cmd.SilenceUsage = true

		radixConfig, err := radixconfig.GetRadixConfig()
		if err != nil {
			return err
		}
		log.Printf("Current context is '%s'", radixConfig.CustomConfig.Context)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getContextCmd)
}
