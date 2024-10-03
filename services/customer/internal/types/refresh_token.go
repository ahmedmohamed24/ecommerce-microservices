package types

import (
	"log"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/customer/database"
)

type RefreshToken struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Token      string    `json:"token"`
	CustomerID uint      `json:"customer_id"`
	Customer   Customer  `json:"customer"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (refreshToken *RefreshToken) Run() {
	log.Println("Flushing expired refresh tokens...")
	refreshTokens := RefreshToken{}
	c, err := config.New()
	if err != nil {
		log.Println(err)
	}
	db, err := database.New(c)
	if err != nil {
		log.Println(err)
	}
	res := db.Model(&refreshTokens).Where("expires_at < ?", time.Now()).Delete(&refreshTokens)
	if err := res.Error; err != nil {
		log.Println(err)
	}
	log.Printf("Flushed expired refresh tokens: %d\n", int(res.RowsAffected))

}
