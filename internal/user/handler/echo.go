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

// GetOneUserById gets one user by id
//
//	@Id				GetOneUserById
//	@Summary		Get One User By ID
//	@Description	Retrieving a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	userdto.GetOneUserByIdResponse
//	@Failure		400		{object}	dto.HttpResponse
//	@Failure		404		{object}	dto.HttpResponse
//	@Failure		500		{object}	dto.HttpResponse
//	@Router			/users/{user_id} [get]
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

	getU, err := h.userUsecase.GetOneById(ctx, u.Id)
	if err != nil {
		return err
	}

	res := userdto.EntityToGetOneUserByIdResponse(getU)

	return c.JSON(http.StatusOK, dto.SimpleResponse(res))
}

// CreateOneUser create one user
//
//	@Id				CreateOneUser
//	@Summary		Create a user based on provided data
//	@Description	Create a user based on provided data. When the user is created, the response will return the newly generated user ID.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		userdto.CreateOneUserRequest	true	"The request should follow the CreateOneUserRequest model"
//	@Success		200		{object}	userdto.CreateOneUserResponse
//	@Failure		400		{object}	dto.HttpResponse
//	@Failure		500		{object}	dto.HttpResponse
//	@Router			/users [post]
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

// GetAllUsers gets all users
//
//	@Id				GetAllUsers
//	@Summary		Get all users
//	@Description	Get all users data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	userdto.GetAllUsersResponse
//	@Failure		500	{object}	dto.HttpResponse
//	@Router			/users [get]
func (h *UserEchoHandler) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	lu, err := h.userUsecase.GetAll(ctx)
	if err != nil {
		return err
	}

	res := userdto.EntityToGetAllUsersResponse(lu)

	return c.JSON(http.StatusOK, dto.SimpleResponse(res))
}
