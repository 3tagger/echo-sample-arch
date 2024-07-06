package handler

import (
	"net/http"
	"time"

	"github.com/3tagger/echo-sample-arch/internal/util/dto"
	"github.com/labstack/echo/v4"
)

type SiteEchoHandler struct {
}

func NewSiteEchoHandler() *SiteEchoHandler {
	return &SiteEchoHandler{}
}

func (h *SiteEchoHandler) Home(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.SimpleMessageResponse("hello world"))
}

func (h *SiteEchoHandler) DemoContextCancellation(c echo.Context) error {
	ctx := c.Request().Context()

	// waiting 5 seconds to success, or request cancelled
	ticker := time.NewTicker(5 * time.Second)

	select {
	case <-ctx.Done():
		c.Logger().Infof("request cancelled: %s", ctx.Err())
		return nil
	case <-ticker.C:
	}

	c.Logger().Info("have waited 5 seconds, continuing...")

	return c.JSON(http.StatusOK, dto.SimpleMessageResponse())
}
