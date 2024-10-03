package tasks_test

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/database"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/scheduler/tasks"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
)

func TestFlushExpiredRefreshTokens(t *testing.T) {
	t.Setenv("ENV", "test")
	c, err := config.New()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	db, err := database.New(c)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// insert admin in DB
	admin := &types.Admin{
		Email:    fmt.Sprintf("%s-%d@%s.com", "test", rand.IntN(1000), "test"),
		Name:     "test",
		Password: "test",
	}
	if err := db.Model(admin).Create(&admin).Error; err != nil {
		t.Errorf("Error: %v", err)
	}
	//insert refresh tokens with expired_at < now()

	token, err := handlers.GenerateRefreshToken()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	refreshToken := &types.RefreshToken{
		Token:     token,
		AdminID:   admin.ID,
		ExpiresAt: time.Now().Add(-time.Hour),
	}
	if err := db.Model(refreshToken).Create(&refreshToken).Error; err != nil {
		t.Errorf("Error: %v", err)
	}
	// get the number of refresh tokens with expired_at < now()
	var refreshTokensExpiredCount int64
	if err := db.Model(&types.RefreshToken{}).Where("expires_at < ?", time.Now()).Count(&refreshTokensExpiredCount).Error; err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	if refreshTokensExpiredCount == 0 {
		t.Errorf("Error: %v", "No expired refresh tokens found.")
	}
	// call FlushExpiredRefreshTokens()
	tasks.FlushExpiredRefreshTokens()
	if err := db.Model(&types.RefreshToken{}).Where("expires_at < ?", time.Now()).Count(&refreshTokensExpiredCount).Error; err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	if refreshTokensExpiredCount != 0 {
		t.Errorf("Error: %v", "Expired refresh tokens not deleted.")
		return
	}
}
