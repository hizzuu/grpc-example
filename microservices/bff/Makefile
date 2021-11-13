## Meta info

PJT_NAME = $(notdir $(PWD))
NET = bff_net
SCHEMA_PATH = $$(dirname $$(dirname $$(pwd)))/schema/proto
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
	docker run --rm --name protoc -v $(PWD)/gen/pb:/pb -v $(SCHEMA_PATH):/proto hizzuu/protoc \
	-I/proto \
	--go_out=plugins=grpc:/pb \
	user.proto

## Generate by graphql
.PHONY: gql
gql:
	docker-compose exec $(SVC) go run github.com/99designs/gqlgen generate


## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)