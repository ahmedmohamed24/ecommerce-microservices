package middlewares

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
)

func MetaDataMiddleware(ctx context.Context, r *http.Request) metadata.MD {
	md := make(map[string]string)
	if method, ok := runtime.RPCMethod(ctx); ok {
		md["method"] = method
	}
	if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
		md["pattern"] = pattern
	}
	return metadata.New(md)
}
