//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package interactor

import (
	"context"

	"github.com/hizzuu/grpc-example-user/internal/domain"
	"github.com/hizzuu/grpc-example-user/internal/usecase/repository"
)

type userInteractor struct {
	userRepo repository.UserRepository
}

type UserInteractor interface {
	Get(ctx context.Context, id int64) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) (*domain.User, error)
}

func NewUserInteractor(userRepo repository.UserRepository) *userInteractor {
	return &userInteractor{
		userRepo: userRepo,
	}
}

func (i *userInteractor) Get(ctx context.Context, id int64) (*domain.User, error) {
	return i.userRepo.Get(ctx, id)
}

func (i *userInteractor) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	return &domain.User{}, nil
}
