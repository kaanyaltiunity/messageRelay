package routes

import (
	"messageHub/controllers"
	"messageHub/router"
)

type MessageRoutes struct {
	controller controllers.MessageController
}

func NewMessageRoutes(controller controllers.MessageController) *MessageRoutes {
	return &MessageRoutes{
		controller: controller,
	}
}

func (m *MessageRoutes) BindControllers(router router.Router) {
	router.POST("/messages", m.controller.RelayMessage)
}
