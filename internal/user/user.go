package user

import "context"

type User struct {
	Id    int64
	Email string
	Name  string
	About string
}

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetOneById(ctx context.Context, userId int64) (*User, error)
	InsertOne(ctx context.Context, user User) (*User, error)
	InsertMany(ctx context.Context, users []User) error
	DeleteOne(ctx context.Context, userId int64) error
	UpdateOneById(ctx context.Context, userId int64, user User) error
}

type Usecase interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetOneById(ctx context.Context, userId int64) (*User, error)
	InsertOne(ctx context.Context, user User) (*User, error)
	InsertMany(ctx context.Context, users []User) error
	DeleteOne(ctx context.Context, userId int64) error
	UpdateOneById(ctx context.Context, userId int64, user User) error
}
