package controllers

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hizzuu/grpc-example-user/gen/pb"
	"github.com/hizzuu/grpc-example-user/internal/domain"
	mock_interactor "github.com/hizzuu/grpc-example-user/internal/usecase/interactor"
)

func TestNewUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock_interactor.NewMockUserInteractor(ctrl)
	type args struct {
		userInteractor *mock_interactor.MockUserInteractor
	}
	tests := []struct {
		name string
		args args
		want *UserController
	}{
		{
			name: "UserController must be returned.",
			args: args{
				userInteractor: mock,
			},
			want: &UserController{
				userInteractor: mock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.userInteractor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_GetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *pb.GetUserReq
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(m *mock_interactor.MockUserInteractor)
		want          *pb.GetUserRes
		wantErr       bool
	}{
		{
			name: "Ability to retrieve specified user information.",
			args: args{
				ctx: context.Background(),
				r:   &pb.GetUserReq{Id: 1},
			},
			prepareMockFn: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().Get(context.Background(), int64(1)).Return(&domain.User{ID: 1}, nil)
			},
			want: &pb.GetUserRes{User: &pb.User{Id: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_interactor.NewMockUserInteractor(ctrl)
			tt.prepareMockFn(mock)

			c := &UserController{
				userInteractor: mock,
			}

			got, err := c.GetUser(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.User.Id, tt.want.User.Id) {
				t.Errorf("UserController.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *pb.CreateUserReq
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(m *mock_interactor.MockUserInteractor)
		want          *pb.CreateUserRes
		wantErr       bool
	}{
		{
			name: "Must be able to register if name is less than 20 characters",
			args: args{
				ctx: context.Background(),
				r: &pb.CreateUserReq{
					Email:    "test+1@example.com",
					Name:     strings.Repeat("a", 20),
					Password: "1234abcd",
				},
			},
			prepareMockFn: func(m *mock_interactor.MockUserInteractor) {
				m.EXPECT().Create(context.Background(), &domain.User{
					Email: "test+1@example.com",
					Name:  strings.Repeat("a", 20),
				}).Return(&domain.User{Email: "test+1@example.com"}, nil)
			},
			want: &pb.CreateUserRes{
				User: &pb.User{Email: "test+1@example.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_interactor.NewMockUserInteractor(ctrl)
			tt.prepareMockFn(mock)

			c := &UserController{
				userInteractor: mock,
			}

			got, err := c.CreateUser(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.User.Email, tt.want.User.Email) {
				t.Errorf("UserController.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
