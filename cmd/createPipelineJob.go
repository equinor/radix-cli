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
	"github.com/equinor/radix-cli/pkg/flagnames"

	"github.com/spf13/cobra"
)

// createJobCmd represents the triggering of pipeline command
var createJobCmd = &cobra.Command{
	Use:     "pipeline-job",
	Aliases: []string{"job"},
	Short:   "Create pipeline job command",
	Long:    `Will be the main command for triggering pipelines.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("please specify the pipeline you want to create")
	},
}

func init() {
	createCmd.AddCommand(createJobCmd)
	setContextSpecificPersistentFlags(createJobCmd)
}

func getGitRefAndType(cmd *cobra.Command) (string, string, error) {
	var errs []error
	branch, err := cmd.Flags().GetString(flagnames.Branch)
	if err != nil {
		errs = append(errs, err)
	}
	tag, err := cmd.Flags().GetString(flagnames.Tag)
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return "", "", errors.Join(errs...)
	}
	if len(tag) > 0 {
		return tag, "tag", nil
	}
	return branch, "branch", nil
}
