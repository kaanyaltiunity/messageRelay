package services

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type UserService interface {
	Register(echo.Context) (*models.User, error)
	GetUsers(echo.Context, string) ([]*models.User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (u *userService) Register(ctx echo.Context) (*models.User, error) {
	user, err := models.NewUser()
	if err != nil {
		return nil, err
	}
	err = u.repository.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) GetUsers(ctx echo.Context, ownId string) ([]*models.User, error) {
	usersIds, err := u.repository.GetUsers(ctx, ownId)
	if err != nil {
		return nil, err
	}
	var users []*models.User

	for _, v := range usersIds {
		user, err := models.UserFromId(v)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
