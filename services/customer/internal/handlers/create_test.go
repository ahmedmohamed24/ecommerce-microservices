package handlers_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
)

func TestCreateCustomer(t *testing.T) {
	s, err := handlers.New()
	if err != nil {
		t.Errorf("CreateCustomer() error = %v", err)
		return
	}

	t.Run("Success", func(t *testing.T) {
		customerRequest := &proto_v1.CreateCustomerRequest{
			Name:     "User",
			Email:    fmt.Sprintf("user-%v@example.com", rand.IntN(100000000)),
			Password: "password",
			Mobile:   fmt.Sprintf("011111%v", rand.IntN(2000000)),
		}
		_, err = s.CreateCustomer(context.Background(), customerRequest)
		if err != nil {
			t.Errorf("CreateCustomer() error = %v", err)
			return
		}
	})
	t.Run("email is unique", func(t *testing.T) {
		email := fmt.Sprintf("user-%v@example.com", rand.IntN(1000000))
		err := s.DB.Create(&types.Customer{
			Name:     "User",
			Email:    email,
			Password: "password",
			Mobile:   fmt.Sprintf("011111%v", rand.IntN(2000000)),
		}).Error
		if err != nil {
			t.Errorf("CreateCustomer() error = %v", err)
			return
		}
		customerRequest := &proto_v1.CreateCustomerRequest{
			Name:     "User",
			Email:    email,
			Password: "password",
			Mobile:   fmt.Sprintf("0111111%v", rand.IntN(20000000)),
		}
		_, err = s.CreateCustomer(context.Background(), customerRequest)
		if err == nil {
			t.Errorf("No validation for email uniqueness")
			return
		}
	})

}
