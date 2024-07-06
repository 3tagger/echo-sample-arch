package main

import (
	"fmt"
	"log"

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

	e.Logger.Infof("server running on %s", srvAddr)
	if err := e.Start(srvAddr); err != nil {
		e.Logger.Fatalf("error running echo server: %s", err)
	}

	e.Logger.Info("server has been stopped")
}
