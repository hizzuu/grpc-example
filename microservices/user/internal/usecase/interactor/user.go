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
	GetByEmailAndPassword(ctx context.Context, email string, password string) (*domain.User, error)
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

func (i *userInteractor) GetByEmailAndPassword(ctx context.Context, email string, password string) (*domain.User, error) {
	user, err := i.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := compareHashAndPassword(user.EncryptedPassword, password); err != nil {
		return nil, err
	}

	return user, nil
}

func (i *userInteractor) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	encryptedPassword, err := genEncryptedPassword(u.Password)
	if err != nil {
		return nil, err
	}

	u.EncryptedPassword = encryptedPassword

	user, err := i.userRepo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}
