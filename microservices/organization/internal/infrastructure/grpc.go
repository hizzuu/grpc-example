package infrastructure

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer() *grpc.Server {
	srv := grpc.NewServer()
	reflection.Register(srv)
	return srv
}
