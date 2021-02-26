package orm

import "gorm.io/gorm"

//User model for ORM
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//UserResponse when created return id
type UserResponse struct {
	ID       uint
	Username string
}

//StatusResponse only reply with a true or false
type StatusResponse struct {
	Status bool
}
