package interactor

import (
	"context"

	"github.com/hizzuu/grpc-sample-user/internal/domain"
	"github.com/hizzuu/grpc-sample-user/internal/usecase/repository"
)

type UserInteractor struct {
	userRepo repository.UserRepository
}

func NewUserInteractor(userRepo repository.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepo: userRepo,
	}
}

func (i *UserInteractor) Get(ctx context.Context, id int64) (*domain.User, error) {
	return i.userRepo.Get(ctx, id)
}
