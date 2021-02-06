package utils

import "golang.org/x/crypto/bcrypt"

//GeneratePassword handle the Hashing of password
func GeneratePassword(pass string) (string, error) {
	bPass := []byte(pass)
	hashed, err := bcrypt.GenerateFromPassword(bPass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil

}
