package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
}

type Image struct {
	gorm.Model
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint    `json:"product_id"`
	Order     int     `json:"order"`
	S3Path    string  `json:"s3"`
}
