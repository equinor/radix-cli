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
	"strconv"

	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scale component replicas",
	Long: `Used for scaling up or down replicas of a Radix application component.

Examples:

# Scale up component to 2 replicas
rx scale --application radix-test --environment dev --component component-abc --replicas 2

# Short version of scaling up component to 0 replicas
rx scale -a radix-test -e dev -n component-abc -r 2
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := getAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}
		envName, err := cmd.Flags().GetString(flagnames.Environment)
		if err != nil {
			return err
		}
		cmpName, err := cmd.Flags().GetString(flagnames.Component)
		if err != nil {
			return err
		}
		replicas, err := cmd.Flags().GetInt(flagnames.Replicas)
		if err != nil {
			return err
		}
		if appName == nil || *appName == "" || envName == "" || cmpName == "" {
			return errors.New("application name, environment name and component name are required fields")
		}
		if replicas < 0 || replicas > 20 {
			return errors.New("required field replicas must be between 0 and 20")
		}

		cmd.SilenceUsage = true

		parameters := component.NewScaleComponentParams().
			WithAppName(*appName).
			WithEnvName(envName).
			WithComponentName(cmpName).
			WithReplicas(strconv.Itoa(replicas))

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}
		_, err = apiClient.Component.ScaleComponent(parameters, nil)
		return err
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
	scaleCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	scaleCmd.Flags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")
	scaleCmd.Flags().StringP(flagnames.Component, "n", "", "Name of the component to scale")
	scaleCmd.Flags().IntP(flagnames.Replicas, "r", 1, "The new desired number of replicas")
	setContextSpecificPersistentFlags(scaleCmd)
}
