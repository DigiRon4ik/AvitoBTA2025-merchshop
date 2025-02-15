# Импортирует переменные окружения из .env для доступа к ним из Makefile
ifneq (,$(wildcard .env))
    include .env
    export
endif

# @docker-compose run --rm migrate -path /migrations -database "$(DSN)" down

.PHONY: m-up m-down m-status d-up d-down d-down-v d-up-app d-up-b lint

DSN=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

m-up:
	@docker-compose run --rm migrate -path /migrations -database "$(DSN)" up

m-down:
	@docker-compose run --rm migrate -path /migrations -database "$(DSN)" down -all

m-status:
	@docker-compose run --rm migrate -path /migrations -database "$(DSN)" version

d-up:
	@docker-compose up -d

d-up-b:
	@docker-compose up -d --build

d-down:
	@docker-compose down

d-down-v:
	@docker-compose down -v

d-up-app:
	@docker-compose up -d postgres app

lint:
	@golangci-lint run
