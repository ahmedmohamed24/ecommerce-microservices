package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/configs"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/logger"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the gRPC server",
	Long:  `Start the gRPC server which will listen on the port specified in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
func Run() {
	c, err := configs.New()
	if err != nil {
		panic(err)
	}
	logger := logger.New(c)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errorCh := make(chan error, 1)
	// start HTTP server (File Server)
	httpServer := server.NewHttpServer(c)
	go func() {
		logger.Info("Serving HTTP on " + fmt.Sprintf("%v:%v", c.FileServer.Host, c.FileServer.Port))
		err := httpServer.Server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			errorCh <- err
		}

	}()
	// start gRPC server
	grpcServer, err := server.NewGRPCServer(c)
	if err != nil {
		logger.Error(err.Error())
	}
	go func() {
		conn, err := net.Listen("tcp", fmt.Sprintf("%v:%v", c.GrpcServer.Host, c.GrpcServer.Port))
		if err != nil {
			logger.Error(err.Error())
		}
		defer conn.Close()

		logger.Info(fmt.Sprintf("Serving gRPC on %v", conn.Addr()))
		err = grpcServer.GS.Serve(conn)
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			errorCh <- err
		}
	}()
	// Handle graceful shutdown
	go func() {
		defer cancel()
		GracefullShutdown(logger, httpServer, grpcServer)
	}()
	for {
		select {
		case err := <-errorCh:
			logger.Error("Error: " + err.Error())
			cancel()
			return
		case <-ctx.Done():
			logger.Info("Context canceled. Shutting down")
			return
		}

	}

}
func GracefullShutdown(logger *slog.Logger, closers ...io.Closer) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	logger.Info("Context canceled. Shutting down")
	for _, c := range closers {
		if err := c.Close(); err != nil {
			fmt.Println("Error closing: ", err)
		}
	}
}
