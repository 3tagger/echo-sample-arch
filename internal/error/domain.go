package error

import "fmt"

type InformativeError struct {
	Message string
}

func (e InformativeError) Error() string {
	return e.Message
}

func ErrUserWithIdNotFound(userId int64) InformativeError {
	return InformativeError{
		Message: fmt.Sprintf("user with id %d not found", userId),
	}
}
