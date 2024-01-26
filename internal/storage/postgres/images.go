package postgres

import models "eshop-products-ms/internal/models/product"

func (s Storage) SaveImage(s3Path string, product models.Product, order int) error {
	image := models.Image{S3Path: s3Path, Product: product, Order: order}
	return s.DB.Save(&image).Error
}

func (s Storage) Images(productID string) ([]models.Image, error) {
	var images []models.Image
	err := s.DB.Where("product_id = ?", productID).Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}
