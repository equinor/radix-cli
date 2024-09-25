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
	"github.com/equinor/radix-cli/pkg/flagnames"
	"github.com/equinor/radix-cli/pkg/utils/completion"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var scaleCmd = &cobra.Command{
	Use:        "scale",
	Short:      "Scale component replicas",
	Long:       `Scale component replicas.`,
	Deprecated: "Please use 'rx scale component' instead. Will be removed after September 2025",
	RunE: func(cmd *cobra.Command, args []string) error {
		return scaleComponentCmd.RunE(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
	scaleCmd.PersistentFlags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	scaleCmd.PersistentFlags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")
	scaleCmd.PersistentFlags().StringP(flagnames.Component, "n", "", "Name of the component to scale")
	scaleCmd.PersistentFlags().IntP(flagnames.Replicas, "r", 1, "The new desired number of replicas")
	scaleCmd.PersistentFlags().Bool(flagnames.Reset, false, "Reset manualy scaled component to use replica count from RadixConfig or managed by horizontal autoscaling")
	scaleCmd.MarkFlagsOneRequired(flagnames.Replicas, flagnames.Reset)
	scaleCmd.MarkFlagsMutuallyExclusive(flagnames.Replicas, flagnames.Reset)

	_ = scaleCmd.RegisterFlagCompletionFunc(flagnames.Application, completion.ApplicationCompletion)
	_ = scaleCmd.RegisterFlagCompletionFunc(flagnames.Component, completion.ComponentCompletion)
	setContextSpecificPersistentFlags(scaleComponentCmd)
}
