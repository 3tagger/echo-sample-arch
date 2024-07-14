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

			var (
				infoErr  apperror.InformativeError
				validErr validator.ValidationErrors
				echoErr  *echo.HTTPError
			)

			resErr := &echo.HTTPError{
				Code:     http.StatusInternalServerError,
				Message:  "Our server is encountering an error, please try again later.",
				Internal: v.Error,
			}

			j["error"] = v.Error.Error()

			if errors.As(v.Error, &echoErr) {
				resErr = echoErr
			} else if errors.As(v.Error, &infoErr) {
				resErr.Code = http.StatusBadRequest
				resErr.Message = infoErr.Error()
			} else if errors.As(v.Error, &validErr) {
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

				resErr.Code = http.StatusBadRequest
				resErr.Message = vldErr
			}

			j["error_message"] = resErr.Message
			j["status"] = resErr.Code

			if resErr.Code/100 == 5 {
				e.Logger.Errorj(j)
			} else {
				e.Logger.Warnj(j)
			}

			return resErr
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
