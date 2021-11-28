package controllers

import (
	"context"

	"github.com/hizzuu/grpc-example-authority/gen/pb"
	"github.com/hizzuu/grpc-example-authority/internal/usecase/interactor"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthorityController struct {
	authInteractor interactor.AuthorityInteractor
}

func NewAuthorityController(autInteractor interactor.AuthorityInteractor) *AuthorityController {
	return &AuthorityController{
		autInteractor,
	}
}

func (c *AuthorityController) CreateToken(ctx context.Context, r *pb.CreateTokenReq) (*pb.CreateTokenRes, error) {
	t, err := c.authInteractor.CreateToken(ctx, r.Uid)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTokenRes{
		Token: &pb.Token{
			IdToken:      string(t.IDToken),
			RefreshToken: string(t.RefreshToken),
		},
	}, nil
}

func (c *AuthorityController) ListPublicKeys(ctx context.Context, in *emptypb.Empty) (*pb.ListPublicKeysRes, error) {
	jwks, err := c.authInteractor.ListPublicKeys()
	if err != nil {
		return nil, err
	}
	return &pb.ListPublicKeysRes{Jwks: jwks}, nil
}
