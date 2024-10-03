package handlers_test

import (
	"context"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/app"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProductImage(t *testing.T) {
	client, closer := GetClientFromGRPCServer(t)
	defer closer()
	ctx := context.Background()
	app, err := app.New()
	if err != nil {
		t.Errorf("error creating app: %v", err)
	}
	t.Run("Test Product Image Deleted From DB", func(t *testing.T) {

		productImage := &types.ProductImage{
			ProductID:    1,
			OriginalName: "test",
			StoredName:   "test",
			Extension:    "png",
			MimeType:     "image/png",
			FileSize:     100,
			IsThumbnail:  false,
		}
		// create product Image
		err := app.DB.Model(productImage).Create(productImage).Error
		if err != nil {
			t.Errorf("error creating product image: %v", err)
		}
		// Get product images count before delete
		var countBeforeDelete int64
		err = app.DB.Model(&types.ProductImage{}).Count(&countBeforeDelete).Error
		if err != nil {
			t.Errorf("error getting images count: %v", err)
		}
		// Delete product image
		_, err = client.DeleteProductImage(ctx, &proto_v1.DeleteProductImageRequest{
			Id: uint32(productImage.ID),
		})
		if err != nil {
			t.Errorf("error deleting product image: %v", err)
		}
		// Get product images count after delete
		var countAfterDelete int64
		err = app.DB.Model(&productImage).Count(&countAfterDelete).Error
		if err != nil {
			t.Errorf("error getting images count: %v", err)
		}
		// Assert product images count before delete - 1 == product images count after delete
		assert.Equal(t, countBeforeDelete-1, countAfterDelete)
	})
}
