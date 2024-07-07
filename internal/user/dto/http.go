package dto

import (
	"strconv"

	"github.com/3tagger/echo-sample-arch/internal/user"
)

type GetOneUserByIdRequest struct {
	Id string `param:"user_id" validate:"required,numeric"`
}

func (req *GetOneUserByIdRequest) GetOneUserByIdRequestToEntity() user.User {
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	return user.User{Id: id}
}

type CreateOneUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	About string `json:"about"`
}

type CreateOneUserResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	About string `json:"about"`
}

func EntityToCreateOneUserResponse(user *user.User) CreateOneUserResponse {
	res := CreateOneUserResponse{}
	if user != nil {
		res.Id = user.Id
		res.Name = user.Name
		res.Email = user.Email
		res.About = user.About
	}
	return res
}
