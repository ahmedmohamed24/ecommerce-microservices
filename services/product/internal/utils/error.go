package utils

import (
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Error(message string, err error, statusCode uint32) error {
	if statusCode >= http.StatusInternalServerError {
		log.Println(message, err, statusCode)
	}
	if message == "" {
		message = err.Error()
	}
	return status.Errorf(codes.Code(statusCode), "%s", message)
}
