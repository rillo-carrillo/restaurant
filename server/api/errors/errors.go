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
	relation   = "42P01"
)

//ValidateUserErrors with DB
func ValidateUserErrors(err error) (res DBErrorResponse) {
	val := err.Error()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.Field = "username"
		res.Message = "username not found."
	} else if strings.Contains(val, relation) {
		res.Field = "error"
		res.Message = "something went wrong with the DB. please contact administrator."
	} else if strings.Contains(val, duplicated) {
		res.Field = "username"
		res.Message = "username already exists."
	}
	return res
}
