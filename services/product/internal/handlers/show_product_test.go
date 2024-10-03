package handlers_test

import (
	"context"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/app"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/stretchr/testify/assert"
)

func TestShowProduct(t *testing.T) {
	client, closer := GetClientFromGRPCServer(t)
	defer closer()
	ctx := context.Background()
	app, err := app.New()
	if err != nil {
		t.Errorf("error creating app: %v", err)
	}
	t.Run("Test Product Saved In DB ", func(t *testing.T) {
		// create product in DB
		product := &types.Product{
			CreatorAdminID: 1,
			Title:          "test",
			Description:    "test",
			Price:          10.0,
			AvailableUnits: 10,
		}
		if err := app.DB.Model(product).Create(product).Error; err != nil {
			t.Errorf("error creating product: %v", err)
		}
		// get product from GRPC
		resp, err := client.GetProduct(ctx, &proto_v1.GetProductRequest{
			Id: uint32(product.ID),
		})
		if err != nil {
			t.Errorf("error getting product: %v", err)
		}
		// assert product from DB == product from response
		assert.Equal(t, product.Title, resp.Title)
		assert.NotEmpty(t, resp.CreatedAt)

	})
	t.Run("Test Product Not Found", func(t *testing.T) {
		lastProduct := &types.Product{}
		app.DB.Last(lastProduct)
		_, err := client.GetProduct(ctx, &proto_v1.GetProductRequest{
			Id: uint32(lastProduct.ID) + 1,
		})
		if err == nil {
			t.Fatal("Expected error, but got none")
		}
	})
}
