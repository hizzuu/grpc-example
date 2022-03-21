package main

import (
	"log"
	"net"

	"github.com/hizzuu/grpc-example-user/internal/infrastructure"
	"github.com/hizzuu/grpc-example-user/internal/interfaces/controllers"
	"github.com/hizzuu/grpc-example-user/internal/interfaces/repository"
	"github.com/hizzuu/grpc-example-user/internal/usecase/interactor"
)

func main() {
	logger, err := infrastructure.NewLogger()
	if err != nil {
		log.Panic(err)
	}

	conn, err := infrastructure.NewMysqlDB()
	if err != nil {
		logger.Log.Panic(err)
	}
	defer conn.Close()

	sqlHandler := infrastructure.NewSqlHandler(conn)

	userCtrl := controllers.NewUserController(
		interactor.NewUserInteractor(
			repository.NewUserRepository(
				sqlHandler,
			),
		),
	)

	srv := infrastructure.NewGrpcServer(userCtrl)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log.Panic(err)
	}

	if err := srv.Serve(listen); err != nil {
		logger.Log.Panic(err)
	}
}
