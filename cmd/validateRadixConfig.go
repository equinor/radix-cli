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
	"fmt"
	"os"

	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	radixv1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	"github.com/equinor/radix-operator/pkg/apis/radixvalidators"
	"github.com/equinor/radix-operator/pkg/apis/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yannh/kubeconform/pkg/validator"
	"sigs.k8s.io/yaml"
)

// logoutCmd represents the logout command
var validateRadixConfigCmd = &cobra.Command{
	Use:   "radix-config",
	Short: "Validate radixconfig.yaml",
	Long:  `Check radixconfig.yaml for structural and logical errors`,
	RunE: func(cmd *cobra.Command, args []string) error {

		cmd.SilenceUsage = true

		radixconfig, err := cmd.Flags().GetString(flagnames.ConfigFile)
		if err != nil {
			return err
		}

		printfile, err := cmd.Flags().GetBool(flagnames.Print)
		if err != nil {
			return err
		}

		schema, err := cmd.Flags().GetString(flagnames.Schema)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stderr, "Validating %s\n", radixconfig)
		if _, err := os.Stat(radixconfig); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("RadixConfig file not found")
		}

		ra, err := utils.GetRadixApplicationFromFile(radixconfig)
		if err != nil {
			return err
		}

		if printfile {
			err = printRA(ra)
			if err != nil {
				return err
			}
		}

		errs, err := validateSchema(radixconfig, schema)
		if err != nil {
			return err
		}

		err = radixvalidators.IsRadixApplicationValid(ra)
		if err != nil {
			errs = append(errs, err)
		}

		err = strictUnmarshalValidation(radixconfig)
		if err != nil {
			errs = append(errs, err)
		}

		if len(errs) > 0 {
			for _, err := range errs {
				fmt.Fprintf(os.Stderr, " - %s\n", err.Error())
			}

			fmt.Fprintln(os.Stderr, "RadixConfig is invalid ")
			os.Exit(2)
		}

		fmt.Fprintln(os.Stderr, "RadixConfig is valid")
		return nil
	},
}

func validateSchema(filename, schema string) ([]error, error) {
	var errs []error
	v, err := validator.New([]string{schema}, validator.Opts{Strict: true})
	if err != nil {
		return nil, fmt.Errorf("failed initializing validator: %s", err)
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening %s: %s", filename, err)
	}
	for _, res := range v.Validate(filename, f) { // A file might contain multiple resources
		for _, err := range res.ValidationErrors {
			errs = append(errs, fmt.Errorf("%s: %s", err.Path, err.Msg))
		}
	}
	return errs, nil
}

func printRA(ra *radixv1.RadixApplication) error {
	b, err := yaml.Marshal(ra)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%s", b)
	return nil
}

func init() {
	validateCmd.AddCommand(validateRadixConfigCmd)
	validateRadixConfigCmd.Flags().StringP(flagnames.ConfigFile, "f", "radixconfig.yaml", "Name of the radixconfig file. Defaults to radixconfig.yaml in current directory")
	validateRadixConfigCmd.Flags().BoolP(flagnames.Print, "p", false, "Print parsed config file")
	validateRadixConfigCmd.Flags().StringP(flagnames.Schema, "s", "https://raw.githubusercontent.com/equinor/radix-operator/release/json-schema/radixapplication.json", "Validate against schema")

	// Allow but hide token-env flag so radix-github-actions won't interfere
	validateRadixConfigCmd.Flags().Bool(flagnames.TokenEnvironment, false, fmt.Sprintf("Take the token from environment variable %s", client.TokenEnvironmentName))
	err := validateRadixConfigCmd.Flags().MarkHidden(flagnames.TokenEnvironment)
	if err != nil {
		panic(err)
	}
}

func strictUnmarshalValidation(filename string) error {
	log.Debug("get radix application yaml from %s", filename)
	radixApp := &radixv1.RadixApplication{}

	raw, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	err = yaml.UnmarshalStrict(raw, radixApp)
	if err != nil {
		return fmt.Errorf("strict test failed: %v", err)
	}

	return nil
}
