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

func TestUpdateAdmin(t *testing.T) {
	t.Setenv("ENV", "test")
	server := handlers.NewAdminServer()
	// create a new admin
	password := "password"
	hashedPasssword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	admin := &types.Admin{
		Email:    fmt.Sprintf("User-%d@ecommerce.com", rand.IntN(100)),
		Password: string(hashedPasssword),
		Name:     fmt.Sprintf("User-%d", rand.IntN(100)),
	}
	if err := server.DB.Create(admin).Error; err != nil {
		t.Fatal(err)
	}
	// generate the auth token
	authToken, err := server.GenerateAuthToken(context.Background(), &pb.GenerateAuthTokenRequest{
		Email:    admin.Email,
		Password: "password",
	})
	if err != nil {
		t.Fatal(err)
	}
	// update the admin
	updatedEmail := fmt.Sprintf("new-email%v@ecommerce.com", rand.IntN(100))
	adminResponse, err := server.UpdateAdmin(context.Background(), &pb.UpdateAdminRequest{
		AuthToken: authToken.GetAuthToken(),
		Name:      "new name",
		Email:     updatedEmail,
		Password:  "new password",
	})
	if err != nil {
		t.Fatal(err)
	}
	if adminResponse == nil {
		t.Fatal("admin is nil")
	}
	// get admin from DB
	var updatedAdmin types.Admin
	if err := server.DB.First(&updatedAdmin, admin.ID).Error; err != nil {
		t.Fatal(err)
	}
	if updatedAdmin.Email != updatedEmail {
		t.Fatal("admin email is incorrect")
	}
	// check if the password is updated
	if err := bcrypt.CompareHashAndPassword([]byte(updatedAdmin.Password), []byte("new password")); err != nil {
		t.Fatal("admin password is incorrect")
	}

}
