package storage

import (
	models "eshop-products-ms/internal/models/product"
	"github.com/getsentry/sentry-go"
)

func (s Storage) SaveImage(bytesImage []byte, s3Path string, product models.Product, order int) (string, error) {
	uploadImage, err := s.MinIO.UploadImage(bytesImage, s3Path)
	if err != nil {
		return "", err
	}
	image := models.Image{S3Path: s3Path, Product: product, Order: order}
	err = s.Postgres.DB.Save(&image).Error
	if err != nil {
		err2 := s.MinIO.CancelUpload(s3Path)
		if err2 != nil {
			sentry.CaptureException(err2)
		}
		return "", err
	}
	return uploadImage, nil
}

func (s Storage) Images(productID string) ([]models.Image, error) {
	var images []models.Image
	err := s.Postgres.DB.Where("product_id = ?", productID).Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}
