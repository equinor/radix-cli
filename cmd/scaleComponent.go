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

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/component"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "Scale component replicas",
	Long: `Used for manually scaling up or down replicas of a Radix application component.
Note: Manual scaling will persist across deployments, and will disable autoscaling.
`,
	Example: `
# Scale up component to 2 replicas
rx scale component --application radix-test --environment dev --component component-abc --replicas 2

# Short version of scaling up component to 0 replicas
rx scale component -a radix-test -e dev -n component-abc -r 2

# Reset manual scaling to resume normal operations:
rx scale component --application radix-test --environment dev --component component-abc --reset
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
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
		reset, err := cmd.Flags().GetBool(flagnames.Reset)
		if err != nil {
			return err
		}
		if appName == "" || envName == "" || cmpName == "" {
			return errors.New("application name, environment name and component name are required fields")
		}
		if !reset && (replicas < 0 || replicas > 20) {
			return errors.New("required field replicas must be between 0 and 20")
		}

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		cmd.SilenceUsage = true

		if reset {
			return resetScaledComponent(apiClient, appName, envName, cmpName)
		}
		return scaleComponent(apiClient, appName, envName, cmpName, strconv.Itoa(replicas))
	},
}

func scaleComponent(apiClient *apiclient.Radixapi, appName, envName, cmpName, replicas string) error {
	parameters := component.NewScaleComponentParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithComponentName(cmpName).
		WithReplicas(replicas)

	if _, err := apiClient.Component.ScaleComponent(parameters, nil); err != nil {
		return err
	}

	logrus.Infof("%s Successfully scaled to %s replicas", cmpName, replicas)
	return nil
}

func resetScaledComponent(apiClient *apiclient.Radixapi, appName, envName, cmpName string) error {
	parameters := component.NewResetScaledComponentParams().
		WithAppName(appName).
		WithEnvName(envName).
		WithComponentName(cmpName)

	if _, err := apiClient.Component.ResetScaledComponent(parameters, nil); err != nil {
		return err
	}

	logrus.Infof("%s Successfully reset to normal scaling", cmpName)
	return nil
}

func init() {
	scaleCmd.AddCommand(scaleComponentCmd)
}
