package productService

import (
	appError "eshop-products-ms/internal/apperror"
	models "eshop-products-ms/internal/models/product"
	"fmt"
	"log/slog"
)

type ProductStorage interface {
	SaveProduct(title string, price float64, quantity int, description string) (uint, error)
	Product(id string) (models.Product, error)
}

type ImageStorage interface {
	SaveImage(s3Path string, product models.Product, order int) error
	Images(productID string) ([]models.Image, error)
}

type Product struct {
	log            *slog.Logger
	productStorage ProductStorage
	imageStorage   ImageStorage
}

func New(log *slog.Logger, productStorage ProductStorage, imageStorage ImageStorage) *Product {
	return &Product{log: log, productStorage: productStorage, imageStorage: imageStorage}
}

func (p Product) CreateProduct(title string, price float64, quantity int, description string) (uint, error) {
	const op = "productService.Product.CreateProduct"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("creating product")

	if title == "" {
		return 0, fmt.Errorf("%s: %w", op, appError.InvalidTitle)
	}

	if price <= 0 {
		return 0, fmt.Errorf("%s: %w", op, appError.InvalidPrice)
	}

	if quantity <= 0 {
		return 0, fmt.Errorf("%s: %w", op, appError.InvalidQuantity)
	}

	productID, err := p.productStorage.SaveProduct(title, price, quantity, description)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return productID, nil
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

func (p Product) GetProductByID(id string) (models.Product, error) {
	const op = "productService.Product.GetProductByID"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("getting product")

	product, err := p.productStorage.Product(id)
	if err != nil {
		return models.Product{}, fmt.Errorf("%s: %w", op, appError.ProductNotFound)
	}
	return product, nil
}
