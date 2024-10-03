package handlers

import (
	"fmt"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/database"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"gorm.io/gorm"
)

type AdminServer struct {
	pb.UnimplementedAdminsServiceServer
	C  *config.Config
	DB *gorm.DB
}

func NewAdminServer() *AdminServer {
	c, err := config.New()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db, err := database.New(c)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &AdminServer{
		C:  c,
		DB: db,
	}

}
