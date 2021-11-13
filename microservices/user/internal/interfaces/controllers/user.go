package controllers

import (
	"context"

	"github.com/hizzuu/grpc-sample-user/gen/pb"
	"github.com/hizzuu/grpc-sample-user/internal/usecase/interactor"
)

type UserController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(userInteractor interactor.UserInteractor) *UserController {
	return &UserController{
		userInteractor: userInteractor,
	}
}

func (c *UserController) GetUser(ctx context.Context, r *pb.GetUserReq) (*pb.GetUserRes, error) {
	user, err := c.userInteractor.Get(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	userProto, err := convUserProto(user)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserRes{User: userProto}, nil
}
