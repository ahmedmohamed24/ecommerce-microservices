package handlers

import (
	"context"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/go-playground/validator/v10"
)

func (p *ProductServiceServer) CreateProduct(ctx context.Context, req *proto_v1.CreateProductRequest) (*proto_v1.CreateProductResponse, error) {
	product := &types.Product{
		Title:          req.GetTitle(),
		Description:    req.GetDescription(),
		Price:          float64(req.GetPrice()),
		AvailableUnits: int(req.GetAvailableUnits()),
	}
	validate := validator.New()
	err := validate.Struct(product)
	if err != nil {
		return nil, utils.Error(err.Error(), err, http.StatusBadRequest)
	}
	if err := p.App.DB.Create(product).Error; err != nil {
		return nil, utils.Error("", err, http.StatusBadRequest)
	}

	return &proto_v1.CreateProductResponse{
		Id: uint32(product.ID),
	}, nil
}
