package app

import (
	"log"
	"messageHub/controllers"
	"messageHub/infrastructure"
	"messageHub/router"
	"messageHub/routes"
	"messageHub/services"
	"net/http"
)

type App struct {
	router router.Router
}

func NewApp() *App {
	return &App{
		router: router.NewRouter(),
	}
}

func (a *App) Start() {
	messageRepository := infrastructure.NewMessageRepository()
	messageService := services.NewMessageService(messageRepository)
	messageController := controllers.NewMessageController(messageService)
	messageRoutes := routes.NewMessageRoutes(messageController)
	messageRoutes.BindControllers(a.router)

	if err := a.router.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
