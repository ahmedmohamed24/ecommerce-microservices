package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/handlers"
	pb "github.com/ahmedmohamed24/ecommerce-microservices/admin/protos/v1/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

func Serve() error {
	s := handlers.NewAdminServer()
	grpcHost := fmt.Sprintf("%v:%v", s.C.GRPC.Host, s.C.GRPC.Port)
	fmt.Println("Starting server on ", grpcHost)
	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer conn.Close()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryLoggingInterceptor),   // Register unary interceptor
		grpc.StreamInterceptor(StreamLoggingInterceptor), // Register stream interceptor
	)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Printf("\nShutting down server\n")
		grpcServer.GracefulStop()
		fmt.Println("Server shutdown completed")

	}()
	pb.RegisterAdminsServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	return grpcServer.Serve(conn)
}
func UnaryLoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	// Retrieve metadata (headers)
	md, _ := metadata.FromIncomingContext(ctx)

	// Call the handler to finish processing the request
	resp, err := handler(ctx, req)

	// Log the request method, headers, and error
	log.Printf("Request - Method: %s, Headers: %v, Duration: %s, Error: %v", info.FullMethod, md, time.Since(start), err)

	return resp, err
}
func StreamLoggingInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()

	// Call the handler to finish processing the stream request
	err := handler(srv, ss)

	// Log details about the request
	log.Printf("Stream Request - Method: %s, Duration: %s, Error: %v", info.FullMethod, time.Since(start), err)

	return err
}
