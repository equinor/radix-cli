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
	"fmt"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// deleteApplicationCmd represents the create application command
var deleteApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Delete application",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("Application name is a required field")
		}

		deleteApplicationParams := application.NewDeleteApplicationParams()
		deleteApplicationParams.SetAppName(*appName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		_, err = apiClient.Application.DeleteApplication(deleteApplicationParams, nil)

		if err != nil {
			println(fmt.Sprintf("%v", err))
		}

		return nil
	},
}

func init() {
	deleteApplicationCmd.Flags().StringP("application", "", "", "Name of the application to create")
	deleteApplicationCmd.Flags().StringP("repository", "", "", "Repository")
	deleteApplicationCmd.Flags().StringP("owner", "", "", "Owner")
	deleteApplicationCmd.Flags().StringSliceP("ad-groups", "", []string{}, "Admin groups")
}
