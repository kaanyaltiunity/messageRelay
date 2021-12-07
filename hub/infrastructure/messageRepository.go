package infrastructure

import (
	"fmt"
	"messageHub/models"
)

type MessageRepository struct {
	publisher Publisher
}

func NewMessageRepository(publisher Publisher) *MessageRepository {
	return &MessageRepository{
		publisher: publisher,
	}
}

func (m *MessageRepository) RelayMessage(message *models.Message) error {
	fmt.Println(message)
	return nil
}
