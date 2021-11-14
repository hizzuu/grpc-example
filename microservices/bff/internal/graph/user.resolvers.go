package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/hizzuu/grpc-sample-bff/gen/pb"
	"github.com/hizzuu/grpc-sample-bff/internal/graph/generated"
)

func (r *queryResolver) GetUser(ctx context.Context, id int64) (*pb.User, error) {
	in := &pb.GetUserReq{
		Id: id,
	}
	res, err := r.userClient.GetUser(ctx, in)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &pb.User{
		Id: res.User.Id,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
