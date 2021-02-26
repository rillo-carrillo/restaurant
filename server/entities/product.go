package entities

import "gorm.io/gorm"

//Product contains the description of the products to offer
type Product struct {
	gorm.Model
	CategoryID    uint
	Name          string
	OrderProducts []OrderProduct
	Modifiers     []Modifier `gorm:"constraint:,OnDelete:CASCADE;"`
}
