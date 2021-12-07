package services

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type UserRepository interface {
	Register(echo.Context, *models.User) error
	GetUsers(echo.Context, string) ([]string, error)
}
