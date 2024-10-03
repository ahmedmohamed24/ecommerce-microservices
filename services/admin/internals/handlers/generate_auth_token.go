package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

func (s AdminServer) GenerateAuthToken(ctx context.Context, req *pb.GenerateAuthTokenRequest) (*pb.GenerateAuthTokenResponse, error) {
	var admin = &types.Admin{}
	// check if the admin exists for the given credentials (email,password)
	if err := s.DB.First(&admin, "email = ?", req.GetEmail()).Error; err != nil {
		return nil, status.Error(http.StatusUnauthorized, types.ErrEmailOrPasswordIncorrect.Error())
	}
	fmt.Println(admin.Password, req.GetPassword())
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(http.StatusUnauthorized, types.ErrEmailOrPasswordIncorrect.Error())
	}
	// generate the auth token
	authJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    admin.Email,
		"admin_id": admin.ID,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	authToken, err := authJWT.SignedString([]byte(s.C.App.Key))
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "something went wrong, please try again")
	}
	// generate the refresh token

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return nil, err
	}
	//  expire old refresh tokens for the admin (could be ignored to allow multiple sessions)
	if err := s.DB.Model(&types.RefreshToken{}).Where("admin_id = ?", admin.ID).Update("expires_at", time.Now().Add(time.Duration(-1)*time.Minute)).Error; err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "something went wrong, please try again")
	}
	// store the new refresh token in the storage
	var refreshTokenType = &types.RefreshToken{
		Token:     refreshToken,
		AdminID:   admin.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := s.DB.Create(&refreshTokenType).Error; err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &pb.GenerateAuthTokenResponse{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}, nil
}

func GenerateRefreshToken() (string, error) {
	tokenBase := make([]byte, 32)
	if _, err := rand.Read(tokenBase); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(tokenBase), nil

}
