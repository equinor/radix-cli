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
	"errors"

	"github.com/spf13/cobra"
)

const triggernabled = true

// triggerCmd represents the triggering of pipeline command
var triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Trigger pipeline",
	Long:  `Will be the main command for triggering pipelines.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Please specify the resource you want to follow")
	},
}

func init() {
	if triggernabled {
		rootCmd.AddCommand(triggerCmd)
	}
}
