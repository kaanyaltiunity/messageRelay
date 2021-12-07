package services

import "messageHub/models"

type MessageService interface {
	RelayMessage(models.RelayMessageDTO) error
}

type messageService struct {
	repository MessageRepository
}

func NewMessageService(repository MessageRepository) MessageService {
	return &messageService{
		repository: repository,
	}
}

func (m *messageService) RelayMessage(messageDTO models.RelayMessageDTO) error {
	message, err := models.NewMessage(messageDTO.Receivers, messageDTO.Text)
	if err != nil {
		return err
	}
	return m.repository.RelayMessage(message)
}
