module github.com/equinor/radix-cli

go 1.16

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/equinor/radix-operator v1.13.2
	github.com/fatih/color v1.9.0
	github.com/go-openapi/analysis v0.20.1 // indirect
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/runtime v0.19.29
	github.com/go-openapi/strfmt v0.20.1
	github.com/go-openapi/swag v0.19.15
	github.com/go-openapi/validate v0.20.2
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.0
	go.mongodb.org/mongo-driver v1.6.0 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	k8s.io/client-go v12.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.19.9
