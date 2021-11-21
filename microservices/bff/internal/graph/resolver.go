package graph

import (
	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userClient pb.UserServiceClient
}

func NewResolver(
	userClient pb.UserServiceClient,
) generated.ResolverRoot {
	return &Resolver{
		userClient,
	}
}
