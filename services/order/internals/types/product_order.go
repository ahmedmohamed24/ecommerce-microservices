package types

import "time"

type ProductOrder struct {
	ProductID uint      `json:"product_id"`
	OrderID   uint      `json:"order_id"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ProductOrder) TableName() string {
	return "products_orders"
}
