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

func TestUpdateCustomer(t *testing.T) {
	s, err := handlers.New()
	if err != nil {
		t.Errorf("UpdateCustomer() error = %v", err)
		return
	}
	t.Run("Success", func(t *testing.T) {
		c := &types.Customer{
			Name:     "test",
			Email:    fmt.Sprintf("user-%v@example.com", rand.IntN(1000000000)),
			Password: "password",
			Mobile:   fmt.Sprintf("011111111%v", rand.IntN(2000)),
		}
		if err := s.DB.Create(c).Error; err != nil {
			t.Errorf("UpdateCustomer() error = %v", err)
			return
		}
		_, err = s.UpdateCustomer(context.Background(), &proto_v1.UpdateCustomerRequest{
			Id:       uint32(c.ID),
			Name:     "new name",
			Email:    c.Email,
			Password: c.Password,
			Mobile:   c.Mobile,
		})
		if err != nil {
			t.Errorf("UpdateCustomer() error = %v", err)
			return
		}
		newCustomer := types.Customer{}
		if err := s.DB.First(&newCustomer, c.ID).Error; err != nil {
			t.Errorf("UpdateCustomer() error = %v", err)
			return
		}
		if newCustomer.Name != "new name" {
			t.Errorf("UpdateCustomer() error = %v, want %v", newCustomer.Name, "new name")
			return
		}

	})

}
