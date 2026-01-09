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
	"os"

	"github.com/equinor/radix-cli/pkg/client"
	"github.com/equinor/radix-cli/pkg/flagnames"
	radixv1 "github.com/equinor/radix-operator/pkg/apis/radix/v1"
	"github.com/equinor/radix-operator/webhook/validation/radixapplication"
	"github.com/fatih/color"
	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

// logoutCmd represents the logout command
var validateRadixConfigCmd = &cobra.Command{
	Use:   "radix-config",
	Short: "Validate radixconfig.yaml",
	Long:  `Check radixconfig.yaml for structural and logical errors`,
	Example: `# Validate radixconfig.yaml in current directory:
rx validate radix-config

# Specify path to radixconfig to validate:
rx validate radix-config --config-file /path/to/anyradixconfig.yaml

# Validate radixconfig without strict validation:
rx validate radix-config --strict-validation=false
`,
	Run: func(cmd *cobra.Command, args []string) {
		radixconfig, err := cmd.Flags().GetString(flagnames.ConfigFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		printfile, err := cmd.Flags().GetBool(flagnames.Print)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		schema, err := cmd.Flags().GetString(flagnames.Schema)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		strictValidation, err := cmd.Flags().GetBool(flagnames.StrictValidation)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "Validating %s\n", radixconfig)
		if _, err := os.Stat(radixconfig); errors.Is(err, os.ErrNotExist) {
			fmt.Fprintln(os.Stderr, "RadixConfig file not found")
			os.Exit(1)
		}

		raw, err := os.ReadFile(radixconfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read file: %v\n", err)
			os.Exit(1)
		}

		ra, err := unmarshalRadixApplication(raw)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		if printfile {
			err = printRA(ra)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}

		validationErrors, err := validateSchema(raw, schema)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		validator := radixapplication.CreateOfflineValidator()
		warnings, err := validator.Validate(cmd.Context(), ra)
		if err != nil {
			validationErrors = append(validationErrors, err)
		}

		if strictValidation {
			err = strictUnmarshalValidation(raw)
			if err != nil {
				validationErrors = append(validationErrors, err)
			}
		}

		if len(warnings) > 0 {
			fmt.Fprintln(os.Stderr, color.HiYellowString("Warnings:"))

			for _, warning := range warnings {
				fmt.Fprintf(os.Stderr, " - %s\n", warning)
			}
		}

		if len(validationErrors) > 0 {
			fmt.Fprintln(os.Stderr, color.RedString("Errors:"))

			for _, err := range validationErrors {
				fmt.Fprintf(os.Stderr, " - %s\n", err)
			}
		}

		if len(validationErrors) == 0 {
			fmt.Fprintln(os.Stderr, color.GreenString("RadixConfig is valid"))
			return
		}

		fmt.Fprintln(os.Stderr, color.RedString("RadixConfig is invalid"))
		os.Exit(2)
	},
}

func validateSchema(raw []byte, schema string) (validationErrors []error, err error) {
	s, err := jsonschema.Compile(schema)
	if err != nil {
		return nil, fmt.Errorf("failed compiling schema %s: %s", schema, err)
	}

	var obj interface{}
	err = yaml.Unmarshal(raw, &obj)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}

	err = s.Validate(obj)
	var verr *jsonschema.ValidationError
	if errors.As(err, &verr) {
		for _, err := range verr.Causes {
			validationErrors = append(validationErrors, fmt.Errorf("%s: %s", err.InstanceLocation, err.Message))
		}
	} else {
		return nil, err
	}

	return validationErrors, nil
}

func strictUnmarshalValidation(raw []byte) error {
	radixApp := &radixv1.RadixApplication{}

	err := yaml.UnmarshalStrict(raw, radixApp)
	if err != nil {
		return fmt.Errorf("strict test failed: %v", err)
	}

	return nil
}
func unmarshalRadixApplication(raw []byte) (*radixv1.RadixApplication, error) {
	radixApp := &radixv1.RadixApplication{}

	err := yaml.Unmarshal(raw, radixApp)
	if err != nil {
		return nil, fmt.Errorf("strict test failed: %v", err)
	}

	return radixApp, nil
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
	validateRadixConfigCmd.Flags().Bool(flagnames.StrictValidation, true, "Enable or disable strict schema validation, which will check for unknown fields in the radixconfig file")
	validateRadixConfigCmd.Flags().String(flagnames.Schema, "https://raw.githubusercontent.com/equinor/radix-operator/release/json-schema/radixapplication.json", "Validate against schema. http://, file:// or path is supported")

	// Allow but hide token-env flag so radix-github-actions won't interfere
	validateRadixConfigCmd.Flags().Bool(flagnames.TokenEnvironment, false, fmt.Sprintf("Take the token from environment variable %s", client.TokenEnvironmentName))
	err := validateRadixConfigCmd.Flags().MarkHidden(flagnames.TokenEnvironment)
	if err != nil {
		panic(err)
	}
}
