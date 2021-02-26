package entities

//Role Entity
type Role struct {
	Code     int `gorm:"primaryKey;autoIncrement:false"`
	Name     string
	Employee []Employee
}
