package routes

import (
	"messageHub/controllers"
	"messageHub/router"
)

type UserRoutes struct {
	controller controllers.UserController
}

func NewUserRoutes(controller controllers.UserController) *UserRoutes {
	return &UserRoutes{
		controller: controller,
	}
}

func (u *UserRoutes) BindControllers(router router.Router) {
	router.POST("/register", u.controller.Register)
	router.GET("/users", u.controller.GetUsers)
}
