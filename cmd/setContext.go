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
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/spf13/cobra"
)

const setContextEnabled = true

// setContextCmd represents the setContext command
var setContextCmd = &cobra.Command{
	Use: "context",
	Short: fmt.Sprintf("Sets the context to be either %s, %s, %s or %s",
		radixconfig.ContextPlatform, radixconfig.ContextPlatform2, radixconfig.ContextPlayground, radixconfig.ContextDevelopment),
	Long: fmt.Sprintf("Sets the context to be either %s, %s, %s or %s",
		radixconfig.ContextPlatform, radixconfig.ContextPlatform2, radixconfig.ContextPlayground, radixconfig.ContextDevelopment),
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString(settings.ContextOption)

		if !radixconfig.IsValidContext(name) {
			return fmt.Errorf("context '%s' is not a valid context", name)
		}

		radixConfig := radixconfig.RadixConfigAccess{}
		config := radixConfig.GetStartingConfig().Config
		config[settings.ContextOption] = name
		persister := radixconfig.PersisterForRadix(radixConfig)
		persister.Persist(config)
		return nil
	},
}

func init() {
	if setContextEnabled {
		setCmd.AddCommand(setContextCmd)
	}
}
