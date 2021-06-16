all: build test check

modules:
	go mod tidy

build: modules
	go build cmd/movieservice/*

test:
	go test ./...

check:
	golangci-lint run