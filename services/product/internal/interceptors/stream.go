package interceptors

import (
	"log"
	"time"

	"google.golang.org/grpc"
)

func (interceptor *Interceptor) StreamAuthInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	err := handler(srv, ss)
	log.Printf("Stream Request - Method: %s, Duration: %s, Error: %v", info.FullMethod, time.Since(start), err)
	return err
}
