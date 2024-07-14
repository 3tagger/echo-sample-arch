package dto

import (
	"strconv"

	"github.com/3tagger/echo-sample-arch/internal/user"
)

type GetOneUserByIdRequest struct {
	Id string `param:"user_id" validate:"required,numeric" example:"1"` // ID of the user
}

func (req *GetOneUserByIdRequest) GetOneUserByIdRequestToEntity() user.User {
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	return user.User{Id: id}
}

type UserResponseEntry struct {
	Id    int64  `json:"id" example:"1"`                        // ID of the user
	Name  string `json:"name" example:"Bob"`                    // Name of the user
	About string `json:"about" example:"He is one of the user"` // 	Description about the user
}

// GetOneUserByIdResponse model info
// @Description This response contains user's information with id, name and about
type GetOneUserByIdResponse struct {
	User UserResponseEntry `json:"user"`
}

func EntityToGetOneUserByIdResponse(user *user.User) GetOneUserByIdResponse {
	res := GetOneUserByIdResponse{}

	if user != nil {
		res.User = UserResponseEntry{
			Id:    user.Id,
			Name:  user.Name,
			About: user.About,
		}
	}
	return res
}

// GetOneUserByIdResponse model info
// @Description This request will create a user by providing user's data
type CreateOneUserRequest struct {
	Name  string `json:"name" validate:"required" example:"Bob"`                    // Name of the user
	Email string `json:"email" validate:"required,email" example:"bob@example.com"` // Email of the user
	About string `json:"about" example:"He is one of the user"`                     // Description about the user
}

// GetOneUserByIdResponse model info
// @Description This response will contain the created user data
type CreateOneUserResponse struct {
	Id    int64  `json:"id" example:"1"`                        // ID of the user
	Name  string `json:"name" example:"Bob"`                    // Name of the user
	Email string `json:"email" example:"bob@example.com"`       // Email of the user
	About string `json:"about" example:"He is one of the user"` // Description about the user
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

// GetAllUsersResponse model info
// @Description This response contains all users information with id, name and about
type GetAllUsersResponse struct {
	Users []UserResponseEntry `json:"users"`
}

func EntityToGetAllUsersResponse(users []*user.User) GetAllUsersResponse {
	res := GetAllUsersResponse{
		Users: []UserResponseEntry{},
	}
	for _, u := range users {
		e := UserResponseEntry{
			Id:    u.Id,
			Name:  u.Name,
			About: u.About,
		}

		res.Users = append(res.Users, e)
	}

	return res
}
