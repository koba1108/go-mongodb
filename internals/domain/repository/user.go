package repository

import (
	"context"

	"github.com/koba1108/go-mongodb/internals/domain/model"
)

type UserRepository interface {
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
}
