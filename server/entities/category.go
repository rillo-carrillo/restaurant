package entities

import "gorm.io/gorm"

//Category entity
type Category struct {
	gorm.Model
	RestaurantID uint
	Name         string
	ParentID     *uint
	Parent       *Category
	Products     []Product
}
