package handlers_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"golang.org/x/crypto/bcrypt"
)

func TestGetAuthCustomer(t *testing.T) {
	s, err := handlers.New()
	if err != nil {
		t.Errorf("ShowCustomer() error = %v", err)
		return
	}

	t.Run("Success", func(t *testing.T) {
		pass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		c := &types.Customer{
			Name:     "test",
			Email:    fmt.Sprintf("user-%v@example.com", rand.IntN(1000000000)),
			Password: string(pass),
			Mobile:   fmt.Sprintf("011111111%v", rand.IntN(2000)),
		}
		if err := s.DB.Create(c).Error; err != nil {
			t.Errorf("ShowCustomer() error = %v", err)
			return
		}
		resp, err := s.GenerateAuthToken(context.Background(), &proto_v1.GenerateAuthTokenRequest{
			Email:    c.Email,
			Password: "password",
		})
		if err != nil {
			t.Errorf("GenerateAuthToken() error = %v", err)
			return
		}
		if resp.GetToken() == "" {
			t.Errorf("GenerateAuthToken() got = %v, want not empty", resp.Token)
			return
		}
		getAuthCustomerResp, err := s.GetAuthCustomer(context.Background(), &proto_v1.GetAuthCustomerRequest{
			Token: resp.Token,
		})
		if err != nil {
			t.Errorf("GetAuthCustomer() error = %v", err)
			return
		}
		if getAuthCustomerResp.GetCustomer().Email != c.Email {
			t.Errorf("GetAuthCustomer() got = %v, want %v", getAuthCustomerResp.GetCustomer().Email, c.Email)
			return
		}

	})
}
