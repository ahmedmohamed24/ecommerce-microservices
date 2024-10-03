package handlers

import (
	"context"
	"net/http"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
)

func (p *ProductServiceServer) ListProducts(ctx context.Context, req *proto_v1.ListProductsRequest) (*proto_v1.ListProductsResponse, error) {
	var products = []types.Product{}
	if err := p.App.DB.Limit(int(req.GetLimit())).Offset(int(req.GetPage()-1) * int(req.GetLimit())).Preload("ProductImages").Find(&products).Error; err != nil {
		return nil, utils.Error("error getting products", err, http.StatusInternalServerError)
	}
	var totalCount int64
	if err := p.App.DB.Model(&types.Product{}).Count(&totalCount).Error; err != nil {
		return nil, utils.Error("error getting products", err, http.StatusInternalServerError)
	}
	productResponse := []*proto_v1.GetProductResponse{}
	for _, product := range products {
		productImages := []*proto_v1.ProductImageResponse{}
		for _, productImage := range product.ProductImages {
			productImages = append(productImages, &proto_v1.ProductImageResponse{
				Id:          uint32(productImage.ID),
				Path:        productImage.StoredName,
				IsThumbnail: productImage.IsThumbnail,
			})
		}
		productResponse = append(productResponse, &proto_v1.GetProductResponse{
			Id:             uint32(product.ID),
			Title:          product.Title,
			Price:          float32(product.Price),
			Description:    product.Description,
			AvailableUnits: uint32(product.AvailableUnits),
			CreatedAt:      product.CreatedAt.Unix(),
			UpdatedAt:      product.UpdatedAt.Unix(),
			Images:         productImages,
		})

	}
	return &proto_v1.ListProductsResponse{
		Products: productResponse,
		Total:    uint32(totalCount),
		Page:     req.GetPage(),
	}, nil
}
