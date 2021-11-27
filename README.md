# Microservices pattern with gRPC Example
The Go language implementation of gRPC

## Architecture
- Clean Architecture
- Microservices pattern
- BFF Pattern (Backend for Frontend)
- JWT (JSON Web Token)

## Microservices
The implementation of each microservice can be found under the `/services` directory.
### BFF - Backend for Frontend
Call the relevant microservice APIs to get the data you need
- Golang
- Graphql
- gRPC

### User Service
Service for create, update, and fetch users
- Golang
- gRPC
- Middleware
  - MySQL

### Authority Service
Service for authentication authorization
- Golang
- gRPC
- JWT

## Schema
The interfaces definition of protocol buffers and graphql can be found under the `/schema` directory.

### Protocol Buffers
- [Docker Image Repository for using Protoc configured in Go](https://github.com/hizzuu/protoc)
- [Protocol Buffers schema document](https://github.com/hizzuu/grpc-example/tree/main/schema/proto)
