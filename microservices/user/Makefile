## Meta info

PJT_NAME = $(notdir $(PWD))
NET = grpc_example_net
SCHEMA_PATH = $$(dirname $$(dirname $$(pwd)))/schema
SVC = user
GRPC_SVC = UserService
DB_SVC = user-db
DB_NAME = user_db

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

## Generate by proto
.PHONY: proto
proto:
	docker run --rm --name protoc -v $(PWD)/gen/pb:/pb -v $(SCHEMA_PATH)/proto:/schema hizzuu/protoc \
		-I/schema \
		-I/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.2 \
		--go_out=plugins=grpc:/pb \
		--validate_out="lang=go:/pb" \
		user.proto

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)