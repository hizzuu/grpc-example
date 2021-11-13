## Meta info

PJT_NAME = $(notdir $(PWD))
NET = bff_net
SCHEMA_PATH = $$(dirname $$(dirname $$(pwd)))/schema
SVC = bff
GRPC_SVC = BffService

## Container up
.PHONY: up
up:
	docker compose up

## Container down
.PHONY: down
down:
	docker compose down

# Cotainer attach
.PHONY: attach
attach:
	docker exec -it $(SVC) sh

## Lint
.PHONY: lint
lint:	
	docker compose run --rm $(SVC) go vet ./...

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)