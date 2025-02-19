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

	"github.com/equinor/radix-cli/generated/radixapi/client/application"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// deleteApplicationCmd represents the create application command
var deleteApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Delete application",
	Long:  `Will delete an application from the cluster`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == "" {
			return errors.New("application name is a required field")
		}

		cmd.SilenceUsage = true

		deleteApplicationParams := application.NewDeleteApplicationParams()
		deleteApplicationParams.SetAppName(appName)

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Application.DeleteApplication(deleteApplicationParams, nil)
		return err
	},
}

func init() {
	deleteCmd.AddCommand(deleteApplicationCmd)
	deleteApplicationCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application to delete")
	_ = deleteApplicationCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	setContextSpecificPersistentFlags(deleteApplicationCmd)
}
