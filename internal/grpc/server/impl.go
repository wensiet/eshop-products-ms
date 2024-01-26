package productsAPI

import (
	"context"
	productv1 "eshop-products-ms/gen/go/products"
	"strconv"
)

func (s *productsAPI) GetProduct(ctx context.Context, in *productv1.GetProductRequest) (*productv1.Product, error) {
	res, err := s.productService.GetProductByID(in.GetId())
	if err != nil {
		return nil, err
	}
	return &productv1.Product{
		Id:          strconv.Itoa(int(res.ID)),
		Title:       res.Title,
		Description: res.Description,
		Price:       float32(res.Price),
		Quantity:    int32(res.Quantity),
	}, nil
}

func (s *productsAPI) GetProducts(ctx context.Context, in *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error) {
	products, err := s.productService.GetProductsWithPaging(int(in.Page), int(in.Limit))
	if err != nil {
		return nil, err
	}
	var productsConverted []*productv1.Product
	for _, product := range products {
		productsConverted = append(productsConverted, &productv1.Product{
			Id:          strconv.Itoa(int(product.ID)),
			Title:       product.Title,
			Description: product.Description,
			Price:       float32(product.Price),
			Quantity:    int32(product.Quantity),
		})
	}
	return &productv1.GetProductsResponse{
		Products: productsConverted,
		Total:    int32(len(products)),
	}, nil
}

func (s *productsAPI) CreateProduct(ctx context.Context, in *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {
	res, err := s.productService.CreateProduct(in.GetTitle(), float64(in.GetPrice()), int(in.GetQuantity()), in.GetDescription())
	if err != nil {
		return nil, err
	}
	return &productv1.CreateProductResponse{
		Id: strconv.Itoa(int(res)),
	}, nil
}

func (s *productsAPI) UpdateProduct(ctx context.Context, in *productv1.UpdateProductRequest) (*productv1.Product, error) {
	return nil, nil
}

func (s *productsAPI) DeleteProduct(ctx context.Context, in *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	return nil, nil
}
