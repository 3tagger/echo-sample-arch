package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	echolog "github.com/labstack/gommon/log"

	"github.com/3tagger/echo-sample-arch/internal/config"
	"github.com/3tagger/echo-sample-arch/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to load config: %v\n", err)
	}

	db, err := database.InitPostgreSQL(cfg.Database.Primary)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	// initializing echo server
	e := echo.New()
	e.Logger.SetLevel(echolog.INFO)

	handlers := initHandlers(db)

	RegisterRoutes(e, handlers)

	srvCfg := cfg.Server
	srvAddr := fmt.Sprintf("%s:%s", srvCfg.Host, srvCfg.Post)

	sigCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// start server
	go func() {
		e.Logger.Infof("server running on %s", srvAddr)
		if err := e.Start(srvAddr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("error running echo server: %s", err)
		}
	}()

	// waiting for os signals, such as SIGINT
	<-sigCtx.Done()

	gracePeriod := cfg.Server.GracePeriod

	e.Logger.Infof("initiating graceful shutdown with duration of %d second(s)", gracePeriod)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(gracePeriod)*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("cannot shutdown the echo server: %s", err)
	}

	e.Logger.Info("server has been stopped")
}
