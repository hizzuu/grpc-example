package infrastructure

import (
	"github.com/hizzuu/grpc-example-authority/gen/pb"
	"github.com/hizzuu/grpc-example-authority/internal/interfaces/controllers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(authCtrl *controllers.AuthorityController) *grpc.Server {
	srv := grpc.NewServer()
	pb.RegisterAuthorityServiceServer(srv, authCtrl)
	reflection.Register(srv)
	return srv
}
