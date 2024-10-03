package types

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var ErrEmailOrPasswordIncorrect = errors.New("email or password is incorrect")

type Admin struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
