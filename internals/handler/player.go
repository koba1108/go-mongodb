package handler

import (
	"net/http"

	"github.com/koba1108/go-mongodb/internals/usecase"
	"github.com/labstack/echo/v4"
)

type PlayerHandler interface {
	List(c echo.Context) error
	GetByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func NewPlayerHandler(pu usecase.PlayerUsecase) PlayerHandler {
	return &playerHandler{playerUsecase: pu}
}

type playerHandler struct {
	playerUsecase usecase.PlayerUsecase
}

func (ph *playerHandler) List(c echo.Context) error {
	players, err := ph.playerUsecase.FindAll(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, players)
}

func (ph *playerHandler) GetByID(c echo.Context) error {
	var req struct {
		ID         string `param:"id" binding:"required"`
		WithVideos bool   `query:"withVideos"`
		Limit      int    `query:"limit"`
		Offset     int    `query:"offset"`
		SortKey    string `query:"sortKey"`
		IsAsc      bool   `query:"isAsc"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.WithVideos {
		if req.Limit == 0 {
			req.Limit = 10
		}
		if req.Offset == 0 {
			req.Offset = 0
		}
		if req.SortKey == "" {
			req.SortKey = "id"
		}
		res, err := ph.playerUsecase.FindByIdWithVideos(c.Request().Context(), req.ID, req.Limit, req.Offset, req.SortKey, req.IsAsc)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}

	player, err := ph.playerUsecase.FindByID(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, player)
}

func (ph *playerHandler) Create(c echo.Context) error {
	var req struct {
		Name            string `json:"name" binding:"required"`
		OfficialSiteUrl string `json:"officialSiteUrl" binding:"required"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	player, err := ph.playerUsecase.Create(c.Request().Context(), req.Name, req.OfficialSiteUrl)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, player)
}

func (ph *playerHandler) Update(c echo.Context) error {
	var req struct {
		ID              string  `param:"id" binding:"required"`
		Name            *string `json:"name"`
		OfficialSiteUrl *string `json:"officialSiteUrl"`
	}
	player, err := ph.playerUsecase.Update(c.Request().Context(), req.ID, req.Name, req.OfficialSiteUrl)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, player)
}

func (ph *playerHandler) Delete(c echo.Context) error {
	if err := ph.playerUsecase.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
