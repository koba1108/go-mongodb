package handler

import (
	"github.com/koba1108/go-mongodb/internals/usecase"
	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	Find(c echo.Context) error
}

func NewArticleHandler(au usecase.ArticleUsecase) ArticleHandler {
	return &articleHandler{
		articleUsecase: au,
	}
}

type articleHandler struct {
	articleUsecase usecase.ArticleUsecase
}

func (ah *articleHandler) Find(c echo.Context) error {
	var req struct {
		PlayerID string `query:"playerID" binding:"required"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	articles, err := ah.articleUsecase.FindByPlayerID(c.Request().Context(), req.PlayerID)
	if err != nil {
		return err
	}
	return c.JSON(200, articles)
}
