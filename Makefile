ENVIRONMENT ?= dev
VERSION ?= 0.0.1
RELEASE_NOTE ?= "First release"

.PHONY: swagger
swagger: bootstrap
	swagger generate client -t ./generated-client -f https://api.radix.equinor.com/swaggerui/swagger.json -A radixapi

.PHONY: push
push:
	docker build . -t ghcr.io/equinor/radix/rx:latest
	docker login ghcr.io/equinor
	docker push ghcr.io/equinor/radix/rx:latest

.PHONY: lint
lint: bootstrap
	golangci-lint run --max-same-issues 0

HAS_SWAGGER       := $(shell command -v swagger;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)

bootstrap:
ifndef HAS_SWAGGER
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
endif
ifndef HAS_GOLANGCI_LINT
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
endif
