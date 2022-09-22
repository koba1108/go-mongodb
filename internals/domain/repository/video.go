package repository

import (
	"context"

	"github.com/koba1108/go-mongodb/internals/domain/model"
)

type VideoRepository interface {
	Create(context.Context, *model.Video) (*model.Video, error)
	Update(context.Context, *model.Video) (*model.Video, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.Video, error)
	FindAll(ctx context.Context, keyword string) ([]*model.Video, error)
}
