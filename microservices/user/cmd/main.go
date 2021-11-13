package main

import (
	"log"

	"github.com/hizzuu/grpc-sample-user/internal/infrastructure"
)

func main() {
	conn, err := infrastructure.NewMysqlDB()
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	sqlHandler := infrastructure.NewSqlHandler(conn)
	log.Println(sqlHandler)
}
