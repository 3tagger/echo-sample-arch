package main

import (
	"errors"
	"net/http"

	"github.com/3tagger/echo-sample-arch/internal/apperror"
	"github.com/3tagger/echo-sample-arch/internal/util/dto"
	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	var (
		infoErr apperror.InformativeError
		msg     string
	)

	if errors.As(err, &infoErr) {
		code = http.StatusBadRequest
		msg = infoErr.Error()
	} else {
		msg = "Our server is encountering an error, please try again later."
	}

	c.JSON(code, dto.SimpleMessageResponse(msg))
}
