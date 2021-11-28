package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph/generated"
	"github.com/hizzuu/grpc-example-bff/internal/graph/model"
)

func (r *queryResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	userRes, err := r.userClient.GetUserByEmailAndPassword(ctx, &pb.GetUserByEmailAndPasswordReq{Email: input.Email, Password: input.Password})
	if err != nil {
		return nil, err
	}

	authRes, err := r.authClient.CreateToken(ctx, &pb.CreateTokenReq{Uid: strconv.FormatInt(userRes.User.Id, 10)})
	if err != nil {
		return nil, err
	}

	return &model.LoginPayload{
		User:  userRes.User,
		Token: authRes.Token,
	}, nil
}

func (r *queryResolver) RefreshIDToken(ctx context.Context) (*pb.Token, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
