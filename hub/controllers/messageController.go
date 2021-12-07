package controllers

import (
	"messageHub/models"
	"messageHub/services"

	"github.com/labstack/echo"
)

type MessageController interface {
	RelayMessage(echo.Context) error
}

type messageController struct {
	service services.MessageService
}

func NewMessageController(service services.MessageService) MessageController {
	return &messageController{
		service: service,
	}
}

func (m *messageController) RelayMessage(ctx echo.Context) error {
	messageDTO := models.RelayMessageDTO{}
	err := ctx.Bind(messageDTO)
	if err != nil {
		panic(err)
	}
	return m.service.RelayMessage(messageDTO)
}
