package models

import "github.com/google/uuid"

type Message struct {
	UUID      uuid.UUID
	Receivers []uuid.UUID
	Text      string
}

type RelayMessageDTO struct {
	Receivers []uuid.UUID `json:"receivers"`
	Text      string      `json:"text"`
}

func NewMessage(receivers []uuid.UUID, text string) (*Message, error) {
	return &Message{
		UUID:      uuid.New(),
		Receivers: receivers,
		Text:      text,
	}, nil
}
