package handler

import "github.com/hizzuu/grpc-example-bff/gen/pb"

type handler struct {
	authClient pb.AuthorityServiceClient
}

type Handler interface{}

func New(authClient pb.AuthorityServiceClient) *handler {
	return &handler{
		authClient: authClient,
	}
}
