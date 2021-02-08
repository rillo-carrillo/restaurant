package entities

import "gorm.io/gorm"

//Product contains the description of the products to offer
type Product struct {
	gorm.Model
	CategoryID uint
	Name       string
	Price      float32
	Comments   string
	Modifiers  []Modifier `gorm:"constraint:,OnDelete:CASCADE;"`
}
