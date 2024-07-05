.PHONY: build
build:
	go build -o bin/server ./cmd/server

.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down

.PHONY: run
run:
	go run ./cmd/server
