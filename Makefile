-include .env
export

run:
	go run ./cmd/api

build:
	go build -a -o ./bin/api ./cmd/api && go build -a -o ./bin/migrate ./cmd/migrate

test:
	go test -v ./...

migrate:
	goose up
