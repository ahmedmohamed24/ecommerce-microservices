package types

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             uint    `gorm:"primarykey"`
	Title          string  `json:"title" validate:"required,min=3,max=100"`
	Price          float64 `json:"price" validate:"required,numeric,gt=0"`
	Description    string  `json:"description,omitempty" validate:"required,min=3,max=1000"`
	CreatorAdminID int     `json:"creator_admin_id"`
	AvailableUnits int     `json:"available_units" validate:"required,numeric"`
	ProductImages  []ProductImage
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
