package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ServerHost = ""
	ServerPort = "8080"
)

func main() {
	// initializing echo server
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	srvAddr := fmt.Sprintf("%s:%s", ServerHost, ServerPort)

	e.Logger.Info("server running on %s", srvAddr)
	if err := e.Start(srvAddr); err != nil {
		e.Logger.Fatal("error running echo server: %s", err)
	}
}