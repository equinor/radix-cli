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

	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/config"
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/settings"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/equinor/radix-cli/pkg/utils/replicalog"
	radixv1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	"github.com/spf13/cobra"
)

// logsEnvironmentComponentOAuth2Cmd represents the logs for oauth2 command
var logsEnvironmentComponentOAuth2Cmd = &cobra.Command{
	Use:   "oauth2",
	Short: "Gets logs for the OAuth2 auxiliary resource of a component",
	Long: `Gets and follows logs for the OAuth2 auxiliary resource of a component in an environment.

It may take few seconds to get the log.

Examples:
  # Get logs for the OAuth2 auxiliary resource of a component
  rx get logs component oauth2 --application radix-test --environment dev --component web-app

  # Short version
  rx get logs component oauth2 -a radix-test -e dev --component web-app
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, err := config.GetAppNameFromConfigOrFromParameter(cmd, flagnames.Application)
		if err != nil {
			return err
		}

		if appName == "" {
			return errors.New("application name is required")
		}

		environmentName, _ := cmd.Flags().GetString(flagnames.Environment)
		componentName, _ := cmd.Flags().GetString(flagnames.Component)
		previousLog, _ := cmd.Flags().GetBool(flagnames.Previous)
		since, _ := cmd.Flags().GetDuration(flagnames.Since)

		if environmentName == "" || componentName == "" {
			return errors.New("both `environment` and `component` are required")
		}

		cmd.SilenceUsage = true

		apiClient, err := client.GetRadixApiForCommand(cmd)
		if err != nil {
			return err
		}

		return replicalog.New(
			cmd.ErrOrStderr(),
			replicalog.GetReplicasForComponentOAuth2(apiClient, appName, environmentName, componentName, radixv1.OAuthProxyAuxiliaryComponentType, previousLog),
			replicalog.GetOAuth2ComponentLog(apiClient, appName, previousLog),
			&since,
		).StreamLogs(cmd.Context(), false)
	},
}

func init() {
	logsEnvironmentComponentCmd.AddCommand(logsEnvironmentComponentOAuth2Cmd)
	logsEnvironmentComponentOAuth2Cmd.Flags().StringP(flagnames.Application, "a", "", "Name of the application owning the component")
	logsEnvironmentComponentOAuth2Cmd.Flags().StringP(flagnames.Environment, "e", "", "Environment the component runs in")
	logsEnvironmentComponentOAuth2Cmd.Flags().String(flagnames.Component, "", "The component to follow")
	logsEnvironmentComponentOAuth2Cmd.Flags().BoolP(flagnames.Previous, "p", false, "If set, print the logs for the previous instance of the OAuth2 auxiliary container, if it exists")
	logsEnvironmentComponentOAuth2Cmd.Flags().DurationP(flagnames.Since, "s", settings.DeltaRefreshApplication, "If set, start get logs from the specified time, eg. 5m or 12h")

	_ = logsEnvironmentComponentOAuth2Cmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = logsEnvironmentComponentOAuth2Cmd.RegisterFlagCompletionFunc(flagnames.Environment, completion.EnvironmentCompletion)
	_ = logsEnvironmentComponentOAuth2Cmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	setContextSpecificPersistentFlags(logsEnvironmentComponentOAuth2Cmd)
}
