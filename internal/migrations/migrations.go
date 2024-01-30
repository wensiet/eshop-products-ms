package migrations

import (
	models "eshop-products-ms/internal/models/product"
	"eshop-products-ms/internal/storage/storage"
)

func MustMigrate(storage storage.Storage) {
	err := storage.Postgres.DB.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err)
	}
	err = storage.Postgres.DB.AutoMigrate(&models.Image{})
	if err != nil {
		panic(err)
	}
}
