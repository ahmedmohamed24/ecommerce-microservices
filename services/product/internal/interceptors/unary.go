package interceptors

import (
	"context"
	"log"
	"strings"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	AdminClient *grpc.ClientConn
}

func New(adminClient *grpc.ClientConn) *Interceptor {
	return &Interceptor{
		AdminClient: adminClient,
	}
}

func (interceptor *Interceptor) UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	adminAuthRoutes := []string{"CreateProduct", "UpdateProduct", "DeleteProduct"}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	for _, router := range adminAuthRoutes {
		if strings.HasSuffix(info.FullMethod, router) {
			if _, ok := md["authorization"]; !ok {
				return nil, status.Error(codes.Unauthenticated, "provide a valid bearer token")
			}
			token := strings.TrimPrefix(md["authorization"][0], "Bearer ")
			adminClient := proto_v1.NewAdminsServiceClient(interceptor.AdminClient)
			resp, err := adminClient.GetAdmin(ctx, &proto_v1.GetAdminRequest{AuthToken: token})
			if err != nil {
				log.Println(err)
				return nil, status.Error(codes.Unauthenticated, "unauthenticated")
			}
			if resp.GetId() == 0 {
				return nil, status.Error(codes.Unauthenticated, "unauthenticated")
			}
			ctx = context.WithValue(ctx, &types.AdminContext{}, resp)

		}
	}
	resp, err := handler(ctx, req)
	return resp, err
}
