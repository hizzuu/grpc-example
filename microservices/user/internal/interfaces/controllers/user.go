package controllers

import (
	"context"

	"github.com/hizzuu/grpc-example-user/gen/pb"
	"github.com/hizzuu/grpc-example-user/internal/usecase/interactor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(userInteractor interactor.UserInteractor) *UserController {
	return &UserController{
		userInteractor: userInteractor,
	}
}

func (c *UserController) GetUser(ctx context.Context, r *pb.GetUserReq) (*pb.User, error) {
	if err := r.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	user, err := c.userInteractor.Get(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	userProto, err := convertUserProto(user)
	if err != nil {
		return nil, err
	}

	return userProto, nil
}

func (c *UserController) CreateUser(ctx context.Context, r *pb.CreateUserReq) (*pb.User, error) {
	return &pb.User{}, nil
}
