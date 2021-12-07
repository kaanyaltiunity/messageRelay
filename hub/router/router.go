package router

import "github.com/labstack/echo"

type Router struct {
	*echo.Echo
}

func NewRouter() Router {
	return Router{
		echo.New(),
	}
}
