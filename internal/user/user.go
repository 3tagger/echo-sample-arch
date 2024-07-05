package main

type User struct {
	Id    int64
	Email string
	Name  string
	About string
}

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	GetOneUserById() (*User, error)
	InsertOneUser(user User) (*User, error)
	DeleteOneUser(user User) error
	UpdateOneUserById(userId int64, user User) error
}
