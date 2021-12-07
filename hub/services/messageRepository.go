package services

import "messageHub/models"

type MessageRepository interface {
	RelayMessage(*models.Message) error
}
