package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/configs"
	"github.com/go-chi/chi"
)

type HTTPServer struct {
	Server *http.Server
}

func NewHttpServer(c *configs.Config) *HTTPServer {

	httpServer := &HTTPServer{
		Server: &http.Server{
			Addr: fmt.Sprintf("%v:%v", c.FileServer.Host, c.FileServer.Port),
		},
	}
	router := chi.NewMux()
	httpServer.Server.Handler = router
	router.Get("/storage/*", http.StripPrefix("/storage", http.FileServer(http.Dir("storage"))).ServeHTTP)
	return httpServer
}

func (s *HTTPServer) Close() error {
	fmt.Println("Closing HTTP server")
	defer fmt.Println("HTTP server closed")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Server.Shutdown(ctx)
}
