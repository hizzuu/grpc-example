package main

import (
	"context"
	"fmt"

	"github.com/hizzuu/grpc-example-bff/gen/graph"
	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/infrastructure/graphql"
	"github.com/hizzuu/grpc-example-bff/internal/infrastructure/server"
	"github.com/hizzuu/grpc-example-bff/utils/logger"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "user:50051", grpc.WithInsecure())
	if err != nil {
		logger.Log.Panic(err)
	}
	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)

	conn, err = grpc.DialContext(ctx, "authority:50051", grpc.WithInsecure())
	if err != nil {
		logger.Log.Panic(err)
	}
	defer conn.Close()

	authClient := pb.NewAuthorityServiceClient(conn)

	resolver := graph.NewResolver(
		userClient,
		authClient,
	)

	srv := graphql.New(resolver)

	s := server.New(srv, authClient)

	if err := s.Listen(); err != nil {
		fmt.Println(err.Error())
	}
}
