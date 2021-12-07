package infrastructure

import (
	"fmt"
	"messageHub/models"

	"github.com/labstack/echo"
)

type MessageRepository struct {
	publisher Publisher
}

func NewMessageRepository(publisher Publisher) *MessageRepository {
	return &MessageRepository{
		publisher: publisher,
	}
}

func (m *MessageRepository) RelayMessage(ctx echo.Context, message *models.Message) error {
	fmt.Println(message)
	for _, v := range message.Receivers {
		m.publisher.Publish(ctx, v.String(), message.Text)
	}
	return nil
}
