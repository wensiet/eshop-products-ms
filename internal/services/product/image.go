package productService

import (
	"context"
	appError "eshop-products-ms/internal/apperror"
	models "eshop-products-ms/internal/models/product"
	"fmt"
	"log/slog"
)

type ImageStorage interface {
	SaveImage(image []byte, s3Path string, product models.Product, order int) (string, error)
	Images(productID string) ([]models.Image, error)
}

func (p Product) AddImage(ctx context.Context, image []byte, imageName string, productID string) (string, error) {
	const op = "productService.Product.AddImage"

	log := p.log.With(
		slog.String("op", op),
	)

	log.Info("adding image")

	product, err := p.GetProductByID(ctx, productID)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	images, err := p.imageStorage.Images(productID)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	s3Path := fmt.Sprintf("%s_%s_%d.png", imageName, productID, len(images)+1)

	filename, err := p.imageStorage.SaveImage(image, s3Path, product, len(images)+1)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.With("image", filename).Info("image added")

	return filename, nil
}

func (p Product) GetImages(productID string) ([]string, error) {
	const op = "productService.Product.GetImages"

	log := p.log.With(
		"op", op,
	)

	log.Info("getting images")

	_, err := p.productStorage.Product(productID)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	images, err := p.imageStorage.Images(productID)
	if err != nil {
		appError.LogIfNotApp(err, log)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var filenames []string
	for _, image := range images {
		filenames = append(filenames, image.S3Path)
	}

	return filenames, nil
}
