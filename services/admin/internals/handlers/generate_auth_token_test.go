package handlers_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"golang.org/x/crypto/bcrypt"
)

// write testing for this function
func TestGenerateAuthToken(t *testing.T) {
	t.Setenv("ENV", "test")
	// create a new admin
	server := handlers.NewAdminServer()
	password := "password"
	hashedPasssword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	admin := &types.Admin{
		Email:    fmt.Sprintf("user-%d@ecommerce.com", rand.IntN(10000)),
		Name:     fmt.Sprintf("user-%d", rand.IntN(100)),
		Password: string(hashedPasssword),
	}
	if err := server.DB.Create(admin).Error; err != nil {
		t.Fatal(err)
	}
	// generate the auth token
	authToken, err := server.GenerateAuthToken(context.Background(), &pb.GenerateAuthTokenRequest{
		Email:    admin.Email,
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}
	// check if the auth token is not empty
	if authToken.GetAuthToken() == "" {
		t.Fatal("auth token is empty")
	}
	// check if the refresh token is not empty
	if authToken.GetRefreshToken() == "" {
		t.Fatal("refresh token is empty")
	}
	// check if the refresh token is stored in the database
	var refreshToken = &types.RefreshToken{}
	if err := server.DB.Last(&refreshToken, "token = ?", authToken.GetRefreshToken()).Error; err != nil {
		t.Fatal(err)
	}
	// check if the refresh token is not expired
	if refreshToken.ExpiresAt.Before(time.Now()) {
		t.Fatal("refresh token is expired")
	}
	// check if the refresh token is for the correct admin
	if refreshToken.AdminID != admin.ID {
		t.Fatal("refresh token is for the wrong admin")
	}
	// check if the refresh token is not expired
	if refreshToken.ExpiresAt.Before(time.Now()) {
		t.Fatal("refresh token is expired")
	}
}
