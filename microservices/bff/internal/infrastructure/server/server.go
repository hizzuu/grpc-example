package server

import (
	"net/http"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/infrastructure/server/handler"
)

type server struct{}

type Server interface {
	Listen() error
}

func New(
	srv *gqlHandler.Server,
	authClient pb.AuthorityServiceClient,
) *server {
	h := handler.New(authClient)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", h.Cors(h.Authentication(srv)))

	return &server{}
}

func (s *server) Listen() error {
	return http.ListenAndServe(":"+"8080", nil)
}
