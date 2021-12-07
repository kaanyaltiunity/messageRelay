package controllers

import (
	"messageHub/models"
	"messageHub/services"

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
	registerUserDTO := models.RegisterUserDTO{}
	err := ctx.Bind(&registerUserDTO)
	if err != nil {
		ctx.Error(err)
		return err
	}

	err = u.service.Register(ctx, registerUserDTO)
	if err != nil {
		ctx.Error(err)
		return err
	}
	return ctx.String(200, "user registered")
}
