all: build test check

modules:
	go mod tidy

build: modules
	go build -v -o bin/movieservice cmd/movieservice/*.go

test:
	go test ./...

check:
	golangci-lint run

newman run /app/bin/Movieservice.postman_collection.json --global-var "localhost=http://localhost:8000"