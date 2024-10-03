package handlers_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/handlers"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
)

func TestGenerateAuthToken(t *testing.T) {
	s, err := handlers.New()
	if err != nil {
		t.Errorf("GenerateAuthToken() error = %v", err)
		return
	}

	t.Run("Success", func(t *testing.T) {
		email := fmt.Sprintf("user-%v@example.com", rand.IntN(1000000))
		customerRequest := &proto_v1.CreateCustomerRequest{
			Name:     "User",
			Email:    email,
			Password: "password",
			Mobile:   fmt.Sprintf("011111%v", rand.IntN(2000000)),
		}
		_, err = s.CreateCustomer(context.Background(), customerRequest)
		if err != nil {
			t.Errorf("GenerateAuthToken() error = %v", err)
			return
		}
		resp, err := s.GenerateAuthToken(context.Background(), &proto_v1.GenerateAuthTokenRequest{
			Email:    email,
			Password: "password",
		})
		if err != nil {
			t.Errorf("GenerateAuthToken() error = %v", err)
			return
		}
		if resp.Token == "" {
			t.Errorf("GenerateAuthToken() got = %v, want not empty", resp.Token)
			return
		}
		if resp.RefreshToken == "" {
			t.Errorf("GenerateAuthToken() got = %v, want not empty", resp.RefreshToken)
			return
		}

	})

}
