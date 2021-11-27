package main

import (
	"context"
	"log"

	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph"
	"github.com/hizzuu/grpc-example-bff/internal/infrastructure"
	"github.com/hizzuu/grpc-example-bff/internal/infrastructure/middleware"
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

	mw := middleware.NewMiddleware(authClient)

	resolver := graph.NewResolver(
		userClient,
		authClient,
	)

	infrastructure.ListenAndServe(resolver, mw)
}

func init() {
	if err := logger.NewLogger(); err != nil {
		log.Panic(err)
	}
}
