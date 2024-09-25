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

install:
	go build ./cli/rx/
	mv rx $$(go env GOPATH)/bin/rx

HAS_SWAGGER       := $(shell command -v swagger;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_GORELEASER    := $(shell command -v goreleaser;)

bootstrap:
ifndef HAS_SWAGGER
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0
endif
ifndef HAS_GOLANGCI_LINT
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.2
endif
ifndef HAS_GORELEASER
	go install github.com/goreleaser/goreleaser@v1.26.2
endif
