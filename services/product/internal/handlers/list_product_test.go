package handlers_test

import (
	"context"
	"testing"

	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {

	client, closer := GetClientFromGRPCServer(t)
	defer closer()
	ctx := context.Background()
	t.Run("Test List Products", func(t *testing.T) {
		resp, err := client.ListProducts(ctx, &proto_v1.ListProductsRequest{
			Limit: 10,
			Page:  1,
		})
		if err != nil {
			t.Errorf("error listing products: %v", err)
		}
		assert.Equal(t, 1, int(resp.Page))
	})

}
