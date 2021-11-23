package main

import (
	"fmt"
	"log"

	"github.com/hizzuu/grpc-example-organization/utils/logger"
)

func init() {
	if err := logger.NewLogger(); err != nil {
		log.Panic(err)
	}
}

func main() {
	fmt.Println("###")
}
