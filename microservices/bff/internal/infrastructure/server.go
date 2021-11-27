package infrastructure

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/grpc-example-bff/internal/graph/directive"
	"github.com/hizzuu/grpc-example-bff/internal/graph/errors"
	"github.com/hizzuu/grpc-example-bff/internal/graph/generated"
	"github.com/hizzuu/grpc-example-bff/internal/infrastructure/middleware"
	"github.com/hizzuu/grpc-example-bff/utils/logger"
)

const defaultPort = "8080"

func ListenAndServe(
	r generated.ResolverRoot,
	mw *middleware.Middleware,
) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conf := generated.Config{Resolvers: r}
	conf.Directives.Authentication = directive.Authentication
	conf.Directives.Constraint = directive.Constraint

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))
	srv.SetErrorPresenter(errors.ErrorPresenter())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query",
		mw.Authentication(
			srv,
		),
	)

	logger.Log.Fatal(http.ListenAndServe(":"+port, nil))
}
