ENVIRONMENT ?= dev
VERSION ?= 0.0.1
RELEASE_NOTE ?= "First release"

.PHONY: generate-client
generate-client:
	swagger generate client -t ./generated-client -f https://api.radix.equinor.com/swaggerui/swagger.json -A radixapi

.PHONY: release
release:
	swagger generate client -t ./generated-client -f https://api.radix.equinor.com/swaggerui/swagger.json -A radixapi
	git tag -a v$(VERSION) -m "$(RELEASE_NOTE)"
	git push origin v$(VERSION)
	git config --global credential.helper cache
	goreleaser --rm-dist
