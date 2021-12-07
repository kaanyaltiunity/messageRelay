package controllers

import (
	"messageHub/models"
	"messageHub/services"
	"net/http"

	"github.com/labstack/echo"
)

type UserController interface {
	Register(echo.Context) error
	GetUsers(echo.Context) error
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
	cookie := new(http.Cookie)
	cookie.Name = "userid"
	cookie.Value = registerUserDTO.UUID
	ctx.SetCookie(cookie)
	return ctx.JSON(http.StatusOK, registerUserDTO)
}

func (u *userController) GetUsers(ctx echo.Context) error {
	// ownId, err := ctx.Cookie("userid")
	// if err != nil {
	// 	fmt.Println(err)
	// 	ctx.Error(err)
	// 	return err
	// }
	ownId := "test"
	users, err := u.service.GetUsers(ctx, ownId)
	if err != nil {
		ctx.Error(err)
		return err
	}
	var getUsersDTO []models.GetUserDTO
	for _, v := range users {
		getUsersDTO = append(getUsersDTO, models.GetUserDTO{UUID: v.UUID.String()})
	}
	return ctx.JSON(http.StatusOK, getUsersDTO)
}
