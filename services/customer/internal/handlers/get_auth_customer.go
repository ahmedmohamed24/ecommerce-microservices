package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) GetAuthCustomer(ctx context.Context, req *proto_v1.GetAuthCustomerRequest) (*proto_v1.GetAuthCustomerResponse, error) {
	token, err := jwt.Parse(req.GetToken(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.Config.APP_KEY), nil
	})
	if err != nil {
		return nil, status.Error(http.StatusUnauthorized, "invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, status.Error(http.StatusUnauthorized, "invalid token")
	}

	customer := &types.Customer{}
	if err := s.DB.Model(&types.Customer{}).Where("id = ?", claims["customer_id"]).First(&customer).Error; err != nil {
		return nil, status.Error(http.StatusUnauthorized, "invalid token")
	}

	return &proto_v1.GetAuthCustomerResponse{
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
