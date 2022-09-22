package database

import (
	"context"
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type articleRepository struct {
	db        *mongo.Database
	playerCol *mongo.Collection
	videoCol  *mongo.Collection
}

func NewArticleRepository(db *mongo.Database) *articleRepository {
	return &articleRepository{
		db: db,
	}
}

func (ar *articleRepository) FindByPlayerID(ctx context.Context, playerID string) ([]*model.Article, error) {
	return nil, nil
}
