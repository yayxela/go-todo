include .env

PROJECT_DIR = $(shell pwd)
SERVER_DIR = $(PROJECT_DIR)/cmd/server
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
SERVER = $(PROJECT_BIN)/server

.PHONY: .install-linter
.install-linter:
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.57.2

.PHONY: lint
lint: .install-linter
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yaml

.PHONY swagger:
swagger:
	swag init -g cmd/server/main.go --parseInternal --parseDepth 1 -o api/swagger

.PHONY docker-run:
docker-build:
	docker compose up -d --build

.PHONY build:
build:
	go build -o $(SERVER) $(SERVER_DIR)/main.go

.PHONY run:
run: build
	$(SERVER)
