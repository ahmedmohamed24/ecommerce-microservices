package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/utils"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s AdminServer) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.GetAdminResponse, error) {
	var admin types.Admin
	claims, err := utils.ParseAndValidateAdminAccessToken(req.GetAuthToken(), s.C.App.Key)
	if err != nil {
		return nil, err
	}
	if err := s.DB.First(&admin, claims["admin_id"]).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(http.StatusNotFound, "admin not found")
		}
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}
	adminResponse := &pb.GetAdminResponse{
		Id:        int64(admin.ID),
		Name:      admin.Name,
		Email:     admin.Email,
		CreatedAt: admin.CreatedAt.Unix(),
		UpdatedAt: admin.UpdatedAt.Unix(),
	}

	return adminResponse, nil

}
