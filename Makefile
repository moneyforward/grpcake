.DEFAULT_GOAL := build

SHELL  = /bin/bash
export GOBIN := $(PWD)/bin
export PATH  := $(GOBIN):$(PATH)

# Whether to run test server with reflection or not
REFLECT?=false

install-tools:
	awk '/_ ".+"/ {print $$2}' $(PWD)/tools/tools.go | xargs go install
.PHONY:install-tools

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	${GOBIN}/golangci-lint run ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o bin/grpcake cmd/main.go 
.PHONY:build

run-test-server:
	go run internal/testing/cmd/testserver/main.go -use-reflection=$(REFLECT)

