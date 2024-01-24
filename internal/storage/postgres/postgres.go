package postgres

import (
	"eshop-products-ms/internal/config"
	models "eshop-products-ms/internal/models/product"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func (s Storage) SaveProduct(title string, price float64, quantity int, description string) (uint, error) {
	product := models.Product{Title: title, Price: price, Quantity: quantity, Description: description}
	err := s.DB.Save(&product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (s Storage) Product(id string) (models.Product, error) {
	var product models.Product
	err := s.DB.First(&product, id).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

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

func NewConnection() (*Storage, error) {
	conf := config.Get()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Database.Host, conf.Database.User, conf.Database.Pass, conf.Database.Name, conf.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return &Storage{DB: db}, err
}
