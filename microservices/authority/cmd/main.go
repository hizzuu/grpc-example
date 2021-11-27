package main

import (
	"log"
	"net"

	"github.com/hizzuu/grpc-example-authority/internal/infrastructure"
	"github.com/hizzuu/grpc-example-authority/internal/interfaces/controllers"
	"github.com/hizzuu/grpc-example-authority/internal/usecase/interactor"
	"github.com/hizzuu/grpc-example-authority/utils/logger"
)

func main() {
	authCtrl := controllers.NewAuthorityController(
		interactor.NewAuthorityInteractor(),
	)
	srv := infrastructure.NewGrpcServer(authCtrl)
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log.Panic(err)
	}
	if err := srv.Serve(listen); err != nil {
		logger.Log.Panic(err)
	}
}

func init() {
	if err := logger.NewLogger(); err != nil {
		log.Panic(err)
	}
}
