package usecase

import (
	"context"

	"github.com/koba1108/go-mongodb/internals/domain/model"
	"github.com/koba1108/go-mongodb/internals/domain/repository"
)

type UserUsecase interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (uu *userUsecase) FindByID(ctx context.Context, id string) (*model.User, error) {
	return uu.userRepo.FindByID(ctx, id)
}

func (uu *userUsecase) FindAll(ctx context.Context) ([]*model.User, error) {
	return uu.userRepo.FindAll(ctx)
}
