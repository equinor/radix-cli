package main

import (
	"github.com/equinor/radix-cli/cmd"

	// Force loading of needed authentication library
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
)

func init() {
	// If you get GOAWAY calling API with token using:
	// az account get-access-token
	// ...enable this line
	// os.Setenv("GODEBUG", "http2server=0,http2client=0")

}

func main() {
	cmd.Execute()
}
