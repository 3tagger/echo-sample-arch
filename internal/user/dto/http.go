package dto

import "github.com/3tagger/echo-sample-arch/internal/user"

type GetOneUserByIdRequest struct {
	Id int64 `param:"user_id"`
}

type CreateOneUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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
