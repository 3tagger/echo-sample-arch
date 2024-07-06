package main

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, handlers *Handlers) {
	siteHandler := handlers.SiteHandler
	userHandler := handlers.UserHandler

	e.GET("/", siteHandler.Home)
	e.GET("/site/demo-cancellation", siteHandler.DemoContextCancellation)
	e.GET("/users", userHandler.GetAllUsers)
	e.GET("/users/:user_id", userHandler.GetOneUserById)
	e.POST("/users", userHandler.CreateOneUser)
}
