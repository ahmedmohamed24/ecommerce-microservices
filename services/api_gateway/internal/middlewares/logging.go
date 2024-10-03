package middlewares

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func LoggingMiddleware(hf runtime.HandlerFunc) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		log.Printf("Request URI: %s, params: %v \n", r.RequestURI, pathParams)
		hf(w, r, pathParams)
	}
}
