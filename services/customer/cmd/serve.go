package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/customer/protos/v1/gen"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/handlers"
	"github.com/ahmedmohamed24/ecommerce-microservices/customer/internal/scheduler"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving the application",
	Long:  `This command will start the server and serve the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	server, err := handlers.New()
	if err != nil {
		panic(err)
	}
	scheduler.Schedule()
	ser := grpc.NewServer()
	go handleGracefullShutdown(ser)

	proto_v1.RegisterCustomerServiceServer(ser, server)
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", server.Config.Server.Host, server.Config.Server.Port))
	reflection.Register(ser)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running on " + lis.Addr().String())
	err = ser.Serve(lis)
	if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		panic(err)
	}

}
func handleGracefullShutdown(ser *grpc.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Printf("\n Shutting down server...\n")
	ser.GracefulStop()
}
