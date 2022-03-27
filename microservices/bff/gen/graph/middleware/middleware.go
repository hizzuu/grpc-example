package middleware

import (
	"github.com/99designs/gqlgen/graphql"
)

type middleware struct{}

type Middleware interface {
	ErrorPresenter() graphql.ErrorPresenterFunc
}

func New() *middleware {
	return &middleware{}
}
