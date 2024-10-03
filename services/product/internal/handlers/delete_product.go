package handlers

import (
	"context"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ProductServiceServer) DeleteProduct(ctx context.Context, req *proto_v1.DeleteProductRequest) (*emptypb.Empty, error) {
	if err := p.App.DB.First(&types.Product{}, req.GetId()).Error; err != nil {
		return nil, utils.Error("product not found", err, http.StatusNotFound)
	}
	if err := p.App.DB.Delete(&types.Product{}, req.GetId()).Error; err != nil {
		return nil, utils.Error("error deleting product", err, http.StatusInternalServerError)
	}
	return nil, nil
}
