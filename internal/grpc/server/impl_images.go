package productsAPI

import (
	"context"
	"errors"
	imagesv1 "eshop-products-ms/gen/go/images"
	appError "eshop-products-ms/internal/apperror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *productsAPI) UploadImage(ctx context.Context, in *imagesv1.UploadImageRequest) (*imagesv1.Empty, error) {
	_, err := s.productService.AddImage(ctx, in.Image, in.Name, in.ProductId)
	if err != nil {
		if errors.Is(err, appError.ProductNotFound) {
			return nil, status.Error(codes.NotFound, "product with such id not found")
		}
		return nil, status.Errorf(codes.Internal, "internal error occurred")
	}
	return &imagesv1.Empty{}, nil
}

func (s *productsAPI) GetProductImages(ctx context.Context, in *imagesv1.GetProductImagesRequest) (*imagesv1.GetProductImagesResponse, error) {
	images, err := s.productService.GetImages(in.ProductId)
	if err != nil {
		if errors.Is(err, appError.ProductNotFound) {
			return nil, status.Error(codes.NotFound, "product with such id not found")
		}
		return nil, status.Errorf(codes.Internal, "internal error occurred")
	}
	return &imagesv1.GetProductImagesResponse{
		Total:      int32(len(images)),
		ImagePaths: images,
	}, nil
}
