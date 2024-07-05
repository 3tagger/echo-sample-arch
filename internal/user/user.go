package user

import "context"

type User struct {
	Id    int64
	Email string
	Name  string
	About string
}

type UserFaker struct {
	Id    int64  `faker:"-"`
	Email string `faker:"email"`
	Name  string `faker:"name"`
	About string `faker:"paragraph"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetOneById(ctx context.Context, userId int64) (*User, error)
	InsertOne(ctx context.Context, user User) (*User, error)
	InsertMany(ctx context.Context, user []User) error
	DeleteOne(ctx context.Context, userId int64) error
	UpdateOneById(ctx context.Context, userId int64, user User) error
}
