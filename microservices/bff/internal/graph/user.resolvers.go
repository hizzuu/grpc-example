package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph/generated"
	"github.com/hizzuu/grpc-example-bff/internal/graph/model"
	"github.com/hizzuu/grpc-example-bff/utils/logger"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.CreateUserPayload, error) {
	userProto, err := r.userClient.CreateUser(ctx, &pb.CreateUserReq{Name: input.Name, Email: input.Email, Password: input.Password})
	if err != nil {
		return nil, err
	}
	return &model.CreateUserPayload{
		User:    userProto.User,
		IDToken: userProto.IdToken,
	}, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id int64) (*pb.User, error) {
	logger.Log.Debug("start GetUser resolver")
	return r.userClient.GetUser(ctx, &pb.GetUserReq{Id: id})
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
