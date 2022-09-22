package repository

import (
	"context"
	"github.com/koba1108/go-mongodb/internals/domain/model"
)

type ArticleRepository interface {
	FindByPlayerID(ctx context.Context, playerID string) ([]*model.Article, error)
}
