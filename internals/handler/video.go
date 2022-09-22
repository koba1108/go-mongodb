package handler

import (
	"github.com/koba1108/go-mongodb/internals/usecase"
	"github.com/labstack/echo/v4"
)

type VideoHandler interface {
	List(c echo.Context) error
	GetByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func NewVideoHandler(vu usecase.VideoUsecase) VideoHandler {
	return &videoHandler{videoUsecase: vu}
}

type videoHandler struct {
	videoUsecase usecase.VideoUsecase
}

func (vh *videoHandler) List(c echo.Context) error {
	return nil
}

func (vh *videoHandler) GetByID(c echo.Context) error {
	return nil
}

func (vh *videoHandler) Create(c echo.Context) error {
	return nil
}

func (vh *videoHandler) Update(c echo.Context) error {
	return nil
}

func (vh *videoHandler) Delete(c echo.Context) error {
	return nil
}
