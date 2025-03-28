# SaladCloud REST API

OPENAPI_SPEC ?= openapi.yaml
OPENAPI_VERSION := v7.12.0

.PHONY: build
build: api/openapi.yaml
	docker run \
		--name openapi-generator \
		--rm \
		--volume "${PWD}:/local" \
		--workdir /local \
		openapitools/openapi-generator-cli:$(OPENAPI_VERSION) generate \
			-c config.json \
			-i $(OPENAPI_SPEC) \
			--git-user-id SaladTechnologies \
			--git-repo-id salad-client

.PHONY: lint
lint:
	golangci-lint run ./... $(LINT_ARGS)

.PHONY: test
test:
	go test -v ./...
