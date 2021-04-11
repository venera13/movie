all: build test check

modules:
	go mod tidy

build: modules
	go build cmd/cinema/*

test:
	go test ./...

check:
	golangci-lint run