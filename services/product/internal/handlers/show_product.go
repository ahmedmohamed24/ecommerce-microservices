package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
)

func (p *ProductServiceServer) GetProduct(ctx context.Context, req *proto_v1.GetProductRequest) (*proto_v1.GetProductResponse, error) {
	product := &types.Product{}
	if err := p.App.DB.Preload("ProductImages").First(product, req.GetId()).Error; err != nil {
		return nil, utils.Error("product not found", err, http.StatusNotFound)
	}
	productImages := []*proto_v1.ProductImageResponse{}
	for _, productImage := range product.ProductImages {
		host := fmt.Sprintf("%v:%v/", p.App.Config.FileServer.Host, p.App.Config.FileServer.Port)
		productImages = append(productImages, &proto_v1.ProductImageResponse{
			Id:          uint32(productImage.ID),
			Path:        host + productImage.StoredName,
			IsThumbnail: productImage.IsThumbnail,
		})
	}

	return &proto_v1.GetProductResponse{
		Id:             uint32(product.ID),
		Title:          product.Title,
		Description:    product.Description,
		Price:          float32(product.Price),
		AvailableUnits: uint32(product.AvailableUnits),
		Images:         productImages,
		CreatedAt:      product.CreatedAt.Unix(),
		UpdatedAt:      product.UpdatedAt.Unix(),
	}, nil
}
