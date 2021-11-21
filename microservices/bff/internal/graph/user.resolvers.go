package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hizzuu/grpc-example-bff/gen/pb"
	"github.com/hizzuu/grpc-example-bff/internal/graph/generated"
	"github.com/hizzuu/grpc-example-bff/utils/logger"
)

func (r *queryResolver) GetUser(ctx context.Context, id int64) (*pb.User, error) {
	logger.Log.Debug("start GetUser resolver")
	return r.userClient.GetUser(ctx, &pb.GetUserReq{Id: id})
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
