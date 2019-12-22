module github.com/equinor/radix-cli

go 1.13

require (
	contrib.go.opencensus.io/exporter/ocagent v0.6.0 // indirect
	github.com/equinor/radix-operator v1.4.1
	github.com/fatih/color v1.7.0
	github.com/go-openapi/errors v0.19.3
	github.com/go-openapi/runtime v0.19.9
	github.com/go-openapi/strfmt v0.19.4
	github.com/go-openapi/swag v0.19.6
	github.com/go-openapi/validate v0.19.5
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	go.opencensus.io v0.22.2 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	k8s.io/apimachinery v0.0.0-20191020214737-6c8691705fc5
	k8s.io/client-go v12.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
