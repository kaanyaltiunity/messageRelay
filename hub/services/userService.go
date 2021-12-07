package services

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type UserService interface {
	Register(echo.Context) (*models.User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (m *userService) Register(ctx echo.Context) (*models.User, error) {
	user, err := models.NewUser()
	if err != nil {
		return nil, err
	}
	err = m.repository.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
