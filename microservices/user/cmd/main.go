package main

import (
	"log"
	"net"

	"github.com/hizzuu/grpc-sample-user/internal/infrastructure"
	"github.com/hizzuu/grpc-sample-user/internal/interfaces/controllers"
	"github.com/hizzuu/grpc-sample-user/internal/interfaces/repository"
	"github.com/hizzuu/grpc-sample-user/internal/usecase/interactor"
)

func main() {
	conn, err := infrastructure.NewMysqlDB()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	sqlHandler := infrastructure.NewSqlHandler(conn)

	userCtrl := controllers.NewUserController(
		*interactor.NewUserInteractor(
			repository.NewUserRepository(sqlHandler),
		),
	)
	srv := infrastructure.NewGrpcServer(userCtrl)
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panic(err)
	}
	if err := srv.Serve(listen); err != nil {
		log.Panic(err)
	}
}
