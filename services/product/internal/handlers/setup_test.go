package handlers_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/handlers"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func GetClientFromGRPCServer(t *testing.T) (proto_v1.ProductServiceClient, func()) {
	t.Setenv("ENV", "test")
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)
	baseServer := grpc.NewServer()
	productService, err := handlers.New()
	if err != nil {
		log.Printf("error creating new server: %v", err)
	}
	proto_v1.RegisterProductServiceServer(baseServer, productService)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	conn, err := grpc.NewClient(
		"passthrough:whatever",
		grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return lis.Dial()
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("error connecting to server:", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Fatal("error closing listener: ", err)
		}
		baseServer.Stop()
	}
	client := proto_v1.NewProductServiceClient(conn)

	return client, closer
}
