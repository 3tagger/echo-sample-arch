package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(e *echo.Echo, handlers *Handlers) {
	siteHandler := handlers.SiteHandler
	userHandler := handlers.UserHandler

	e.GET("/", siteHandler.Home)
	e.GET("/site/demo-cancellation", siteHandler.DemoContextCancellation)

	usersgroup := e.Group("/users")
	{
		usersgroup.GET("", userHandler.GetAllUsers)
		usersgroup.GET("/:user_id", userHandler.GetOneUserById)
		usersgroup.POST("", userHandler.CreateOneUser)
	}

	swaggroup := e.Group("/swagger")
	{
		swaggroup.GET("", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
		})
		swaggroup.GET("/*", echoSwagger.WrapHandler)
	}
}
