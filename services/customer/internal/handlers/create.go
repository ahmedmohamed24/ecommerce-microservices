package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	validate "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) CreateCustomer(ctx context.Context, req *proto_v1.CreateCustomerRequest) (*proto_v1.CreateCustomerResponse, error) {
	customer := &types.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Mobile:   req.Mobile,
	}
	err := validate.New().Struct(customer)
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	if err := s.DB.Model(customer).Where("email = ?", req.Email).First(&customer).Error; err == nil {
		return nil, status.Errorf(http.StatusBadRequest, "%s", "customer email already exists")
	}
	if err := s.DB.Model(customer).Where("mobile = ?", req.Mobile).First(&customer).Error; err == nil {
		return nil, status.Errorf(http.StatusBadRequest, "%s", "customer mobile already exists")
	}
	p, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(http.StatusInternalServerError, "%s", "internal server error")
	}
	customer.Password = string(p)
	if err := s.DB.Create(customer).Error; err != nil {
		return nil, err
	}

	return &proto_v1.CreateCustomerResponse{
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
