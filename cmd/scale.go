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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scale component replicas",
	Long:  `Scale component replicas.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Warn("This command is deprecated. Please use 'rx scale component' instead. Will be removed after September 2025")
		return scaleComponentCmd.RunE(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
	scaleCmd.PersistentFlags().StringP(flagnames.Application, "a", "", "Name of the application namespace")
	scaleCmd.PersistentFlags().StringP(flagnames.Environment, "e", "", "Name of the environment of the application")
	scaleCmd.PersistentFlags().StringP(flagnames.Component, "n", "", "Name of the component to scale")
	scaleCmd.PersistentFlags().IntP(flagnames.Replicas, "r", 1, "The new desired number of replicas")
	scaleCmd.PersistentFlags().Bool(flagnames.Reset, false, "Reset manual scaling to resume normal operations")
	scaleCmd.MarkFlagsOneRequired(flagnames.Replicas, flagnames.Reset)
	scaleCmd.MarkFlagsMutuallyExclusive(flagnames.Replicas, flagnames.Reset)
	setContextSpecificPersistentFlags(scaleComponentCmd)
}
