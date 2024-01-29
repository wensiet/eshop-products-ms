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
	ManyProducts(limit, offset int) ([]models.Product, error)
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
		appError.LogIfNotApp(err, p.log)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return productID, nil
}

func (p Product) GetProductByID(id string) (models.Product, error) {
	const op = "productService.Product.GetProductByID"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("getting product")

	product, err := p.productStorage.Product(id)
	if err != nil {
		appError.LogIfNotApp(err, p.log)
		return models.Product{}, fmt.Errorf("%s: %w", op, appError.ProductNotFound)
	}
	return product, nil
}

func (p Product) GetProductsWithPaging(page int, pageSize ...int) ([]models.Product, error) {
	const op = "productService.Product.GetProductsWithPaging"
	const defaultPageSize = 10

	var pSize int
	if len(pageSize) == 0 {
		pSize = defaultPageSize
	} else {
		if pageSize[0] <= 0 {
			pSize = defaultPageSize
		} else {
			pSize = pageSize[0]
		}
	}

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("getting products")

	products, err := p.productStorage.ManyProducts(pSize, (page-1)*pSize)
	if err != nil {
		appError.LogIfNotApp(err, p.log)
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return products, nil
}
