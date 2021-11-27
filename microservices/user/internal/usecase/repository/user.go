package repository

import (
	"context"

	"github.com/hizzuu/grpc-example-user/internal/domain"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) (*domain.User, error)
}
