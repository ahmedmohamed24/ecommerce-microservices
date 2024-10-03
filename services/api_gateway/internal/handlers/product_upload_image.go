package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ahmedmohamed24/ecommerce-microservices/gateway/config"
	protoV1 "github.com/ahmedmohamed24/ecommerce-microservices/gateway/protos/v1/gen/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UploadProductImage(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()
	cfg, err := config.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client, err := grpc.NewClient(fmt.Sprintf("%v:%v", cfg.ProductService.Host, cfg.ProductService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	clientStream, err := protoV1.NewProductServiceClient(client).UploadProductImage(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer client.Close()
	chunk := make([]byte, 64*1024) // 64 KB
	for {
		chunk = chunk[:cap(chunk)]
		n, err := f.Read(chunk)
		if err != nil && errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		chunk = chunk[:n]
		fmt.Println("Sending chunk", len(chunk))
		err = clientStream.Send(&protoV1.ProductImageRequest{
			Chunk:        chunk,
			OriginalName: header.Filename,
			MimeType:     header.Header.Get("Content-Type"),
			Extension:    filepath.Ext(header.Filename),
			ProductId:    1,
			FileSize:     uint64(header.Size),
		})
	}
	response, err := clientStream.CloseAndRecv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(map[string]string{"url": response.GetPath()})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(b)

}
