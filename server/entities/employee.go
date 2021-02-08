package entities

import "gorm.io/gorm"

//Employee entity
type Employee struct {
	gorm.Model
	Name         string
	RestaurantID uint
	RoleID       int
	Order        []Order
}
