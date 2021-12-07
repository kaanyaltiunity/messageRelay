package services

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type UserRepository interface {
	Register(echo.Context, *models.User) error
}
