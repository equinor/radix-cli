package main

import (
	"os"

	"github.com/equinor/radix-cli/cmd"
	radixconfig "github.com/equinor/radix-cli/pkg/config"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func init() {
	// If you get GOAWAY calling API with token using:
	// az account get-access-token
	// ...enable this line
	// os.Setenv("GODEBUG", "http2server=0,http2client=0")

}

func main() {
	klog.SetLogger(klog.New(log.NullLogSink{})) // HACK: Temporarily disable client-go warning https://github.com/kubernetes/client-go/blob/c2f61ae20ae1b13893992f7ceadd6304ba7025e3/plugin/pkg/client/auth/azure/azure.go#L91
	ensureRadixConfigFilesExist()
	cmd.Execute()
}

func ensureRadixConfigFilesExist() error {
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
