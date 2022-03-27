package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/hizzuu/grpc-example-bff/gen/graph/directive"
	"github.com/hizzuu/grpc-example-bff/gen/graph/generated"
	"github.com/hizzuu/grpc-example-bff/gen/graph/middleware"
)

func New(r generated.ResolverRoot) *handler.Server {
	d := directive.New()
	m := middleware.New()

	conf := generated.Config{Resolvers: r}
	conf.Directives.Authentication = d.Authentication
	conf.Directives.Constraint = d.Constraint

	schema := generated.NewExecutableSchema(conf)

	srv := handler.NewDefaultServer(schema)
	srv.SetErrorPresenter(m.ErrorPresenter())

	return srv
}
