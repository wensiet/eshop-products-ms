package productService

import (
	models "eshop-products-ms/internal/models/product"
	"fmt"
	"log/slog"
)

type ImageStorage interface {
	SaveImage(s3Path string, product models.Product, order int) error
	Images(productID string) ([]models.Image, error)
}

func (p Product) AddImage(s3Path string, productID string) error {
	const op = "productService.Product.UpdateProduct"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("adding image")

	product, err := p.GetProductByID(productID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	images, err := p.imageStorage.Images(productID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = p.imageStorage.SaveImage(s3Path, product, len(images)+1)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
