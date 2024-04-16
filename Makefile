include .env

.PHONY swagger:
	swag init -g cmd/server/main.go --parseInternal --parseDepth 1 -o api/swagger

.PHONY docker-build:
	docker compose up -d --build