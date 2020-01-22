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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/models"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// getApplicationCmd represents the getApplicationCmd command
var getApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Will get the radix application",
	Long:  `Will get the radix application`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil {
			return errors.New("Application name is a required field")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		getApplicationParams := application.NewGetApplicationParams()
		getApplicationParams.SetAppName(*appName)
		resp, err := apiClient.Application.GetApplication(getApplicationParams, nil)
		if err == nil {
			json, err := getPrettyJSON(resp.Payload)
			if err == nil {
				fmt.Println(*json)
			} else {
				println(fmt.Sprintf("%v", err))
			}

		} else {
			println(fmt.Sprintf("%v", err))
		}

		return nil
	},
}

func getPrettyJSON(data *models.Application) (*string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, b, "", "\t")
	if err != nil {
		return nil, err
	}

	jsonString := string(prettyJSON.Bytes())
	return &jsonString, nil
}

func init() {
	getApplicationCmd.Flags().StringP("application", "a", "", "Name of the application")
}
