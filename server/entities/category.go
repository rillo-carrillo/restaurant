package entities

import "gorm.io/gorm"

//Category entity
type Category struct {
	gorm.Model
	RestaurantID uint
	Name         string
	Products     []Product
}
