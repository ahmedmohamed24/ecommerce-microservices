package cmd

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmedmohamed24/ecommerce-microservices/payment/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/payment/internal/handlers"
	payment_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/payment/protos/v1/gen"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Payment Service",
	Long:  `Handle payment operations like creating a payment,`,
	Run: func(cmd *cobra.Command, args []string) {
		grpcServer := grpc.NewServer()
		config, err := config.New()
		if err != nil {
			log.Println(err)
			return
		}
		s, err := handlers.New(config)
		if err != nil {
			log.Println(err)
			return
		}
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sig
			for _, c := range s.Closers {
				if err := c.Close(); err != nil {
					log.Println(err)
				}
			}
			fmt.Println("shutting down server")
			grpcServer.GracefulStop()

		}()
		reflection.Register(grpcServer)
		payment_protos_v1.RegisterPaymentServiceServer(grpcServer, s)
		fmt.Println("serve started successfully at: ", fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
		lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
		err = grpcServer.Serve(lis)
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Println(err)
			return
		}
		fmt.Println("serve shutdown successfully")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
