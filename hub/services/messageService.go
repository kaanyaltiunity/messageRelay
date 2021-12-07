package services

import (
	"messageHub/models"

	"github.com/labstack/echo"
)

type MessageService interface {
	RelayMessage(echo.Context, models.RelayMessageDTO) error
}

type messageService struct {
	repository MessageRepository
}

func NewMessageService(repository MessageRepository) MessageService {
	return &messageService{
		repository: repository,
	}
}

func (m *messageService) RelayMessage(ctx echo.Context, messageDTO models.RelayMessageDTO) error {
	message, err := models.NewMessage(messageDTO.Receivers, messageDTO.Text)
	if err != nil {
		return err
	}
	return m.repository.RelayMessage(ctx, message)
}
