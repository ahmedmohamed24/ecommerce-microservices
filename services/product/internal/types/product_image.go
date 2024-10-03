package types

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	ProductID    uint    `json:"product_id" validate:"required,number"`
	Product      Product `json:"omitempty" validate:"-"`
	OriginalName string  `json:"original_name" validate:"required,min=1,max=255"`
	StoredName   string  `json:"stored_name" validate:"required,min=1,max=255"`
	Extension    string  `json:"extension" validate:"required,min=1,max=255"`
	MimeType     string  `json:"mime_type" validate:"required,min=1,max=255"`
	FileSize     int64   `json:"file_size" validate:"required,number"`
	IsThumbnail  bool    `json:"is_thumbnail" validate:"boolean"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (ProductImage) TableName() string {
	return "product_images"
}
