package controllers

import (
	"context"

	"github.com/hizzuu/grpc-example-user/gen/pb"
	"github.com/hizzuu/grpc-example-user/internal/domain"
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

	return convertUserProto(user)
}

func (c *UserController) CreateUser(ctx context.Context, r *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	user := &domain.User{}
	switch x := r.Data.(type) {
	case *pb.CreateUserReq_Info:
		user.Email = x.Info.Email
		user.Name = x.Info.Name
	}
	user, err := c.userInteractor.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	userProto, err := convertUserProto(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserRes{User: userProto}, nil
}
