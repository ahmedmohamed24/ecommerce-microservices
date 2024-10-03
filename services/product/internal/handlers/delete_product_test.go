package handlers_test

import (
	"context"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/app"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProduct(t *testing.T) {
	client, closer := GetClientFromGRPCServer(t)
	defer closer()
	ctx := context.Background()
	app, err := app.New()
	if err != nil {
		t.Errorf("error creating app: %v", err)
	}

	t.Run("Test Product Deleted From DB", func(t *testing.T) {
		var product = &types.Product{
			CreatorAdminID: 1,
			Title:          "test",
			Description:    "test",
			Price:          10.0,
			AvailableUnits: 10,
		}
		// create product
		err = app.DB.Model(product).Create(product).Error
		if err != nil {
			t.Errorf("error creating product: %v", err)
		}
		// get count of products before delete
		var countBeforeDelete int64
		err = app.DB.Model(&types.Product{}).Count(&countBeforeDelete).Error
		if err != nil {
			t.Errorf("error getting products count: %v", err)
		}
		// delete product
		_, err = client.DeleteProduct(ctx, &proto_v1.DeleteProductRequest{
			Id: uint32(product.ID),
		})
		if err != nil {
			t.Errorf("error deleting product: %v", err)
		}
		// get count of products after delete
		var countAfterDelete int64
		err = app.DB.Model(&product).Count(&countAfterDelete).Error
		if err != nil {
			t.Errorf("error getting products count: %v", err)
		}
		// assert count of products before delete - 1 == count of products after delete
		assert.Equal(t, countBeforeDelete-1, countAfterDelete)

	})
}
