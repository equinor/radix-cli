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
	"fmt"

	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/json"
	"github.com/equinor/radix-common/utils/pointers"
	"github.com/spf13/cobra"
)

// getResourcesCmd represents the getResourcesCmd command
var getResourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "Gets resources used by the Radix application",
	Long:  `Gets resources used by the Radix application or its environment or a component`,
	Example: `
# Get resources used by the application for the last 30 days
rx get resources -a myapp

# Get resources used by the application in the environment for the last 30 days
rx get resources -a myapp -e dev

# Get resources used by the application in the environment for a component for the last 30 days
rx get resources -a myapp -e dev -n mycomponent

# Get resources used by the application in the environment for a component for the last 5 minutes
rx get resources -a myapp -e dev -n mycomponent --duration 5m

# Get resources used by the application in the environment for a component for the last 12 hours starting with two days ago
rx get resources -a myapp -e dev -n mycomponent --duration 12h --since 2d
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}
		if appName == "" {
			return errors.New("application name is required field")
		}
		envName, err := cmd.Flags().GetString(flagnames.Environment)
		if err != nil {
			return err
		}
		componentName, err := cmd.Flags().GetString(flagnames.Component)
		if err != nil {
			return err
		}
		since, err := cmd.Flags().GetString(flagnames.Since)
		if err != nil {
			return err
		}
		duration, err := cmd.Flags().GetString(flagnames.Duration)
		if err != nil {
			return err
		}

		getResourcesParams := application.NewGetResourcesParams()
		getResourcesParams.SetAppName(appName)
		getResourcesParams.SetEnvironment(&envName)
		getResourcesParams.SetComponent(&componentName)
		if duration != "" {
			getResourcesParams.SetDuration(pointers.Ptr(duration))
		}
		if since != "" {
			getResourcesParams.SetSince(pointers.Ptr(since))
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}
		resp, err := apiClient.Application.GetResources(getResourcesParams, nil)
		if err != nil {
			return err
		}
		prettyJSON, err := json.Pretty(resp.Payload)
		if err != nil {
			return err
		}
		fmt.Println(*prettyJSON)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getResourcesCmd)
	getResourcesCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	getResourcesCmd.Flags().StringP(flagnames.Environment, "e", "", "Optional, name of the environment")
	getResourcesCmd.Flags().StringP(flagnames.Component, "n", "", "Optional, name of the component")
	getResourcesCmd.Flags().String(flagnames.Duration, "", "If set, get resources during the specified period (default is 30 days), eg. 5m or 12h")
	getResourcesCmd.Flags().String(flagnames.Since, "", "If set, get resources starting from the specified time in the past, eg. 5m or 12h")
	setContextSpecificPersistentFlags(getResourcesCmd)
}
