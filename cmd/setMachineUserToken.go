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

const setMachineUserTokenEnabled = true

var setMachineUserToken = &cobra.Command{
	Use:   "machine-user-token",
	Short: "Generate machine user token",
	Long:  `Will generate machine user token and return it.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, "application")
		if err != nil {
			return err
		}

		if appName == nil || *appName == "" {
			return errors.New("Application name is a required field")
		}

		regenerateMachineUserTokenParams := application.NewRegenerateMachineUserTokenParams()
		regenerateMachineUserTokenParams.SetAppName(*appName)

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		token, err := apiClient.Application.RegenerateMachineUserToken(regenerateMachineUserTokenParams, nil)

		if err != nil {
			println(fmt.Sprintf("%v", err))
			return err
		}
		if token.Payload != nil && token.Payload.Token != nil {
			fmt.Print(*token.Payload.Token)
		}
		return nil
	},
}

func init() {
	if setMachineUserTokenEnabled {
		setCmd.AddCommand(setMachineUserToken)
		setMachineUserToken.Flags().StringP(applicationOption, "a", "", "Name of the application to generate machine user token for")
	}
}
