package handlers

import (
	"github.com/ahmedmohamed24/ecommerce-microservices/customer/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/customer/database"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"gorm.io/gorm"
)

type GrpcServer struct {
	DB     *gorm.DB
	Config *config.Config
	proto_v1.UnimplementedCustomerServiceServer
}

func New() (*GrpcServer, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}
	db, err := database.New(c)
	if err != nil {
		return nil, err
	}

	return &GrpcServer{
		DB:     db,
		Config: c,
	}, nil

}
