package middleware

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph/model"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func (m *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		idToken := getIDTokenFromHeader(r)
		if idToken == "" {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.CtxAuthErrorCtxKey, fmt.Errorf("failed to authenticate"))))
			return
		}

		res, err := m.authClient.ListPublicKeys(r.Context(), &pb.ListPublicKeysReq{})
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.CtxAuthErrorCtxKey, fmt.Errorf("failed to authenticate"))))
			return
		}

		token, err := verifyIDToken(idToken, res.Jwks)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.CtxAuthErrorCtxKey, fmt.Errorf("failed to authenticate"))))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.CtxJwtTokenKey, token)))
	})
}

func getIDTokenFromHeader(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	return strings.Replace(authorization, "Bearer ", "", 1)
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

	return token, nil
}
