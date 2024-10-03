package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

type GenerateAuthTokenRequest struct {
	Email    string `json:"email" validate:"required,email,min=3"`
	Password string `json:"password" validate:"required,min=3"`
}

func (s *GrpcServer) GenerateAuthToken(ctx context.Context, req *proto_v1.GenerateAuthTokenRequest) (*proto_v1.GenerateAuthTokenResponse, error) {
	LoginRequest := &GenerateAuthTokenRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	validate := validator.New()
	if err := validate.Struct(LoginRequest); err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	c := &types.Customer{}
	if err := s.DB.Where("email = ?", req.Email).First(&c).Error; err != nil {
		return nil, status.Error(http.StatusUnauthorized, "invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(req.Password)); err != nil {
		return nil, status.Error(http.StatusUnauthorized, "invalid credentials")
	}
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": c.ID,
		"iss":         s.Config.ServiceName,
		"exp":         time.Now().Add(time.Minute * 15).Unix(),
		"iat":         time.Now().Unix(),
	}).SignedString([]byte(s.Config.APP_KEY))
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "internal server error")
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": c.ID,
		"iss":         s.Config.ServiceName,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
		"iat":         time.Now().Unix(),
	}).SignedString([]byte(s.Config.APP_KEY))
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "internal server error")
	}
	if err = s.DB.Create(&types.RefreshToken{
		Token:      refreshToken,
		CustomerID: c.ID,
		ExpiresAt:  time.Now().Add(time.Hour * 24 * 7),
	}).Error; err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "internal server error")
	}

	return &proto_v1.GenerateAuthTokenResponse{
		Token:        signedString,
		RefreshToken: refreshToken,
	}, nil
}
