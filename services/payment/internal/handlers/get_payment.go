package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/payment/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/payment/internal/utils"
	payment_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/payment/protos/v1/gen"
	"google.golang.org/grpc/status"
)

func (s *GrpcServer) GetPaymentOperation(ctx context.Context, req *payment_protos_v1.GetPaymentOperationRequest) (*payment_protos_v1.GetPaymentOperationResponse, error) {
	payment := types.Payment{}
	if err := s.DB.First(&payment, "order_id = ?", req.GetOrderId()).Error; err != nil {
		return nil, status.Error(http.StatusNotFound, "payment not found")
	}
	jsonStruct, err := utils.ConvertDatatypesJSONToStruct(payment.PaymentInfo)
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "something went wrong, please try again later")
	}
	return &payment_protos_v1.GetPaymentOperationResponse{
		OrderId:       uint32(payment.OrderID),
		PaymentId:     uint32(payment.ID),
		PaidAmount:    float32(payment.PaidAmount),
		CardLast_4:    payment.CardLast4,
		PaymentInfo:   jsonStruct,
		CreatedAt:     payment.CreatedAt.Unix(),
		UpdatedAt:     payment.UpdatedAt.Unix(),
		PaymentMethod: payment.PaymentMethod,
	}, nil
}
