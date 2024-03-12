package main

import (
	"fmt"
	"os"

	"github.com/equinor/radix-cli/cmd"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
)

func init() {
	// If you get GOAWAY calling API with token using:
	// az account get-access-token
	// ...enable this line
	// os.Setenv("GODEBUG", "http2server=0,http2client=0")

}

func main() {
	err := ensureRadixConfigFolderExists()
	if err != nil {
		fmt.Printf("Error creating radix config folder: %v\n", err)
		os.Exit(1)
	}

	cmd.Execute()
}

func ensureRadixConfigFolderExists() error {
	if _, err := os.Stat(radixconfig.RadixConfigDir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err = os.MkdirAll(radixconfig.RadixConfigDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
