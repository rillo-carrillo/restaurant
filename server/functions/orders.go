package functions

import (
	"github.com/rillo-carrillo/restaurant/server/db"
	"github.com/rillo-carrillo/restaurant/server/entities"
	"github.com/rillo-carrillo/restaurant/server/model"
)

//CreateOrder handles the creation of order.
func CreateOrder(id uint, or *model.OrderRequest) (uint, error) {
	order := entities.Order{
		EmployeeID: id,
		ClientID:   or.Client,
	}
	if err := db.SaveToDB(&order); err != nil {
		return 0, err
	}
	for _, o := range or.Order {
		orderProduct := entities.OrderProduct{
			OrderID:   order.ID,
			ProductID: o.Product,
			Comment:   o.Comment,
		}
		if err := db.SaveToDB(&orderProduct); err != nil {
			return 0, err
		}
		for _, m := range o.Modifiers {
			orderDetails := entities.OrderDetail{
				OrderProductID: orderProduct.ID,
				ModifierID:     m,
			}
			if err := db.SaveToDB(&orderDetails); err != nil {
				return 0, err
			}
		}
	}
	return order.ID, nil
}
