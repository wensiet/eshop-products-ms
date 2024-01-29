package postgres

import (
	models "eshop-products-ms/internal/models/product"
)

func (s Storage) SaveProduct(title string, price float64, quantity int, description string) (uint, error) {
	product := models.Product{Title: title, Price: price, Quantity: quantity, Description: description}
	err := s.DB.Save(&product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (s Storage) Product(id string) (models.Product, error) {
	cached, err := s.Redis.GetProduct(id)
	if err == nil {
		return cached, nil
	}
	var product models.Product
	err = s.DB.First(&product, id).Error
	if err != nil {
		return models.Product{}, err
	}
	err = s.Redis.CacheProduct(product)
	return product, nil
}

func (s Storage) ManyProducts(limit, offset int) ([]models.Product, error) {
	var products []models.Product
	err := s.DB.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
