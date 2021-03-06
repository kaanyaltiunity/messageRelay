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
	messagePublisher := infrastructure.NewPublisher()
	messageRepository := infrastructure.NewMessageRepository(messagePublisher)
	messageService := services.NewMessageService(messageRepository)
	messageController := controllers.NewMessageController(messageService)
	messageRoutes := routes.NewMessageRoutes(messageController)
	messageRoutes.BindControllers(a.router)

	cache := infrastructure.NewCache()
	userRepository := infrastructure.NewUserRepository(cache)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	userRoutes := routes.NewUserRoutes(userController)
	userRoutes.BindControllers(a.router)

	if err := a.router.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
