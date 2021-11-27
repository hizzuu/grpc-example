package middleware

import "github.com/hizzuu/grpc-example-bff/gen/pb"

type Middleware struct {
	authClient pb.AuthorityServiceClient
}

func NewMiddleware(authClient pb.AuthorityServiceClient) *Middleware {
	return &Middleware{
		authClient: authClient,
	}
}
