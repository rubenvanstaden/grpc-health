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
	go build -o bin/health-cli ./cmd/cli/*
