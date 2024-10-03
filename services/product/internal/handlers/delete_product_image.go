package handlers

import (
	"context"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ProductServiceServer) DeleteProductImage(ctx context.Context, req *proto_v1.DeleteProductImageRequest) (*emptypb.Empty, error) {
	if err := p.App.DB.First(&types.ProductImage{}, req.GetId()).Error; err != nil {
		return nil, utils.Error("product image not found", err, http.StatusNotFound)
	}
	if err := p.App.DB.Delete(&types.ProductImage{}, req.GetId()).Error; err != nil {
		return nil, utils.Error("error deleting product image", err, http.StatusInternalServerError)
	}
	return nil, nil
}
