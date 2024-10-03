package handlers_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"golang.org/x/crypto/bcrypt"
)

func TestRefreshToken(t *testing.T) {
	t.Setenv("ENV", "test")
	server := handlers.NewAdminServer()
	// create a new admin
	password := "password"
	hashedPasssword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	admin := &types.Admin{
		Email:    fmt.Sprintf("user-%d", rand.IntN(100)),
		Password: string(hashedPasssword),
		Name:     fmt.Sprintf("user-%d", rand.IntN(100)),
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
	// refresh the auth token
	refreshToken, err := server.RefreshAuthToken(context.Background(), &pb.RefreshAuthTokenRequest{
		AdminId:      int64(admin.ID),
		RefreshToken: authToken.GetRefreshToken(),
	})
	if err != nil {
		t.Fatal(err)
	}
	// check if the auth token is not empty
	if refreshToken.GetAuthToken() == "" {
		t.Fatal("auth token is empty")
	}

	// check logging user with the new token
	adminResponse, err := server.GetAdmin(context.Background(), &pb.GetAdminRequest{
		AuthToken: refreshToken.GetAuthToken(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if adminResponse == nil {
		t.Fatal("admin is nil")
	}
	if adminResponse.GetId() != int64(admin.ID) {
		t.Fatal("admin id is incorrect")
	}
}
