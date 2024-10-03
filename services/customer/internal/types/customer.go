package types

import (
	"time"
)

type Customer struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-" validate:"required,min=6"`
	Mobile    string    `json:"mobile" validate:"required,min=9,max=14"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
