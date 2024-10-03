package handlers

import (
	"context"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ProductServiceServer) UpdateProduct(ctx context.Context, req *proto_v1.UpdateProductRequest) (*emptypb.Empty, error) {
	var product types.Product
	if err := p.App.DB.First(&product, req.GetId()).Error; err != nil {
		return nil, utils.Error("product not found", err, http.StatusBadRequest)
	}
	product.Title = req.GetTitle()
	product.Description = req.GetDescription()
	product.Price = float64(req.GetPrice())
	product.AvailableUnits = int(req.GetAvailableUnits())
	if err := p.App.DB.Save(&product).Error; err != nil {
		return nil, utils.Error("error updating product", err, http.StatusInternalServerError)
	}

	return nil, nil
}
