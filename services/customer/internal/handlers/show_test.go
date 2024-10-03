package handlers_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"google.golang.org/grpc/status"
)

func TestShow(t *testing.T) {
	s, err := handlers.New()
	if err != nil {
		t.Errorf("ShowCustomer() error = %v", err)
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
			t.Errorf("ShowCustomer() error = %v", err)
			return
		}
		res, err := s.GetCustomer(context.Background(), &proto_v1.GetCustomerRequest{
			Id: uint32(c.ID),
		})
		if err != nil {
			t.Errorf("ShowCustomer() error = %v", err)
			return
		}
		if res.Customer.Email != c.Email {
			t.Errorf("different customer email")
			return
		}

	})
	t.Run("Not Found Customer", func(t *testing.T) {
		c := &types.Customer{
			Name:     "test",
			Email:    fmt.Sprintf("user-%v@example.com", rand.IntN(1000000000)),
			Password: "password",
			Mobile:   fmt.Sprintf("011111111%v", rand.IntN(2000)),
		}
		if err := s.DB.Create(c).Error; err != nil {
			t.Errorf("ShowCustomer() error = %v", err)
			return
		}
		_, err := s.GetCustomer(context.Background(), &proto_v1.GetCustomerRequest{
			Id: uint32(c.ID + 1),
		})
		if err == nil {
			t.Errorf("Found Customer that does not exist")
			return
		}
		status, _ := status.FromError(err)
		if status.Code() != http.StatusNotFound {
			t.Errorf("Expected Not Found Error")
			return
		}

	})
}
