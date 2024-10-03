package types

import (
	"time"

	"gorm.io/datatypes"
)

type Payment struct {
	ID            uint           `gorm:"primarykey"`
	OrderID       uint           `json:"order_id"`
	PaymentMethod string         `json:"payment_method"`
	PaidAmount    float64        `json:"paid_amount"`
	CardLast4     string         `json:"card_last_4" gorm:"column:card_last_four"`
	PaymentInfo   datatypes.JSON `json:"payment_info"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
