package handlers_test

import (
	"context"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/app"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	client, closer := GetClientFromGRPCServer(t)
	defer closer()
	ctx := context.Background()
	app, err := app.New()
	if err != nil {
		t.Errorf("error creating app: %v", err)
	}
	t.Run("Test Product Saved In DB ", func(t *testing.T) {
		product := &types.Product{}

		var preRequestProductsCount, afterRequestProductsCount int64
		if err := app.DB.Model(&product).Where("title = ?", "test").Count(&preRequestProductsCount).Error; err != nil {
			t.Errorf("error getting count of products: %v", err.Error())
		}
		_, err = client.CreateProduct(ctx, &proto_v1.CreateProductRequest{
			Title:          "test",
			Description:    "test",
			Price:          10.0,
			AvailableUnits: 10,
		})
		if err != nil {
			t.Errorf("error creating product: %v", err)
		}
		if err := app.DB.Model(&product).Where("title = ?", "test").Count(&afterRequestProductsCount).Error; err != nil {
			t.Errorf("error getting count of products: %v", err.Error())
		}
		assert.Equal(t, preRequestProductsCount+1, afterRequestProductsCount)

	})

}
