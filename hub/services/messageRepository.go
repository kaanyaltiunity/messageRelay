package services

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type MessageRepository interface {
	RelayMessage(echo.Context, *models.Message) error
}
