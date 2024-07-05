.PHONY: build
build:
	go build -o bin/server ./cmd/server

.PHONY: run
run:
	go run ./cmd/server
