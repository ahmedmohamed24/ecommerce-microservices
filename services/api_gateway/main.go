package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmedmohamed24/ecommerce-microservices/order/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/order/internal/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/order/internal/middlewares"
	admin_proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen/admin"
	customer_proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen/customer"
	order_proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen/order"
	payment_proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen/payment"
	product_proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen/product"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx := context.Background()
	mux := runtime.NewServeMux(
		runtime.WithMiddlewares(middlewares.LoggingMiddleware),
		runtime.WithMetadata(middlewares.MetaDataMiddleware),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = admin_proto_v1.RegisterAdminsServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%v:%v", cfg.AdminService.Host, cfg.AdminService.Port), opts)
	if err != nil {
		log.Fatal(err)
	}
	err = product_proto_v1.RegisterProductServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%v:%v", cfg.ProductService.Host, cfg.ProductService.Port), opts)
	if err != nil {
		log.Fatal(err)
	}
	err = customer_proto_v1.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%v:%v", cfg.CustomerService.Host, cfg.CustomerService.Port), opts)
	if err != nil {
		log.Fatal(err)
	}
	err = payment_proto_v1.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%v:%v", cfg.PaymentService.Host, cfg.PaymentService.Port), opts)
	if err != nil {
		log.Fatal(err)
	}
	err = order_proto_v1.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%v:%v", cfg.OrderService.Host, cfg.OrderService.Port), opts)
	if err != nil {
		log.Fatal(err)
	}
	mux.HandlePath(http.MethodPost, "/api/v1/product/image", handlers.UploadProductImage)
	server := http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf("%v:%v", cfg.ApiGateway.Host, cfg.ApiGateway.Port),
	}
	go handleGracefulShutdown(&server, &ctx)
	log.Println("HTTP server is running on: " + fmt.Sprintf("%v:%v", cfg.ApiGateway.Host, cfg.ApiGateway.Port))
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
		return
	}
	log.Println("Server stopped")

}

func handleGracefulShutdown(server *http.Server, ctx *context.Context) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("Shutting down the server")
	shutdownCtx, cancel := context.WithTimeout(*ctx, 20)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println(err)
	}
}
