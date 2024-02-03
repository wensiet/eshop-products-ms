package productService

import (
	"context"
	appError "eshop-products-ms/internal/apperror"
	models "eshop-products-ms/internal/models/product"
	"fmt"
	"github.com/getsentry/sentry-go"
	"log/slog"
	"time"
)

type ProductStorage interface {
	SaveProduct(title string, price float64, quantity int, description string) (uint, error)
	Product(id string) (models.Product, error)
	ManyProducts(limit, offset int) ([]models.Product, error)
	UpdateProduct(product *models.Product) error
	BeginProductUpdateTransaction(product *models.Product) (string, time.Time, error)
	EndProductUpdateTransaction(transactionID string, status bool) error
}

func (p Product) CreateProduct(title string, price float64, quantity int, description string) (uint, error) {
	const op = "productService.Product.CreateProduct"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("creating product")

	transaction := sentry.StartTransaction(context.Background(), op)
	defer transaction.Finish()

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

func (p Product) GetProductByID(ctx context.Context, id string) (models.Product, error) {
	const op = "productService.Product.GetProductByID"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("getting product")

	transaction := sentry.StartTransaction(ctx, op)
	defer transaction.Finish()

	product, err := p.productStorage.Product(id)
	if err != nil {

		appError.LogIfNotApp(err, p.log)
		return models.Product{}, fmt.Errorf("%s: %w", op, appError.ProductNotFound)
	}

	return product, nil
}

func (p Product) GetProductsWithPaging(page int, pageSize ...int) ([]models.Product, error) {
	const op = "productService.Product.GetProductsWithPaging"

	transaction := sentry.StartTransaction(context.Background(), op)
	defer transaction.Finish()

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("getting products")

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

	products, err := p.productStorage.ManyProducts(pSize, (page-1)*pSize)
	if err != nil {
		appError.LogIfNotApp(err, p.log)
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return products, nil
}

func (p Product) BeginTransaction(productID string, quantity int32) (string, error) {
	const op = "productService.Product.BeginTransaction"

	transaction := sentry.StartTransaction(context.Background(), op)
	defer transaction.Finish()

	log := p.log.With("op", op)

	log.Info(fmt.Sprintf("begin order transaction, productID: %s, quantity: %d", productID, quantity))

	product, err := p.productStorage.Product(productID)
	if err != nil {
		appError.LogIfNotApp(err, p.log)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	newQuantity := product.Quantity - int(quantity)
	if newQuantity < 0 {
		return "", fmt.Errorf("%s: %w", op, appError.NotEnoughProducts)
	}

	product.Quantity = newQuantity

	trID, expiration, err := p.productStorage.BeginProductUpdateTransaction(&product)
	if err != nil {
		appError.LogIfNotApp(err, p.log)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.With("transactionID", trID).With("expires", expiration).Info("transaction created")

	return trID, nil
}

func (p Product) ApplyTransaction(trID string, success bool) error {
	const op = "productService.Product.ApplyTransaction"

	transaction := sentry.StartTransaction(context.Background(), op)
	defer transaction.Finish()

	log := p.log.With("op", op)

	err := p.productStorage.EndProductUpdateTransaction(trID, success)
	if err != nil {
		log.With("transactionID", trID).Error("failed to apply transaction")
		appError.LogIfNotApp(err, p.log)
		return err
	}

	log.With("transactionID", trID).With("success", success).Info("transaction applied")

	return nil
}
