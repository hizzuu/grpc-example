package directive

import (
	"context"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/grpc-example-bff/internal/graph/errors"
	"github.com/hizzuu/grpc-example-bff/internal/graph/model"
	"github.com/lestrrat-go/jwx/jwt"
)

func Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if err := getAuthErrFromCtx(ctx); err != nil {
		return nil, errors.AuthenticationError(err.Error())
	}

	token, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, errors.AuthenticationError(err.Error())
	}

	uid, err := getUIDFromToken(token)
	if err != nil {
		return nil, errors.AuthenticationError("failed to get uid")
	}

	int64UID, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, errors.AuthenticationError("failed to parse storing to int64")
	}

	return next(context.WithValue(ctx, model.CtxUIDKey, int64UID))
}

func getAuthErrFromCtx(ctx context.Context) error {
	err, ok := ctx.Value(model.CtxAuthErrorCtxKey).(error)
	if ok {
		return err
	}
	return nil
}

func getTokenFromCtx(ctx context.Context) (jwt.Token, error) {
	v := ctx.Value(model.CtxJwtTokenKey)

	token, ok := v.(jwt.Token)
	if !ok {
		return nil, fmt.Errorf("missing token in 'Authorization' header")
	}

	return token, nil
}

func getUIDFromToken(t jwt.Token) (string, error) {
	uid, ok := t.Get(string(model.UIDKey))
	if !ok {
		return "", errors.AuthenticationError("failed to get uid")
	}
	return uid.(string), nil
}
