package errors

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

//DBErrorResponse struct
type DBErrorResponse struct {
	Field   string
	Message string
}

//ServerError strunc message
type ServerError struct {
	Message string
}

const (
	//
	duplicated = "23505"
)

//ValidateUserErrors with DB
func ValidateUserErrors(err error) (res DBErrorResponse) {
	val := err.Error()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.Field = "username"
		res.Message = "username not found."
	} else if strings.Contains(val, duplicated) {
		res.Field = "username"
		res.Message = "username already exists."
	}
	return res
}
