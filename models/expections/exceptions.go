package exceptions

import (
	"fmt"
)

type UnprocessableEntity struct {
	Status  int
	Message string
}

func (e *UnprocessableEntity) Error() string {
	e.Status = 422
	return fmt.Sprint("Transaction error: ", e.Message)
}

type ServerError struct {
	Status  int
	Message string
}

func (e *ServerError) Error() string {
	e.Status = 500
	return fmt.Sprint("Internal Server error: ", e.Message)
}

type UserNotFound struct {
	Status  int
	Message string
}

func (e *UserNotFound) Error() string {
	e.Status = 404
	return fmt.Sprint("User not found error: ", e.Message)
}
