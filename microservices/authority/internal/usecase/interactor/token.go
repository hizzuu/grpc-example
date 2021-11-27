package interactor

import (
	"encoding/json"
	"time"

	"github.com/hizzuu/grpc-example-authority/conf"
	"github.com/hizzuu/grpc-example-authority/internal/domain"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

func genToken(userID string) (*domain.Token, error) {
	idToken, err := genJwt(userID, domain.IdTokenSub, domain.IdTokenExpSec)
	if err != nil {
		return nil, err
	}

	refreshToken, err := genJwt(userID, domain.RefreshTokenSub, domain.RefreshTokenExpSec)
	if err != nil {
		return nil, err
	}

	return &domain.Token{
		IDToken:      domain.TokenType(idToken),
		RefreshToken: domain.TokenType(refreshToken),
	}, nil
}

func genJwt(userID string, sub string, expSec int64) (string, error) {
	exp := time.Now().Add(time.Duration(expSec) * time.Second).Unix()

	token := jwt.New()
	token.Set(jwt.ExpirationKey, exp)
	token.Set(jwt.SubjectKey, sub)
	token.Set(domain.UIDKey, userID)

	headers := jws.NewHeaders()
	headers.Set(jws.AlgorithmKey, jwa.RS256)
	headers.Set(jws.TypeKey, domain.TypeName)
	headers.Set(jws.KeyIDKey, "test-kid")

	return signToken(token, headers)
}

func signToken(t jwt.Token, h jws.Headers) (string, error) {
	buf, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	signedToken, err := jws.Sign(buf, jwa.RS256, conf.CredentialsConf.PrivateKey, jws.WithHeaders(h))
	if err != nil {
		return "", err
	}
	return string(signedToken), nil
}
