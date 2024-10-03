package types

import "time"

const (
	_ = iota
	OrderStatusPending
	OrderStatusCompleted
	OrderStatusCancelled
)

type OrderStatus struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderStatus) TableName() string {
	return "order_status"
}
