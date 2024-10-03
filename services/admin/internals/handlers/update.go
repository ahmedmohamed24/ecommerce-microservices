package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/utils"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

func (s AdminServer) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminRequest) (*pb.UpdateAdminResponse, error) {
	var admin = &types.Admin{}
	claims, err := utils.ParseAndValidateAdminAccessToken(req.GetAuthToken(), s.C.App.Key)
	if err != nil {
		return nil, status.Error(http.StatusUnauthorized, err.Error())
	}
	if err := s.DB.First(&admin, claims["admin_id"]).Error; err != nil {
		return nil, status.Error(http.StatusUnauthorized, err.Error())
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "something went wrong, please try again")
	}
	admin.Name = req.GetName()
	admin.Email = req.GetEmail()
	admin.Password = string(password)
	if err := s.DB.Save(&admin).Error; err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "something went wrong, please try again")
	}
	return &pb.UpdateAdminResponse{
		Id:        int64(admin.ID),
		Name:      admin.Name,
		Email:     admin.Email,
		CreatedAt: admin.CreatedAt.Unix(),
		UpdatedAt: admin.UpdatedAt.Unix(),
	}, nil
}
