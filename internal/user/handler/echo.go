package handler

import (
	"net/http"

	"github.com/3tagger/echo-sample-arch/internal/user"
	userdto "github.com/3tagger/echo-sample-arch/internal/user/dto"
	"github.com/3tagger/echo-sample-arch/internal/util/dto"
	"github.com/labstack/echo/v4"
)

type UserEchoHandler struct {
	userUsecase user.Usecase
}

func NewUserEchoHandler(userUsecase user.Usecase) *UserEchoHandler {
	return &UserEchoHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserEchoHandler) GetOneUserById(c echo.Context) error {
	var req userdto.GetOneUserByIdRequest
	ctx := c.Request().Context()

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.SimpleMessageResponse())
	}

	res, err := h.userUsecase.GetOneById(ctx, req.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, dto.SimpleResponse(res))
}
