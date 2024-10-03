package handlers

import (
	"context"

	payment_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/payment/protos/v1/gen"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

func (s *GrpcServer) FailPaymentOperation(ctx context.Context, req *payment_protos_v1.FailPaymentOperationRequest) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("order payment canceled"),
	}, nil
}
