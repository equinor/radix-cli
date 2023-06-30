package main

import (
	"github.com/equinor/radix-cli/cmd"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/log"

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
	klog.SetLogger(klog.New(log.NullLogSink{})) // HACK: Temporarily disable client-go warning https://github.com/kubernetes/client-go/blob/c2f61ae20ae1b13893992f7ceadd6304ba7025e3/plugin/pkg/client/auth/azure/azure.go#L91
	cmd.Execute()
}
