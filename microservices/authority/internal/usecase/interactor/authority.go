package interactor

import (
	"context"
	"encoding/json"

	"github.com/hizzuu/grpc-example-authority/conf"
	"github.com/hizzuu/grpc-example-authority/internal/domain"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authorityInteractor struct {
}

type AuthorityInteractor interface {
	CreateToken(ctx context.Context, uid string) (*domain.Token, error)
	ListPublicKeys() (string, error)
}

func NewAuthorityInteractor() AuthorityInteractor {
	return &authorityInteractor{}
}

func (i *authorityInteractor) CreateToken(ctx context.Context, uid string) (*domain.Token, error) {
	t, err := genToken(uid)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	return t, nil
}

func (i *authorityInteractor) ListPublicKeys() (string, error) {
	key, err := jwk.New(conf.CredentialsConf.PrivateKey.PublicKey)
	if err != nil {
		return "", status.Error(codes.Internal, "failed to create jwk")
	}

	if err = key.Set(jws.KeyIDKey, "test-kid"); err != nil {
		return "", status.Error(codes.Internal, "failed to create jwks")
	}

	if err = key.Set(jws.AlgorithmKey, jwa.RS256); err != nil {
		return "", status.Error(codes.Internal, "failed to create jwks")
	}

	set := jwk.NewSet()
	set.Add(key)

	buf, err := json.Marshal(set)
	if err != nil {
		return "", status.Error(codes.Internal, "failed to create jwks")
	}

	return string(buf), nil
}
