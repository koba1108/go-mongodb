package usecase

import (
	"context"
	"github.com/koba1108/go-mongodb/internals/domain/model"
	"github.com/koba1108/go-mongodb/internals/domain/repository"
)

type ArticleUsecase interface {
	FindByPlayerID(ctx context.Context, playerID string) ([]*model.Article, error)
}

type articleUsecase struct {
	articleRepo repository.ArticleRepository
}

func NewArticleUsecase(articleRepo repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{articleRepo: articleRepo}
}

func (au *articleUsecase) FindByPlayerID(ctx context.Context, playerID string) ([]*model.Article, error) {
	return au.articleRepo.FindByPlayerID(ctx, playerID)
}
