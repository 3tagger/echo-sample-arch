package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/3tagger/echo-sample-arch/docs"

	"github.com/3tagger/echo-sample-arch/internal/config"
	"github.com/3tagger/echo-sample-arch/internal/database"
	"github.com/go-playground/validator/v10"
)

//	@title			echo-sample-arch
//	@version		0.1.0
//	@description	Sample of simple web server built using Echo framework. You can visit the GitHub repository at https://github.com/3tagger/echo-sample-arch

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

	validate := validator.New(validator.WithRequiredStructEnabled())

	// initializing echo server
	e := initEcho()

	RegisterRoutes(e, initHandlers(db, validate))

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

	gracefulShutdown(e, cfg.Server.GracePeriod)
}
