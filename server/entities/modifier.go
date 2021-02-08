package entities

import "gorm.io/gorm"

//Modifier Entity
type Modifier struct {
	gorm.Model
	ProductID   uint
	Description string
	Price       float32
}
