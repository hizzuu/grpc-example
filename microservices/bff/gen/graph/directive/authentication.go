package directive

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/grpc-example-bff/gen/graph/gqlerr"
	"github.com/hizzuu/grpc-example-bff/gen/graph/model"
	"github.com/lestrrat-go/jwx/jwt"
)

func (d *directive) Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if err := getAuthErrFromCtx(ctx); err != nil {
		return nil, gqlerr.AuthenticationError(err.Error())
	}

	token, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError(err.Error())
	}

	claims, err := getClaimsFromToken(token)
	if err != nil {
		return nil, gqlerr.AuthenticationError("failed to get claims")
	}

	return next(context.WithValue(ctx, model.CtxClaimsKey, claims))
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

func getClaimsFromToken(t jwt.Token) (*model.JwtClaims, error) {
	v, ok := t.Get(string(model.ClaimsKey))
	if !ok {
		return nil, fmt.Errorf("failed to get claims from context")
	}

	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal")
	}

	var c *model.JwtClaims
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal")
	}

	return c, nil
}
