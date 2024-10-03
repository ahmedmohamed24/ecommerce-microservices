package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) UpdateCustomer(ctx context.Context, req *proto_v1.UpdateCustomerRequest) (*proto_v1.UpdateCustomerResponse, error) {
	validate := validator.New()
	err := validate.Struct(&types.Customer{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Mobile:   req.GetMobile(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	customer := &types.Customer{}
	if err := s.DB.Where("id = ?", req.GetId()).First(customer).Error; err != nil {
		return nil, status.Error(http.StatusNotFound, "customer not found")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "internal server error")
	}

	err = s.DB.Model(customer).Updates(&types.Customer{
		Email:    req.GetEmail(),
		Name:     req.GetName(),
		Mobile:   req.GetMobile(),
		Password: string(hashedPassword),
	}).Error
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "internal server error")
	}
	return &proto_v1.UpdateCustomerResponse{
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
