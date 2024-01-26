package productService

import (
	"log/slog"
)

type Product struct {
	log            *slog.Logger
	productStorage ProductStorage
	imageStorage   ImageStorage
}

func New(log *slog.Logger, productStorage ProductStorage, imageStorage ImageStorage) *Product {
	return &Product{log: log, productStorage: productStorage, imageStorage: imageStorage}
}
