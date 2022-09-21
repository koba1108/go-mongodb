package handler

import (
	"log"
	"net/http"

	"github.com/koba1108/go-mongodb/internals/usecase"
	"github.com/labstack/echo"
)

type UserHandler interface {
	List(c echo.Context) error
	GetByID(c echo.Context) error
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uu}
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func (uh *userHandler) List(c echo.Context) error {
	log.Println("List")
	users, err := uh.userUsecase.FindAll(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (uh *userHandler) GetByID(c echo.Context) error {
	user, err := uh.userUsecase.FindByID(c.Request().Context(), c.Param("userId"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
