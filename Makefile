all: build test check

up:
	docker-compose -f docker-compose.yml up -d

down:
	docker-compose -f docker-compose.yml down

modules:
	go mod tidy

build: modules
	go build -v -o bin/movieservice cmd/movieservice/*.go

test:
	go test ./...

check:
	golangci-lint run

api_tests: up
	docker run --rm -v $(shell pwd)/api-tests:/app --network host postman/newman run --global-var url=localhost:8000 /app/Movieservice.postman_collection.json