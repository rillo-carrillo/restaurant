package model

//OrderDetails array with Entity definition
type OrderDetails struct {
	// `json:"products"`
	Product   uint
	Comment   string
	Modifiers []uint
}

// OrderCreateResponse struct
type OrderCreateResponse struct {
	ID     int
	Status bool
}

//OrderRequest obj
type OrderRequest struct {
	Client uint
	Order  []OrderDetails `json:"order"`
}
