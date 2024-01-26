package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
}

func (p Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
}

func (p Product) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

type Image struct {
	gorm.Model
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint    `json:"product_id"`
	Order     int     `json:"order"`
	S3Path    string  `json:"s3"`
}
