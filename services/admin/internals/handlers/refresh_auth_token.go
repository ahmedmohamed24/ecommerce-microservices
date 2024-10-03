package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/status"
)

func (s AdminServer) RefreshAuthToken(ctx context.Context, req *pb.RefreshAuthTokenRequest) (*pb.RefreshAuthTokenResponse, error) {
	var refreshTokenType = &types.RefreshToken{}
	err := s.DB.Model(refreshTokenType).Preload("Admin").Where("token = ?", req.GetRefreshToken()).Where("admin_id = ?", req.GetAdminId()).Where("expires_at > ?", time.Now()).First(refreshTokenType).Error
	if err != nil {
		return nil, status.Error(http.StatusUnauthorized, types.ErrInvalidRefreshToken.Error())
	}
	if refreshTokenType.Admin.ID == 0 {
		return nil, status.Error(http.StatusUnauthorized, types.ErrInvalidRefreshToken.Error())
	}
	authJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    refreshTokenType.Admin.Email,
		"admin_id": refreshTokenType.Admin.ID,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	authToken, err := authJWT.SignedString([]byte(s.C.App.Key))
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}
	return &pb.RefreshAuthTokenResponse{
		AuthToken: authToken,
	}, nil
}
