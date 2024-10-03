package handlers_test

import (
	"context"
	"os"
	"testing"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/app"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/stretchr/testify/assert"
)

func TestUploadProductImage(t *testing.T) {
	ctx := context.Background()
	client, closer := GetClientFromGRPCServer(t)
	defer closer()
	app, err := app.New()
	if err != nil {
		t.Errorf("error creating app: %v", err)
	}
	//make sure there is a product in the database
	app.DB.Create(&types.Product{
		Title:          "test",
		Description:    "test",
		Price:          10.0,
		AvailableUnits: 10,
		CreatorAdminID: 1,
	})

	uploadImageClient, err := client.UploadProductImage(ctx)
	if err != nil {
		t.Errorf("error uploading product image: %v", err)
	}
	err = uploadImageClient.Send(&proto_v1.ProductImageRequest{
		ProductId:    1,
		OriginalName: "test.jpg",
		Extension:    "jpg",
		MimeType:     "image/jpeg",
		FileSize:     1000,
		IsThumbnail:  false,
		Chunk:        []byte("test"),
	})
	if err != nil {
		t.Errorf("error sending upload image request: %v", err)
	}
	pr, err := uploadImageClient.CloseAndRecv()
	if err != nil {
		t.Errorf("error closing and receiving response: %v", err)
	}
	if pr.Path == "" {
		t.Errorf("received empty path")
	}
	assert.FileExists(t, pr.Path)
	assert.Contains(t, pr.Path, "tmp/storage")
	t.Cleanup(func() {
		os.Remove(pr.Path)
	})

}
