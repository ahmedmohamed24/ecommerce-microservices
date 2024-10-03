package handlers

import (
	"context"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/app"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type ProductServiceServer struct {
	proto_v1.UnimplementedProductServiceServer
	App *app.App
}

func New() (*ProductServiceServer, error) {
	app, err := app.New()
	if err != nil {
		return nil, utils.Error("we ran into a problem, please try again later", err, 500)
	}
	return &ProductServiceServer{
		App: app,
	}, nil
}

func (s *ProductServiceServer) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *ProductServiceServer) Watch(in *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
