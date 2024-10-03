package cmd

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmedmohamed24/ecommerce-microservices/order/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/order/internals/consumers"
	"github.com/ahmedmohamed24/ecommerce-microservices/order/internals/handlers"
	order_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start GRPC server",
	Long:  `Start GRPC server`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.New()
		if err != nil {
			log.Println(err)
			return
		}
		h, err := handlers.New(c)
		if err != nil {
			log.Println(err)
			return
		}
		lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port))
		if err != nil {
			log.Println(err)
			return
		}
		defer lis.Close()
		grpcServer := grpc.NewServer()
		reflection.Register(grpcServer)
		go handleGracefullShutdown(grpcServer, h)
		order_protos_v1.RegisterOrderServiceServer(grpcServer, h)

		// register customers client
		customerClientConn, err := grpc.NewClient(fmt.Sprintf("%v:%v", c.Services.Customer.Host, c.Services.Customer.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println(err)
			return
		}
		h.CustomerServiceClient = order_protos_v1.NewCustomerServiceClient(customerClientConn)

		// register payments client
		paymentClientConn, err := grpc.NewClient(fmt.Sprintf("%v:%v", c.Services.Payment.Host, c.Services.Payment.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println(err)
			return
		}
		h.PaymentServiceClient = order_protos_v1.NewPaymentServiceClient(paymentClientConn)

		// register products client
		productClientConn, err := grpc.NewClient(fmt.Sprintf("%v:%v", c.Services.Product.Host, c.Services.Product.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println(err)
			return
		}
		h.ProductServiceClient = order_protos_v1.NewProductServiceClient(productClientConn)

		// start the consumer of the RabbitMQ
		if err = consumers.SuccessPaymentConsumer(h); err != nil {
			log.Println(err)
			return
		}

		err = grpcServer.Serve(lis)
		if !errors.Is(err, grpc.ErrServerStopped) {
			panic(err)
		} else {
			fmt.Println("Server stopped gracefully")
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
func handleGracefullShutdown(grpcServer *grpc.Server, s *handlers.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("Shutting down server...")
	for _, closer := range s.Closers {
		if err := closer.Close(); err != nil {
			log.Println(err)
		}

	}
	grpcServer.GracefulStop()

}
