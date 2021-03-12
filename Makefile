include .env
export

migrate-up:
	migrate -database $(DATABASE_CONNECTION_STRING) -path db/migrations up

migrate-down:
	migrate -database $(DATABASE_CONNECTION_STRING) -path db/migrations down

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	golangci-lint run

fmt:
	go fmt ./...

test:
	go test ./...

run: fmt test
	go run cmd/*.go

build: fmt test
	go build cmd/*.go


.PHONY: migrate-up migrate-down cover lint fmt test run build