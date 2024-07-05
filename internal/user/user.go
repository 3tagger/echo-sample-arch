package user

import "context"

type User struct {
	Id    int64
	Email string
	Name  string
	About string
}

type Repository interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetOneUserById(ctx context.Context, userId int64) (*User, error)
	InsertOneUser(ctx context.Context, user User) (*User, error)
	DeleteOneUser(ctx context.Context, userId int64) error
	UpdateOneUserById(ctx context.Context, userId int64, user User) error
}
