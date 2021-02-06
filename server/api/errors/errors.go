package errors

import (
	"strings"
)

//DBErrorResponse struct
type DBErrorResponse struct {
	Field   string
	Message string
}

const (
	//
	duplicated = "23505"
)

//ValidateUserErrors with DB
func ValidateUserErrors(err error) (res DBErrorResponse) {
	val := err.Error()
	if strings.Contains(val, duplicated) {
		res.Field = "username"
		res.Message = "username already exists."
	}
	return res
}
