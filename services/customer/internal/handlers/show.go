package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *GrpcServer) GetCustomer(ctx context.Context, req *proto_v1.GetCustomerRequest) (*proto_v1.GetCustomerResponse, error) {
	customer := &types.Customer{}
	if req.Id == 0 {
		return nil, status.Error(http.StatusBadRequest, "id is required")
	}
	if err := s.DB.Model(customer).Where("id = ?", req.Id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(http.StatusNotFound, "customer not found")
		}
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "internal server error")
	}
	return &proto_v1.GetCustomerResponse{
		Customer: &proto_v1.Customer{
			Id:        uint32(customer.ID),
			Name:      customer.Name,
			Email:     customer.Email,
			Mobile:    customer.Mobile,
			CreatedAt: customer.CreatedAt.String(),
			UpdatedAt: customer.UpdatedAt.String(),
		},
	}, nil
}
