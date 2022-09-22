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
	FindByIdWithVideos(ctx context.Context, id string, limit, offset int, sortKey string, orderBy model.OrderBy) ([]*model.PlayerWithVideos, error)
	FindAll(ctx context.Context) ([]*model.Player, error)
}
