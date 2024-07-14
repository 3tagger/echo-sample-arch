ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: build
build:
	go build -o bin/server ./cmd/server

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: run
run:
	go run ./cmd/server

.PHONY: seed-user
seed-user:
	go run ./cmd/seeder -n=10 -entity=user

.PHONY: seed
seed: seed-user

.PHONY: migrate-up
migrate-up:
	migrate -database "postgresql://$(POSTGRESQL_USERNAME):$(POSTGRESQL_PASSWORD)@$(POSTGRESQL_HOST):$(POSTGRESQL_PORT)/$(POSTGRESQL_DBNAME)?sslmode=disable" -verbose -path ./db/migrations/ up

.PHONY: migrate-down
migrate-down:
	migrate -database "postgresql://$(POSTGRESQL_USERNAME):$(POSTGRESQL_PASSWORD)@$(POSTGRESQL_HOST):$(POSTGRESQL_PORT)/$(POSTGRESQL_DBNAME)?sslmode=disable" -verbose -path ./db/migrations/ down

.PHONY: mock
mock:
	mockery

.PHONY: gen-doc
gen-doc:
	swag fmt
	swag init -g ../cmd/server/main.go -o ./docs -d ./internal