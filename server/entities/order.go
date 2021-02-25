package entities

import (
	"gorm.io/gorm"
)

//STATUS INICIO, PREPARACION, REPARTO,LISTO

//Order contains the products, of the Order as well Employee ID
type Order struct {
	gorm.Model
	EmployeeID    uint
	ClientID      uint
	Status        int `gorm:"default:1;check:status IN (1,2,3,4)"`
	Total         float32
	OrderProducts []OrderProduct
	// Products   []Product `gorm:"many2many:order_products;"`
}

//BeforeCreate Hook on new Orders set status as 1.
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ClientID == 0 {
		o.ClientID = 1
	}
	o.Status = 1
	return nil
}

//OrderProduct define the join table
type OrderProduct struct {
	//`gorm:"primaryKey;autoIncrement:false"`
	ID           uint `gorm:"primaryKey"`
	OrderID      uint
	ProductID    uint
	Comment      string
	OrderDetails []OrderDetail
	// Modifiers []Modifier `gorm:"many2many:order_details;"`
}

//OrderDetail define the join table between OrderProduct and it's modifiers
type OrderDetail struct {
	ID             uint `gorm:"primaryKey"`
	OrderProductID uint
	ModifierID     uint
}

// //BeforeSave Hook to update the total of Order
// func (op *OrderProduct) BeforeSave(tx *gorm.DB) (err error) {
// 	var o Order
// 	if err := tx.First(&o, "id=?", op.OrderID); err.Error != nil {
// 		return err.Error
// 	}
// 	var p Modifier
// 	if err := tx.First(&p, "id=?", op.ModifierID); err.Error != nil {
// 		return err.Error
// 	}
// 	tx.Model(&o).Update("total", o.Total+p.Price)
// 	return nil
// }

// //BeforeDelete Hook to update the total of Order
// func (op *OrderProduct) BeforeDelete(tx *gorm.DB) (err error) {
// 	var o Order
// 	if err := tx.First(&o, "id=?", op.OrderID); err.Error != nil {
// 		return err.Error
// 	}
// 	var p Modifier
// 	if err := tx.First(&p, "id=?", op.ModifierID); err.Error != nil {
// 		return err.Error
// 	}
// 	tx.Model(&o).Update("total", o.Total-p.Price)
// 	return nil
// }
