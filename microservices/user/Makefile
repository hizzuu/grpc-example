## Meta info

PJT_NAME = $(notdir $(PWD))
NET = user_net
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

## Generate by Proto
.PHONY: proto
proto:
	docker run --rm --name protoc -v $(PWD)/gen:/pb -v $(SCHEMA_PATH):/proto hizzuu/protoc \
	-I/proto \
	--go_out=plugins=grpc:/pb \
	user.proto

# proto:
# 	docker run --rm --name protoc -v $(CWD)/pb:/pb -v $(CWD)/schema:/proto hizzuu/protoc \
# 	-I/proto \
# 	-I/go/src/github.com/envoyproxy/protoc-gen-validate \
# 	--go_out=plugins=grpc:/pb \
# 	--validate_out="lang=go:/pb" \
# 	same.proto

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)