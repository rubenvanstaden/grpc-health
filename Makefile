SHELL := /bin/bash

INTEGRATION_TEST_PATH=./grpc/...

fmt:
	gofmt -l -s -w .

tidy:
	GOPROXY="" go mod tidy -compat=1.17

protos:
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pb/*.proto

build:
	GOOS="linux" go build -o bin/health-cli ./cmd/*

test.integration:
	go test -count=1 -run=Integration $(INTEGRATION_TEST_PATH)

test.integration.debug:
	go test -count=1 -run=Integration $(INTEGRATION_TEST_PATH) -v
