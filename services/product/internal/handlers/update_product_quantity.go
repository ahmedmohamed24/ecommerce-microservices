package handlers

import (
	"context"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ProductServiceServer) UpdateProductQuantity(ctx context.Context, req *proto_v1.UpdateProductQuantityRequest) (*emptypb.Empty, error) {
	var product types.Product
	req.GetProductId()
	if err := p.App.DB.First(&product, req.GetProductId()).Error; err != nil {
		return nil, utils.Error("product not found", err, http.StatusBadRequest)
	}
	if product.AvailableUnits < int(req.GetQuantity()) {
		return nil, utils.Error("not enough units available", nil, http.StatusBadRequest)
	}
	quantity := product.AvailableUnits - int(req.GetQuantity())
	product.AvailableUnits = quantity
	if err := p.App.DB.Save(&product).Error; err != nil {
		return nil, utils.Error("error updating product", err, http.StatusInternalServerError)
	}

	return nil, nil
}
