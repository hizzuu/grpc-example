package main

import (
	"context"
	"log"

	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph"
	"github.com/hizzuu/grpc-example-bff/internal/infrastructure"
	"github.com/hizzuu/grpc-example-bff/utils/logger"
	"google.golang.org/grpc"
)

func init() {
	if err := logger.NewLogger(); err != nil {
		log.Panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "user:50051", grpc.WithInsecure())
	if err != nil {
		logger.Log.Panic(err)
	}
	defer conn.Close()

	resolver := graph.NewResolver(
		pb.NewUserServiceClient(conn),
	)
	infrastructure.ListenAndServe(resolver)
}
