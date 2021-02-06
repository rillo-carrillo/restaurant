package orm

import "gorm.io/gorm"

//User model for ORM
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//UserCreationResponse when created return id
type UserCreationResponse struct {
	ID       uint
	Username string
}
