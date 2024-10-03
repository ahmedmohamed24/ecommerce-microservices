package handlers

import (
	"context"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/order/internals/types"
	order_protos_v1 "github.com/ahmedmohamed24/ecommerce-microservices/order/protos/v1/gen"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOrderDetails(ctx context.Context, req *order_protos_v1.GetOrderDetailsRequest) (*order_protos_v1.GetOrderDetailsResponse, error) {
	var order = &types.Order{}
	if err := s.DB.Preload("Status").First(order, req.GetOrderId()).Error; err != nil {
		return nil, status.Error(http.StatusNotFound, "order not found")
	}
	var productOrders []*types.ProductOrder
	if err := s.DB.Where("order_id = ?", order.ID).Find(&productOrders).Error; err != nil {
		return nil, status.Error(http.StatusNotFound, "order products not found")
	}
	var orderProductResponses = []*order_protos_v1.OrderProductResponse{}
	for _, po := range productOrders {
		pr, err := s.ProductServiceClient.GetProduct(ctx, &order_protos_v1.GetProductRequest{Id: uint32(po.ProductID)})
		if err != nil {
			return nil, status.Errorf(http.StatusNotFound, "product with ID %v not found", po.ProductID)
		}
		orderProductResponses = append(orderProductResponses, &order_protos_v1.OrderProductResponse{
			Id:       uint32(po.ProductID),
			Quantity: uint32(po.Quantity),
			Price:    uint32(pr.Price),
			Name:     pr.GetTitle(),
		})

	}

	return &order_protos_v1.GetOrderDetailsResponse{
		OrderId:       uint32(order.ID),
		TotalPrice:    float32(order.TotalPrice),
		StatusId:      uint32(order.StatusID),
		Status:        order.Status.Name,
		OrderProducts: orderProductResponses,
		CreatedAt:     order.CreatedAt.Unix(),
		UpdatedAt:     order.UpdatedAt.Unix(),
	}, nil

}
