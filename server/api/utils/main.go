package utils

import (
	"errors"
)

//GetUserID check in session for the key and return userId, error
func GetUserID(u interface{}) (userID uint, err error) {
	switch v := u.(type) {
	case uint:
		return v, nil
	default:
		return 0, errors.New("UserID not found")
	}
}
