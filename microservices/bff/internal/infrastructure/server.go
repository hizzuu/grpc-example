package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/grpc-sample-bff/internal/graph/generated"
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
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
