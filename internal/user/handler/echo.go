package handler

import (
	"net/http"

	"github.com/3tagger/echo-sample-arch/internal/user"
	userdto "github.com/3tagger/echo-sample-arch/internal/user/dto"
	"github.com/3tagger/echo-sample-arch/internal/util/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserEchoHandler struct {
	userUsecase user.Usecase
	validate    *validator.Validate
}

func NewUserEchoHandler(userUsecase user.Usecase, validate *validator.Validate) *UserEchoHandler {
	return &UserEchoHandler{
		userUsecase: userUsecase,
		validate:    validate,
	}
}

func (h *UserEchoHandler) GetOneUserById(c echo.Context) error {
	var req userdto.GetOneUserByIdRequest
	ctx := c.Request().Context()

	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.validate.StructCtx(ctx, &req)
	if err != nil {
		return err
	}

	u := req.GetOneUserByIdRequestToEntity()

	res, err := h.userUsecase.GetOneById(ctx, u.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.SimpleResponse(res))
}

func (h *UserEchoHandler) CreateOneUser(c echo.Context) error {
	var req userdto.CreateOneUserRequest
	ctx := c.Request().Context()

	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.validate.StructCtx(ctx, &req)
	if err != nil {
		return err
	}

	cu, err := h.userUsecase.InsertOne(ctx, user.User{
		Email: req.Email,
		Name:  req.Name,
		About: req.About,
	})
	if err != nil {
		return err
	}

	res := userdto.EntityToCreateOneUserResponse(cu)

	return c.JSON(http.StatusOK, dto.SimpleResponse(res))
}

func (h *UserEchoHandler) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.userUsecase.GetAll(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.SimpleResponse(res))
}
