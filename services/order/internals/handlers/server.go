package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/order/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/order/database"
	"github.com/ahmedmohamed24/ecommerce-microservices/order/internals/utils"
	order_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	Config                *config.Config
	DB                    *gorm.DB
	CustomerServiceClient order_protos_v1.CustomerServiceClient
	ProductServiceClient  order_protos_v1.ProductServiceClient
	PaymentServiceClient  order_protos_v1.PaymentServiceClient
	order_protos_v1.UnimplementedOrderServiceServer
	Closers []io.Closer
}

func New(c *config.Config) (*Server, error) {
	db, err := database.New(c)
	if err != nil {
		return nil, err

	}
	return &Server{
		Config: c,
		DB:     db,
	}, nil
}

func (s *Server) GetAuthCustomerFromContext(ctx context.Context) (*order_protos_v1.Customer, error) {
	token, err := utils.GetAuthTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.CustomerServiceClient.GetAuthCustomer(ctx, &order_protos_v1.GetAuthCustomerRequest{
		Token: token,
	})
	if err != nil {
		return nil, status.Error(http.StatusUnauthorized, "please, provide the authorization token")
	}
	if resp.Customer.Id == 0 {
		return nil, status.Error(http.StatusUnauthorized, "please, provide the authorization token")
	}
	return resp.Customer, nil
}
