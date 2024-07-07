package main

import (
	"database/sql"

	userhandler "github.com/3tagger/echo-sample-arch/internal/user/handler"
	userrepository "github.com/3tagger/echo-sample-arch/internal/user/repository"
	userusecase "github.com/3tagger/echo-sample-arch/internal/user/usecase"
	sitehandler "github.com/3tagger/echo-sample-arch/internal/util/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echolog "github.com/labstack/gommon/log"
)

type Handlers struct {
	UserHandler *userhandler.UserEchoHandler
	SiteHandler *sitehandler.SiteEchoHandler
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.Logger.SetHeader(`{"time":"${time_rfc3339_nano}","level":"${level}"}`)
	e.Logger.SetLevel(echolog.INFO)
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogHost:      true,
		LogRequestID: true,
		LogMethod:    true,
		HandleError:  true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			j := echolog.JSON{
				"request_id": v.RequestID,
				"method":     v.Method,
				"host":       v.Host,
				"uri":        v.URI,
				"status":     v.Status,
			}

			// give different log level for 5xx status
			if v.Status/100 == 5 {
				e.Logger.Errorj(j)
			} else {
				e.Logger.Infoj(j)
			}

			return v.Error
		},
	}))

	return e
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
