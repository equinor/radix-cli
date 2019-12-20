package main

import (
	"github.com/equinor/radix-cli/cmd"

	// Force loading of needed authentication library
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
)

func main() {
	cmd.Execute()
}
