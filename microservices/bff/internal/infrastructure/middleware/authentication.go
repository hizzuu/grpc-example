package middleware

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hizzuu/grpc-example-bff/internal/graph/model"
	"github.com/hizzuu/grpc-example-bff/utils/logger"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idToken, err := getIDTokenFromHeader(r)
		if err != nil {
			logger.Log.Infof("unauthenticate: %s", err.Error())
			next.ServeHTTP(w, r.WithContext(
				context.WithValue(
					ctx,
					model.CtxAuthErrorCtxKey,
					status.Error(codes.Unauthenticated, "unauthenticate"),
				),
			))
			return
		}

		res, err := m.authClient.ListPublicKeys(r.Context(), &emptypb.Empty{})
		if err != nil {
			logger.Log.Errorf("unauthenticate: %s", err.Error())
			next.ServeHTTP(w, r.WithContext(
				context.WithValue(
					ctx,
					model.CtxAuthErrorCtxKey,
					status.Error(codes.Unauthenticated, "unauthenticate"),
				),
			))
			return
		}

		token, err := verifyIDToken(idToken, res.Jwks)
		if err != nil {
			logger.Log.Errorf("unauthenticate: %s", err.Error())
			next.ServeHTTP(w, r.WithContext(
				context.WithValue(
					ctx,
					model.CtxAuthErrorCtxKey,
					status.Error(codes.Unauthenticated, "unauthenticate"),
				),
			))
			return
		}

		next.ServeHTTP(w, r.WithContext(
			context.WithValue(
				ctx,
				model.CtxJwtTokenKey,
				token,
			),
		))
	})
}

func getIDTokenFromHeader(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return "", fmt.Errorf("failed to get Authorization from header")
	}

	idToken := strings.TrimPrefix(authorization, "Bearer ")
	if idToken == "" {
		return "", fmt.Errorf("failed to get token from authorization")
	}

	return idToken, nil
}

func verifyIDToken(idToken string, jwks string) (jwt.Token, error) {
	key, err := jwk.Parse(bytes.NewBufferString(jwks).Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwks")
	}

	token, err := jwt.Parse([]byte(idToken), jwt.WithKeySet(key))
	if err != nil {
		return nil, fmt.Errorf("failed to verify token")
	}

	if time.Now().After(token.Expiration().Local()) {
		return nil, fmt.Errorf("token has expired")
	}

	return token, nil
}
