package infrastructure

import (
	"github.com/hizzuu/grpc-example-user/gen/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(userCtrl pb.UserServiceServer) *grpc.Server {
	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, userCtrl)
	reflection.Register(srv)
	return srv
}
