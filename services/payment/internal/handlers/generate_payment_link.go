package handlers

import (
	"context"
	"log"

	payment_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/payment/protos/v1/gen"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

func (s *GrpcServer) GeneratePaymentLink(ctx context.Context, req *payment_protos_v1.GeneratePaymentLinkRequest) (*payment_protos_v1.GeneratePaymentLinkResponse, error) {
	stripe.Key = s.Config.Services.Stripe.Key
	var products = make([]*stripe.CheckoutSessionLineItemParams, 0, len(req.Products))
	for _, p := range req.Products {
		products = append(products, &stripe.CheckoutSessionLineItemParams{
			Quantity: stripe.Int64(int64(p.GetQuantity())),
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency:   stripe.String("USD"),
				UnitAmount: stripe.Int64(int64(p.GetAmount())),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Description: stripe.String(p.Description),
					Name:        stripe.String(p.Name),
				},
			},
		})
	}
	params := &stripe.CheckoutSessionParams{
		LineItems: products,
		Metadata: map[string]string{
			"order_id": req.GetOrderId(),
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(s.Config.Services.Stripe.SuccessURL + "?session={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(s.Config.Services.Stripe.CancelURL + "?session={CHECKOUT_SESSION_ID}"),
	}
	session, err := session.New(params)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &payment_protos_v1.GeneratePaymentLinkResponse{
		PaymentLink: session.URL,
	}, nil

}
