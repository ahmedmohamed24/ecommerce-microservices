package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ahmedmohamed24/ecommerce-microservices/payment/internal/types"
	payment_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/payment/protos/v1/gen"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/status"
	"gorm.io/datatypes"
)

func (s *GrpcServer) SuccessPaymentOperation(ctx context.Context, req *payment_protos_v1.SuccessPaymentOperationRequest) (*httpbody.HttpBody, error) {
	result, order, err := validateOrderPayment(req, s)
	if err != nil {
		return nil, err
	}
	orderInfo, err := json.Marshal(map[string]string{
		"payment_id":     result.PaymentIntent.ID,
		"status":         string(result.Status),
		"session_id":     req.GetSession(),
		"payment_method": string(result.PaymentIntent.PaymentMethod.Type),
	})
	if err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusBadRequest, "something went wrong, please try again later")
	}
	payment := &types.Payment{
		OrderID:       uint(order),
		PaymentMethod: string(result.PaymentIntent.PaymentMethod.Type),
		PaidAmount:    float64(result.AmountTotal),
		CardLast4:     result.PaymentIntent.PaymentMethod.Card.Last4,
		PaymentInfo:   datatypes.JSON(orderInfo),
	}

	if err := s.DB.Create(payment).Error; err != nil {
		log.Println(err)
		return nil, status.Error(http.StatusInternalServerError, "something went wrong, please try again later")
	}
	if err := emitSuccessPaymentEvent(s.AmqpChannel, uint32(order)); err != nil {
		return nil, err
	}

	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Order paid successfully"),
	}, nil

}
func validateOrderPayment(req *payment_protos_v1.SuccessPaymentOperationRequest, s *GrpcServer) (*stripe.CheckoutSession, int, error) {
	if req.GetSession() == "" {
		return nil, 0, status.Error(http.StatusNotFound, "session not found")
	}
	stripe.Key = s.Config.Services.Stripe.Key
	result, err := session.Get(req.GetSession(), &stripe.CheckoutSessionParams{
		Expand: []*string{
			stripe.String("payment_intent.payment_method"),
		},
	})
	if err != nil {
		return nil, 0, status.Error(http.StatusNotFound, "session not found")
	}
	if result.Status != stripe.CheckoutSessionStatusComplete {
		return nil, 0, status.Error(http.StatusBadRequest, "Payment status is "+string(result.Status))

	}
	if _, ok := result.Metadata["order_id"]; !ok {
		return nil, 0, status.Error(http.StatusBadRequest, "order_id not found")
	}
	orderId, _ := result.Metadata["order_id"]
	order, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, 0, status.Error(http.StatusBadRequest, "order_id not valid")
	}
	p := &types.Payment{}
	s.DB.Where("order_id = ?", order).First(p)
	if p.ID != 0 {
		return nil, 0, status.Error(http.StatusBadRequest, "order already paid")
	}
	return result, order, nil
}

func emitSuccessPaymentEvent(amqpChannel *amqp.Channel, orderID uint32) error {
	q, err := amqpChannel.QueueDeclare("success_payment", false, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return status.Error(http.StatusBadRequest, "something went wrong, please try again later")
	}
	b, err := json.Marshal(payment_protos_v1.PaymentDoneSuccessfully{OrderId: orderID})
	if err != nil {
		log.Println(err)
		return status.Error(http.StatusBadRequest, "something went wrong, please try again later")
	}

	err = amqpChannel.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: b})
	if err != nil {
		log.Println(err)
		return status.Error(http.StatusBadRequest, "something went wrong, please try again later")
	}
	return nil
}
