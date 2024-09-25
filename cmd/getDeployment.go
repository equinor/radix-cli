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

	apiclient "github.com/equinor/radix-cli/generated-client/client"
	"github.com/equinor/radix-cli/generated-client/client/application"
	"github.com/equinor/radix-cli/generated-client/client/environment"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/json"

	"github.com/equinor/radix-cli/generated-client/client/deployment"
	"github.com/equinor/radix-cli/pkg/client"
	"github.com/spf13/cobra"
)

// getDeploymentCmd represents the getDeploymentCmd command
var getDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Gets deployments for a given application",
	Long: `Gets deployments for a given application and deployment name or environment

Examples:
  # Get a deployments an application radix-test 
  rx get deployment --application radix-test

  # Get a deployment deployment-abc for an application radix-test 
  rx get deployment --application radix-test --deployment deployment-abc

  # Get a deployments for an application radix-test and its environment test
  rx get deployment --application radix-test --environment test
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}
		if appName == nil || *appName == "" {
			return errors.New("application name is required field")
		}

		deploymentName, err := cmd.Flags().GetString(flagnames.Deployment)
		if err != nil {
			return err
		}
		envName, err := cmd.Flags().GetString(flagnames.Environment)
		if err != nil {
			return err
		}
		if deploymentName != "" && envName != "" {
			return errors.New("options 'deployment' and 'environment' cannot be used together")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetForCommand(cmd)
		if err != nil {
			return err
		}

		if deploymentName == "" && envName == "" {
			return getDeploymentForAllEnvironments(apiClient, *appName)
		}
		if deploymentName != "" {
			return getDeployment(apiClient, *appName, deploymentName)
		}
		return getDeploymentForEnvironment(apiClient, *appName, envName)
	},
}

func getDeployment(apiClient *apiclient.Radixapi, appName, deploymentName string) error {
	params := deployment.NewGetDeploymentParams()
	params.WithAppName(appName)
	if deploymentName != "" {
		params.WithDeploymentName(deploymentName)
	}
	resp, err := apiClient.Deployment.GetDeployment(params, nil)
	if err != nil {
		return err
	}
	prettyJSON, err := json.Pretty(resp.Payload)
	if err != nil {
		return err
	}
	fmt.Println(*prettyJSON)
	return nil
}

func getDeploymentForAllEnvironments(apiClient *apiclient.Radixapi, appName string) error {
	params := application.NewGetDeploymentsParams()
	params.WithAppName(appName)
	resp, err := apiClient.Application.GetDeployments(params, nil)
	if err != nil {
		return err
	}
	prettyJSON, err := json.Pretty(resp.Payload)
	if err != nil {
		return err
	}
	fmt.Println(*prettyJSON)
	return nil
}

func getDeploymentForEnvironment(apiClient *apiclient.Radixapi, appName, envName string) error {
	params := environment.NewGetApplicationEnvironmentDeploymentsParams()
	params.WithAppName(appName)
	params.WithEnvName(envName)
	resp, err := apiClient.Environment.GetApplicationEnvironmentDeployments(params, nil)
	if err != nil {
		return err
	}
	prettyJSON, err := json.Pretty(resp.Payload)
	if err != nil {
		return err
	}
	fmt.Println(*prettyJSON)
	return nil
}

func init() {
	getCmd.AddCommand(getDeploymentCmd)
	getDeploymentCmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application")
	getDeploymentCmd.Flags().StringP(flagnames.Deployment, "d", "", "Optional, name of a deployment. It cannot be used together with an option 'environment'.")
	getDeploymentCmd.Flags().StringP(flagnames.Environment, "e", "", "Optional, name of the environment. It cannot be used together with an option 'deployment'.")

	_ = getDeploymentCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = getDeploymentCmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	setContextSpecificPersistentFlags(getDeploymentCmd)
}
