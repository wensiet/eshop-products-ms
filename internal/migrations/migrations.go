package migrations

import (
	models "eshop-products-ms/internal/models/product"
	"eshop-products-ms/internal/storage/postgres"
)

func MustMigrate(storage postgres.Storage) {
	err := storage.DB.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err)
	}
	err = storage.DB.AutoMigrate(&models.Image{})
	if err != nil {
		panic(err)
	}
}
