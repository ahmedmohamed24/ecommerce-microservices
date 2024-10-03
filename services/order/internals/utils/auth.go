package utils

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetAuthTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(http.StatusUnauthorized, "please, provide the authorization token")
	}
	accessToken := ""
	if authToken, ok := md["authorization"]; ok {
		accessToken = authToken[0]
	} else {
		return "", status.Error(http.StatusUnauthorized, "please, provide the authorization token")
	}
	token := strings.Split(accessToken, "Bearer ")
	if len(token) != 2 {
		return "", status.Error(http.StatusUnauthorized, "please, provide the authorization token")
	}
	return token[1], nil

}
