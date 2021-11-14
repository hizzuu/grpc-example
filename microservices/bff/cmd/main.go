package main

import (
	"context"
	"log"

	"github.com/hizzuu/grpc-sample-bff/gen/pb"
	"github.com/hizzuu/grpc-sample-bff/internal/graph"
	"github.com/hizzuu/grpc-sample-bff/internal/infrastructure"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "user:50051", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	resolver := graph.NewResolver(
		pb.NewUserServiceClient(conn),
	)
	infrastructure.ListenAndServe(resolver)
}
