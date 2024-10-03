package tasks

import (
	"log"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/database"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
)

func FlushExpiredRefreshTokens() {
	log.Println("Flushing expired refresh tokens...")
	refreshTokens := types.RefreshToken{}
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
