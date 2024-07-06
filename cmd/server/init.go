package main

import (
	"database/sql"

	userhandler "github.com/3tagger/echo-sample-arch/internal/user/handler"
	userrepository "github.com/3tagger/echo-sample-arch/internal/user/repository"
	userusecase "github.com/3tagger/echo-sample-arch/internal/user/usecase"
	sitehandler "github.com/3tagger/echo-sample-arch/internal/util/handler"
)

type Handlers struct {
	UserHandler *userhandler.UserEchoHandler
	SiteHandler *sitehandler.SiteEchoHandler
}

func initHandlers(db *sql.DB) *Handlers {
	userRepository := userrepository.NewUserRepositoryPostgreSQL(db)
	userUsecase := userusecase.NewUserUsecase(userRepository)
	userHandler := userhandler.NewUserEchoHandler(userUsecase)

	siteHandler := sitehandler.NewSiteEchoHandler()

	return &Handlers{
		SiteHandler: siteHandler,
		UserHandler: userHandler,
	}
}
