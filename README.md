# Microservices pattern with gRPC Example
The Go language implementation of gRPC

## Architecture
- Clean Architecture
- Microservices pattern
- BFF Pattern (Backend for Frontend)

## Microservices
The implementation of each microservice can be found under the `/services` directory.
### BFF - Backend for Frontend
Call the relevant microservice APIs to get the data you need
- Golang
- Graphql
- gRPC

### User Service
Service for CRUD and authentication authorization of users
- Golang
- gRPC
- Middleware
  - MySQL

## Schema
The interfaces definition of protocol buffers and graphql can be found under the `/schema` directory.

- [Protocol Buffers schema document](https://github.com/hizzuu/grpc-example/tree/main/schema/proto)
