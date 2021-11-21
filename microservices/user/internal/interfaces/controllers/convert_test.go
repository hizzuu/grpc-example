package controllers

import (
	"reflect"
	"testing"
	"time"

	"github.com/hizzuu/grpc-example-user/gen/pb"
	"github.com/hizzuu/grpc-example-user/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_convertUserProto(t *testing.T) {
	type args struct {
		u *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.User
		wantErr bool
	}{
		{
			name: "Ability to convert from model to proto.",
			args: args{
				u: &domain.User{
					ID:        1,
					Name:      "TEST USER",
					Email:     "test+1@example.com",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			want: &pb.User{
				Id:        1,
				Name:      "TEST USER",
				Email:     "test+1@example.com",
				CreatedAt: timestamppb.New(time.Now()),
				UpdatedAt: timestamppb.New(time.Now()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertUserProto(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertUserProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Id, tt.want.Id) {
				t.Errorf("convertUserProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
