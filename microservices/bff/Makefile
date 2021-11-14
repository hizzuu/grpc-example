## Meta info

PJT_NAME = $(notdir $(PWD))
NET = grpc_sample_net
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

## Generate by Proto
.PHONY: proto
proto:
	docker run --rm --name protoc -v $(PWD)/gen/pb:/pb -v $(SCHEMA_PATH)/proto:/schema hizzuu/protoc \
	-I/schema \
	--go_out=plugins=grpc:/pb \
	user.proto

## Generate by graphql
.PHONY: gql
gql:
	docker compose run --rm -v $(SCHEMA_PATH)/graphql:/app/schema $(SVC) go run github.com/99designs/gqlgen generate
	rm -rf $(PWD)/schema

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
