package services

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type UserService interface {
	Register(echo.Context, models.RegisterUserDTO) error
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (m *userService) Register(ctx echo.Context, registerUserDTO models.RegisterUserDTO) error {
	user, err := models.NewUser(registerUserDTO.UUID)
	if err != nil {
		return err
	}
	return m.repository.Register(ctx, user)
}
