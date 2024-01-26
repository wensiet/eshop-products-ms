package productsAPI

import (
	productv1 "eshop-products-ms/gen/go/products"
	productService "eshop-products-ms/internal/services/product"
	userService "eshop-products-ms/internal/services/user"
	"google.golang.org/grpc"
)

type productsAPI struct {
	productv1.UnimplementedProductServServer
	productService productService.Product
	userService    userService.User
}

func Register(serv *grpc.Server, products productService.Product, user userService.User) {
	productv1.RegisterProductServServer(serv, &productsAPI{
		productService: products,
		userService:    user,
	})
}

//func New(productService productService.Product, userService userService.User) *productsAPI {
//	return &productsAPI{productService: productService, userService: userService}
//}
//
//func MustRun(product productService.Product, port string) {
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		panic(err)
//	}
//
//	s := grpc.NewServer()
//
//	serv := New(
//		product,
//		userService.User{},
//	)
//
//	productv1.RegisterProductServServer(s, serv)
//
//	err = s.Serve(lis)
//	if err != nil {
//		panic(err)
//	}
//}
