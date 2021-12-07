package models

import "github.com/google/uuid"

type Message struct {
	UUID      uuid.UUID
	Receivers []uuid.UUID
	Text      string
}

type RelayMessageDTO struct {
	Receivers []string `json:"receivers"`
	Text      string   `json:"text"`
}

func NewMessage(receivers []string, text string) (*Message, error) {
	var uuids []uuid.UUID

	for _, v := range receivers {
		uuid, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, uuid)
	}

	return &Message{
		UUID:      uuid.New(),
		Receivers: uuids,
		Text:      text,
	}, nil
}
