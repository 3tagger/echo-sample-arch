package main

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

func gracefulShutdown(e *echo.Echo, gracePeriod int) {
	e.Logger.Infof("initiating graceful shutdown with duration of %d second(s)", gracePeriod)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(gracePeriod)*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("cannot shutdown the echo server: %s", err)
	}

	e.Logger.Info("server has been stopped")
}
