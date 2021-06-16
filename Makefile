all: build test check

modules:
	go mod tidy

build: modules
	go build cmd/movieservice/*

test:
	go test ./...

check:
	golangci-lint run

newman run /app/bin/Movieservice.postman_collection.json --global-var "localhost=http://localhost:8000"