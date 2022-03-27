package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type directive struct {
}

type Directive interface {
	Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error)
	Constraint(ctx context.Context, obj interface{}, next graphql.Resolver, label string, notEmpty *bool, notBlank *bool, pattern *string, min *int, max *int) (interface{}, error)
}

func New() *directive {
	return &directive{}
}
