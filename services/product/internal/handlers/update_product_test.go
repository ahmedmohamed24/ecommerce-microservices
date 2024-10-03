package handlers_test

import (
	"context"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/app"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProduct(t *testing.T) {
	client, closer := GetClientFromGRPCServer(t)
	defer closer()
	ctx := context.Background()
	app, err := app.New()
	if err != nil {
		t.Errorf("error creating app: %v", err)
	}
	t.Run("Test Product Updated In DB", func(t *testing.T) {
		// create product in DB
		var product = &types.Product{
			CreatorAdminID: 1,
			Title:          "test",
			Description:    "test",
			Price:          10.0,
			AvailableUnits: 10,
		}
		err = app.DB.Model(product).Create(product).Error
		if err != nil {
			t.Errorf("error creating product: %v", err)
		}
		// update product with gRPC
		_, err = client.UpdateProduct(ctx, &proto_v1.UpdateProductRequest{
			Id:          uint32(product.ID),
			Title:       "test2",
			Description: "test2",
			Price:       20.0,
		})
		if err != nil {
			t.Errorf("error updating product: %v", err)
		}
		// get product from DB
		var updatedProduct = &types.Product{}
		err = app.DB.Model(updatedProduct).Where("id = ?", product.ID).First(updatedProduct).Error
		if err != nil {
			t.Errorf("error getting product: %v", err)
		}
		// assert product from DB == product from response
		assert.Equal(t, "test2", updatedProduct.Title)
		assert.Equal(t, "test2", updatedProduct.Description)
		assert.Equal(t, 20.0, updatedProduct.Price)

	})
	t.Run("Test Product Not Found", func(t *testing.T) {
		lastProduct := &types.Product{}
		app.DB.Last(lastProduct)
		_, err := client.UpdateProduct(ctx, &proto_v1.UpdateProductRequest{
			Id: uint32(lastProduct.ID) + 1,
		})
		if err == nil {
			t.Fatal("Expected Not found error, but got none")
		}
	})
}
