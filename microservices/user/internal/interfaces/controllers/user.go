package controllers

import (
	"context"
	"time"

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

func (c *UserController) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.GetUserRes, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user, err := c.userInteractor.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	pbUser, err := convertUserProto(user)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserRes{
		User: pbUser,
	}, nil
}

func (c *UserController) GetUserByEmailAndPassword(ctx context.Context, r *pb.GetUserByEmailAndPasswordReq) (*pb.GetUserRes, error) {
	user, err := c.userInteractor.GetByEmailAndPassword(ctx, r.Email, r.Password)
	if err != nil {
		return nil, err
	}

	pbUser, err := convertUserProto(user)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserRes{
		User: pbUser,
	}, nil
}

func (c *UserController) CreateUser(ctx context.Context, r *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	nowDt := time.Now()
	user := &domain.User{
		Email:     r.Email,
		Name:      r.Name,
		Password:  r.Password,
		CreatedAt: nowDt,
		UpdatedAt: nowDt,
	}

	user, err := c.userInteractor.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	pbUser, err := convertUserProto(user)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserRes{User: pbUser}, nil
}
