package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) RefreshAuthToken(ctx context.Context, req *proto_v1.RefreshAuthTokenRequest) (*proto_v1.RefreshAuthTokenResponse, error) {
	refreshToken := &types.RefreshToken{}
	if err := s.DB.Preload("Customer").Model(refreshToken).Where("token = ?", req.GetRefreshToken()).Where("expires_at > ?", time.Now()).Take(refreshToken).Error; err != nil {
		return nil, status.Error(http.StatusUnauthorized, "invalid refresh token")
	}
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": refreshToken.CustomerID,
		"iss":         s.Config.ServiceName,
		"exp":         time.Now().Add(time.Minute * 15).Unix(),
		"iat":         time.Now().Unix(),
	}).SignedString([]byte(s.Config.APP_KEY))
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "internal server error")
	}

	return &proto_v1.RefreshAuthTokenResponse{
		Token: signedString,
	}, nil
}
