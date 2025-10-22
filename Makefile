ENVIRONMENT ?= dev
VERSION ?= 0.0.1
RELEASE_NOTE ?= "First release"

.PHONY: swagger
swagger: bootstrap
	mkdir -p ./generated/radixapi
	mkdir -p ./generated/vulnscanapi
	swagger generate client -t ./generated/radixapi -f https://api.radix.equinor.com/swaggerui/swagger.json -A radixapi
	swagger generate client -t ./generated/vulnscanapi -f https://server-radix-vulnerability-scanner-api-prod.radix.equinor.com/swaggerui/swagger.json -A vulnscanapi
	go mod tidy

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
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.5.0
endif
ifndef HAS_GORELEASER
	go install github.com/goreleaser/goreleaser/v2@v2.12.6
endif
