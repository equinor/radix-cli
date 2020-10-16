module github.com/equinor/radix-cli

go 1.13

require (
	github.com/equinor/radix-operator v1.5.27
	github.com/fatih/color v1.7.0
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/runtime v0.19.18
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.1.0
	golang.org/x/text v0.3.3 // indirect
	k8s.io/client-go v12.0.0+incompatible
)

replace (
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.0.0-20190818123050-43acd0e2e93f
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
