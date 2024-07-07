package main

import (
	"errors"
	"net/http"

	"github.com/3tagger/echo-sample-arch/internal/util/dto"
	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	var (
		httpError *echo.HTTPError
		msg       any
	)

	if errors.As(err, &httpError) {
		code = httpError.Code
		msg = httpError.Message
	} else {
		msg = "We encountered error, please try again later."
	}

	if txtMsg, ok := msg.(string); ok {
		c.JSON(code, dto.SimpleMessageResponse(txtMsg))
	} else {
		c.JSON(code, dto.SimpleResponse(msg, "Error"))
	}
}
