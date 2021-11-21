package controllers

import (
	"github.com/hizzuu/grpc-example-user/gen/pb"
	"github.com/hizzuu/grpc-example-user/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserProto(u *domain.User) (*pb.User, error) {
	return &pb.User{
		Id:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}, nil
}
