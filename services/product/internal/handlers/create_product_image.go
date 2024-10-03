package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/types"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/utils"
	proto_v1 "github.com/ahmedmohamed24/ecommerce-microservices/product/protos/v1/gen"
	"github.com/go-playground/validator/v10"
)

func (p *ProductServiceServer) UploadProductImage(stream proto_v1.ProductService_UploadProductImageServer) error {

	var fileName string
	var file *os.File
	var productImage = &types.ProductImage{}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			if err := p.App.DB.Create(productImage).Error; err != nil {
				return utils.Error("", err, http.StatusBadRequest)
			}
			err := stream.SendAndClose(&proto_v1.ProductImageResponse{
				Path:        fileName,
				IsThumbnail: req.GetIsThumbnail(),
			})
			if err != nil {
				return utils.Error("error sending response", err, http.StatusBadRequest)
			}
			return nil
		}
		if err != nil {
			return utils.Error("error receiving chunk", err, http.StatusBadRequest)
		}

		if file == nil {
			fileName = fileName + fmt.Sprintf("%v/%v-%s", p.App.Config.Storage.Path, time.Now().Unix(), req.GetOriginalName())
			productImage.ProductID = uint(req.GetProductId())
			productImage.OriginalName = req.GetOriginalName()
			productImage.StoredName = fileName
			productImage.Extension = req.GetExtension()
			productImage.MimeType = req.GetMimeType()
			productImage.FileSize = int64(req.GetFileSize())
			productImage.IsThumbnail = req.GetIsThumbnail()
			validate := validator.New()
			if err := validate.Struct(productImage); err != nil {
				return utils.Error("validation error", err, http.StatusBadRequest)
			}

			if err := os.MkdirAll(p.App.Config.Storage.Path, os.ModePerm); err != nil {
				return utils.Error("error while uploading file", err, http.StatusInternalServerError)
			}
			file, err = os.Create(fileName)
			if err != nil {
				return utils.Error("error while uploading file", err, http.StatusInternalServerError)
			}
		}

		_, err = file.Write(req.GetChunk())
		if err != nil {
			return utils.Error("error while uploading file", err, http.StatusInternalServerError)
		}
	}

}
