package types

import "time"

type Order struct {
	ID         uint        `gorm:"primarykey" json:"id"`
	CustomerID uint        `json:"customer_id"`
	TotalPrice float64     `json:"total_price"`
	StatusID   uint        `json:"status_id"`
	Status     OrderStatus `json:"status" gorm:"foreignKey:id;references:status_id"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
