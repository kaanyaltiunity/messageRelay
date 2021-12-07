package infrastructure

import (
	"fmt"
	"messageHub/models"
)

type MessageRepository struct {
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (m *MessageRepository) RelayMessage(message *models.Message) error {
	fmt.Println(message)
	return nil
}
