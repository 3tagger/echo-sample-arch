package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/3tagger/echo-sample-arch/internal/config"
	"github.com/3tagger/echo-sample-arch/internal/database"
	userhandler "github.com/3tagger/echo-sample-arch/internal/user/handler"
	userrepository "github.com/3tagger/echo-sample-arch/internal/user/repository"
	userusecase "github.com/3tagger/echo-sample-arch/internal/user/usecase"
	"github.com/3tagger/echo-sample-arch/internal/util/dto"
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

	userRepository := userrepository.NewUserRepositoryPostgreSQL(db)
	userUsecase := userusecase.NewUserUsecase(userRepository)
	userhandler := userhandler.NewUserEchoHandler(userUsecase)

	e.GET("/users/:user_id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.SimpleMessageResponse("hello world"))
	})

	e.GET("/users/:user_id", userhandler.GetOneUserById)

	srvCfg := cfg.Server
	srvAddr := fmt.Sprintf("%s:%s", srvCfg.Host, srvCfg.Post)

	e.Logger.Info("server running on %s", srvAddr)
	if err := e.Start(srvAddr); err != nil {
		e.Logger.Fatal("error running echo server: %s", err)
	}

	e.Logger.Info("server has been stopped")
}
