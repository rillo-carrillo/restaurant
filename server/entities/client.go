package entities

import "gorm.io/gorm"

//Client entity
type Client struct {
	gorm.Model
	RestaurantID   uint
	Name           string
	FirstLastname  string
	SecondLastname string
	Address        string
	Number         int
	CP             int
	RFC            string `gorm:"size:13"`
	Orders         []Order
}
