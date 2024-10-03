package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/order/internals/types"
	order_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type orderProductRequest struct {
	ProductID uint `json:"product_id" validate:"required,number,min=1"`
	Quantity  uint `json:"quantity" validate:"required,number,min=1"`
}
type orderRequest struct {
	Products []orderProductRequest `json:"order_products" validate:"required,dive"`
}

func (s *Server) CreateOrder(ctx context.Context, req *order_protos_v1.CreateOrderRequest) (*order_protos_v1.CreateOrderResponse, error) {
	// authorize the request
	customer, err := s.GetAuthCustomerFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// validate the products and the products available units
	orderRequest := &orderRequest{}
	paymentProducts := []*order_protos_v1.Product{}
	var totalPrice float32
	for _, p := range req.OrderProducts {
		pResp, err := s.ProductServiceClient.GetProduct(ctx, &order_protos_v1.GetProductRequest{Id: p.ProductId})
		if err != nil {
			fmt.Println(err)
			return nil, status.Errorf(http.StatusNotFound, "product with ID %v not found", p.ProductId)
		}
		if pResp.AvailableUnits < p.Quantity {
			return nil, status.Errorf(http.StatusBadRequest, "product with ID %v has only %v available units", p.ProductId, pResp.AvailableUnits)
		}
		totalPrice += float32(p.Quantity) * pResp.Price * 100
		orderRequest.Products = append(orderRequest.Products, orderProductRequest{
			ProductID: uint(p.ProductId),
			Quantity:  uint(p.Quantity),
		})
		paymentProducts = append(paymentProducts, &order_protos_v1.Product{
			Quantity:    uint32(p.GetQuantity()),
			Amount:      pResp.GetPrice(),
			Name:        pResp.GetTitle(),
			Description: pResp.GetDescription(),
		})
	}
	validate := validator.New()
	err = validate.Struct(orderRequest)
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, "invalid request")
	}
	order := &types.Order{
		CustomerID: uint(customer.Id),
		TotalPrice: float64(totalPrice),
		StatusID:   types.OrderStatusPending,
	}

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		err = s.DB.Create(order).Error
		if err != nil {
			return err
		}
		for _, p := range req.OrderProducts {
			err = s.DB.Create(&types.ProductOrder{
				OrderID:   uint(order.ID),
				ProductID: uint(p.ProductId),
				Quantity:  uint(p.Quantity),
			}).Error
			if err != nil {
				return err
			}
			_, err = s.ProductServiceClient.UpdateProductQuantity(ctx, &order_protos_v1.UpdateProductQuantityRequest{
				ProductId: p.GetProductId(),
				Quantity:  p.GetQuantity(),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	paymentLinkResp, err := s.PaymentServiceClient.GeneratePaymentLink(ctx, &order_protos_v1.GeneratePaymentLinkRequest{
		OrderId:  fmt.Sprintf("%v", order.ID),
		Products: paymentProducts,
	})
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	return &order_protos_v1.CreateOrderResponse{
		OrderId:     uint32(order.ID),
		PaymentLink: paymentLinkResp.GetPaymentLink(),
	}, nil

}
