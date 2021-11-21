package infrastructure

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/grpc-example-bff/internal/graph/errors"
	"github.com/hizzuu/grpc-example-bff/internal/graph/generated"
	"github.com/hizzuu/grpc-example-bff/utils/logger"
)

const defaultPort = "8080"

func ListenAndServe(r generated.ResolverRoot) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: r,
			},
		),
	)

	srv.SetErrorPresenter(errors.ErrorPresenter())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	logger.Log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Log.Fatal(http.ListenAndServe(":"+port, nil))
}
