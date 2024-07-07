package apperror

import "fmt"

type InformativeError struct {
	message string
}

func (e InformativeError) Error() string {
	return e.message
}

func ErrUserWithIdNotFound(userId int64) InformativeError {
	return InformativeError{
		message: fmt.Sprintf("user with id %d not found", userId),
	}
}
