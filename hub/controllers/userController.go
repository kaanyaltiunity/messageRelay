package controllers

import (
	"messageHub/models"
	"messageHub/services"
	"net/http"

	"github.com/labstack/echo"
)

type UserController interface {
	Register(echo.Context) error
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) Register(ctx echo.Context) error {
	user, err := u.service.Register(ctx)
	if err != nil {
		ctx.Error(err)
		return err
	}
	registerUserDTO := models.RegisterUserDTO{
		UUID: user.UUID.String(),
	}
	return ctx.JSON(http.StatusOK, registerUserDTO)
}
