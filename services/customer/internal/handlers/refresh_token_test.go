package handlers_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/handlers"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
)

func TestRefreshAuthToken(t *testing.T) {
	s, err := handlers.New()
	if err != nil {
		t.Errorf("RefreshAuthToken() error = %v", err)
		return
	}

	t.Run("Success", func(t *testing.T) {
		email := fmt.Sprintf("user-%v@example.com", rand.IntN(1000000000))
		customerRequest := &proto_v1.CreateCustomerRequest{
			Name:     "User",
			Email:    email,
			Password: "password",
			Mobile:   fmt.Sprintf("011111111%v", rand.IntN(2000)),
		}
		_, err = s.CreateCustomer(context.Background(), customerRequest)
		if err != nil {
			t.Errorf("RefreshAuthToken() error = %v", err)
			return
		}
		resp, err := s.GenerateAuthToken(context.Background(), &proto_v1.GenerateAuthTokenRequest{
			Email:    email,
			Password: "password",
		})
		if err != nil {
			t.Errorf("RefreshAuthToken() error = %v", err)
			return
		}
		refreshTokenResp, err := s.RefreshAuthToken(context.Background(), &proto_v1.RefreshAuthTokenRequest{
			RefreshToken: resp.RefreshToken,
		})
		if err != nil {
			t.Errorf("RefreshAuthToken() error = %v", err)
			return
		}
		if refreshTokenResp.Token == "" {
			t.Errorf("expected token to be not empty")
			return
		}
		authCustomerResp, err := s.GetAuthCustomer(context.Background(), &proto_v1.GetAuthCustomerRequest{
			Token: refreshTokenResp.Token,
		})
		if err != nil {
			t.Errorf("RefreshAuthToken() error = %v", err)
			return
		}
		if authCustomerResp.GetCustomer().Email != email {
			t.Errorf("expected email to be %v, got %v", email, authCustomerResp.GetCustomer().Email)
			return
		}

	})
}
