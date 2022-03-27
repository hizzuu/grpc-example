package middleware

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/grpc-example-bff/gen/graph/gqlerr"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *middleware) ErrorPresenter() graphql.ErrorPresenterFunc {
	return graphql.ErrorPresenterFunc(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		s, ok := status.FromError(err.Unwrap())
		if !ok {
			return err
		}
		switch s.Code() {
		case codes.InvalidArgument, codes.AlreadyExists, codes.NotFound:
			err = gqlerr.UserInputError(e.Error())
		case codes.Unauthenticated:
			err = gqlerr.AuthenticationError(e.Error())
		case codes.PermissionDenied:
			err = gqlerr.ForbiddenError(e.Error())
		default:
			err = gqlerr.InternalServerError(e.Error())
		}
		err.Extensions["grpc_code"] = s.Code().String()
		return err
	})
}
