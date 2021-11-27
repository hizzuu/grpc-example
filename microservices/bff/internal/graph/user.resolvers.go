package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph/generated"
	"github.com/hizzuu/grpc-example-bff/internal/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.CreateUserPayload, error) {
	res, err := r.userClient.CreateUser(ctx, &pb.CreateUserReq{Name: input.Name, Email: input.Email, Password: input.Password})
	if err != nil {
		return nil, err
	}

	return &model.CreateUserPayload{
		User: res.User,
	}, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id int64) (*model.GetUserPayload, error) {
	uid := ctx.Value(model.UIDKey).(int64)
	fmt.Println(uid, "UID")
	res, err := r.userClient.GetUser(ctx, &pb.GetUserReq{Id: id})
	if err != nil {
		return nil, err
	}

	return &model.GetUserPayload{
		User: res.User,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
