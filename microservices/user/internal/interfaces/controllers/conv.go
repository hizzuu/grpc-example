package controllers

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/hizzuu/grpc-sample-user/gen/pb"
	"github.com/hizzuu/grpc-sample-user/internal/domain"
)

func convUserProto(u *domain.User) (*pb.User, error) {
	cAt, err := ptypes.TimestampProto(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	uAt, err := ptypes.TimestampProto(u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}
