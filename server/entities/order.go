package entities

import "gorm.io/gorm"

//STATUS INICIO, PREPARACION, REPARTO,LISTO

//Order contains the products, of the Order as well Employee ID
type Order struct {
	gorm.Model
	EmployeeID uint
	ClientID   uint
	Status     int `gorm:"default:1;check:status IN (1,2,3,4)"`
	Total      float32
	Products   []Product `gorm:"many2many:order_products;"`
}

//BeforeCreate Hook on new Orders set status as 1.
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ClientID == 0 {
		o.ClientID = 1
	}
	o.Status = 1
	err = nil
	return err
}

//OrderProduct define the join table
type OrderProduct struct {
	ProductID uint `gorm:"primaryKey"`
	OrderID   uint `gorm:"primaryKey"`
	Modifier  int
}

//AfterSave Hook to update the total of Order
func (op *OrderProduct) AfterSave(tx *gorm.DB) (err error) {
	var o Order
	if err := tx.First(&o, "id=?", op.OrderID); err.Error != nil {
		return err.Error
	}
	var p Product
	if err := tx.First(&p, "id=?", op.ProductID); err.Error != nil {
		return err.Error
	}
	tx.Model(&o).Update("total", o.Total+p.Price)
	return nil
}
