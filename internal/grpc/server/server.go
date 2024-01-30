package productsAPI

import (
	imagesv1 "eshop-products-ms/gen/go/images"
	productv1 "eshop-products-ms/gen/go/products"
	productService "eshop-products-ms/internal/services/product"
	userService "eshop-products-ms/internal/services/user"
	"google.golang.org/grpc"
)

type productsAPI struct {
	productv1.UnimplementedProductServServer
	imagesv1.UnimplementedImagesServer
	productService productService.Product
	userService    userService.User
}

func Register(serv *grpc.Server, products productService.Product, user userService.User) {
	productsAPI := &productsAPI{
		productService: products,
		userService:    user,
	}

	productv1.RegisterProductServServer(serv, productsAPI)
	imagesv1.RegisterImagesServer(serv, productsAPI)
}
