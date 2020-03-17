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
	"fmt"
	"log"

	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/spf13/cobra"
)

const getContextEnabled = true

// getContextCmd represents the getContext command
var getContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Gets current context",
	Long: fmt.Sprintf("Gets the current context. It can be one of %s, %s or %s",
		radixconfig.ContextProdction, radixconfig.ContextPlayground, radixconfig.ContextDevelopment),
	Run: func(cmd *cobra.Command, args []string) {
		radixConfig := radixconfig.RadixConfigAccess{}
		config := radixConfig.GetStartingConfig().Config
		log.Printf("Current context is %s", config[settings.ContextOption])
	},
}

func init() {
	if getContextEnabled {
		getCmd.AddCommand(getContextCmd)
	}
}
