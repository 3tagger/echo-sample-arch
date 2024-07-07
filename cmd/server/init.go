package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/3tagger/echo-sample-arch/internal/apperror"
	userhandler "github.com/3tagger/echo-sample-arch/internal/user/handler"
	userrepository "github.com/3tagger/echo-sample-arch/internal/user/repository"
	userusecase "github.com/3tagger/echo-sample-arch/internal/user/usecase"
	sitehandler "github.com/3tagger/echo-sample-arch/internal/util/handler"
	"github.com/go-playground/validator/v10"
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
		LogError:     true,
		LogMethod:    true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			j := echolog.JSON{
				"request_id": v.RequestID,
				"method":     v.Method,
				"host":       v.Host,
				"uri":        v.URI,
				"status":     v.Status,
			}

			if v.Error == nil {
				e.Logger.Infoj(j)
				return nil
			}

			code := http.StatusInternalServerError

			var (
				infoErr  apperror.InformativeError
				validErr validator.ValidationErrors
				msg      any
			)

			j["error"] = v.Error.Error()

			err := v.Error
			if errors.As(err, &infoErr) {
				code = http.StatusBadRequest
				msg = infoErr.Error()
			} else if errors.As(err, &validErr) {
				code = http.StatusBadRequest
				vldErr := map[string]string{}
				for _, err := range validErr {
					switch err.Tag() {
					case "required":
						vldErr[err.Field()] = fmt.Sprintf("%s is required",
							err.Field())
					case "email":
						vldErr[err.Field()] = fmt.Sprintf("%s is not valid email",
							err.Field())
					case "gte":
						vldErr[err.Field()] = fmt.Sprintf("%s value must be greater than %s",
							err.Field(), err.Param())
					case "lte":
						vldErr[err.Field()] = fmt.Sprintf("%s value must be lower than %s",
							err.Field(), err.Param())
					case "numeric":
						vldErr[err.Field()] = fmt.Sprintf("%s must be numeric",
							err.Field())
					default:
						vldErr[err.Field()] = fmt.Sprintf("%s is not valid", err.Field())
					}
				}

				msg = vldErr
			} else {
				msg = "Our server is encountering an error, please try again later."
			}

			j["error_message"] = msg
			j["status"] = code

			if code/100 == 5 {
				e.Logger.Errorj(j)
			} else {
				e.Logger.Warnj(j)
			}

			return &echo.HTTPError{
				Code:     code,
				Message:  msg,
				Internal: err,
			}
		},
	}))

	return e
}

func initHandlers(db *sql.DB, validate *validator.Validate) *Handlers {
	userRepository := userrepository.NewUserRepositoryPostgreSQL(db)
	userUsecase := userusecase.NewUserUsecase(userRepository)
	userHandler := userhandler.NewUserEchoHandler(userUsecase, validate)

	siteHandler := sitehandler.NewSiteEchoHandler()

	return &Handlers{
		SiteHandler: siteHandler,
		UserHandler: userHandler,
	}
}
