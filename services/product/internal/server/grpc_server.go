package server

import (
	"fmt"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/configs"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/interceptors"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	GS           *grpc.Server
	AdminsClient *grpc.ClientConn
}

func (server *GrpcServer) Close() error {
	if err := server.AdminsClient.Close(); err != nil {
		return err
	}
	fmt.Println("Closing gRPC server")
	server.GS.GracefulStop()
	fmt.Println("gRPC server closed")
	return nil
}

func NewGRPCServer(c *configs.Config) (*GrpcServer, error) {
	adminClient, _ := grpc.NewClient(fmt.Sprintf("%v:%v", c.Services.Admin.Host, c.Services.Admin.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	productServiceServer, err := handlers.New()
	if err != nil {
		return nil, err
	}

	interceptor := interceptors.New(adminClient)
	gs := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryAuthInterceptor),
		grpc.StreamInterceptor(interceptor.StreamAuthInterceptor),
	)
	server := &GrpcServer{
		GS:           gs,
		AdminsClient: adminClient,
	}
	grpc_health_v1.RegisterHealthServer(gs, productServiceServer)
	reflection.Register(server.GS)
	proto_v1.RegisterProductServiceServer(server.GS, productServiceServer)
	return server, nil
}
