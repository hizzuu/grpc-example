package errors

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorPresenter() graphql.ErrorPresenterFunc {
	return graphql.ErrorPresenterFunc(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		s, ok := status.FromError(err.Unwrap())
		if !ok {
			return err
		}
		switch s.Code() {
		case codes.InvalidArgument, codes.AlreadyExists, codes.NotFound:
			err = UserInputError(e.Error())
		case codes.Unauthenticated:
			err = AuthenticationError(e.Error())
		case codes.PermissionDenied:
			err = ForbiddenError(e.Error())
		default:
			err = InternalServerError(e.Error())
		}
		err.Extensions["grpc_code"] = s.Code().String()
		return err
	})
}
