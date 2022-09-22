package repository

import (
	"context"

	"github.com/koba1108/go-mongodb/internals/domain/model"
)

type PlayerRepository interface {
	Create(context.Context, *model.Player) (*model.Player, error)
	Update(context.Context, *model.Player) (*model.Player, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.Player, error)
	FindAll(ctx context.Context) ([]*model.Player, error)
}
