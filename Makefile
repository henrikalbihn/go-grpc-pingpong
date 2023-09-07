SHELL := /bin/bash
.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: serve
serve: ## Run the gRPC server.
	./bin/serve

.PHONY: ping
ping:	## Ping the gRPC server.
	./bin/ping

.PHONY: build
build: build-server build-ping build-utils build-pb

.PHONY: build-server
build-server:  ## Build the server executable.
	cd server && GOOS=linux GOARCH=amd64 go build -o ../bin/serve

.PHONY: build-ping
build-ping:  ## Build the ping executable.
	cd client && GOOS=linux GOARCH=amd64 go build -o ../bin/ping

.PHONY: build-utils
build-utils:  ## Build the utils executable.
	cd utils && GOOS=linux GOARCH=amd64 go build -o ../bin/utils

.PHONY: build-pb
build-pb:  ## Build the protobuf files.
	cd pb && GOOS=linux GOARCH=amd64 go build -o ../bin/pb

.PHONY: clean
clean:  ## Clean the bin directory.
	rm -rf bin/*
