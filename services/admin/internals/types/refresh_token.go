package types

import (
	"errors"
	"time"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")

type RefreshToken struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Token     string    `json:"token"`
	AdminID   uint      `json:"admin_id"`
	Admin     Admin     `json:"admin"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
