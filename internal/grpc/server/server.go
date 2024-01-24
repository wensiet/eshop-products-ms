package grpcServer

import (
	"context"
	productv1 "eshop-products-ms/gen/go/products"
	productService "eshop-products-ms/internal/services/product"
	userService "eshop-products-ms/internal/services/user"
	"eshop-products-ms/internal/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"strconv"
)

func MustRun(logger *slog.Logger, storage postgres.Storage) {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	productv1.RegisterProductServServer(s, &GRPCServer{
		productService: *productService.New(logger, storage, storage),
		userService:    userService.User{},
		log:            logger,
	})

	log.Println("Listening on port :9090")
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

type GRPCServer struct {
	productv1.UnimplementedProductServServer
	productService productService.Product
	userService    userService.User
	log            *slog.Logger
}

func (s *GRPCServer) GetProduct(ctx context.Context, in *productv1.GetProductRequest) (*productv1.Product, error) {
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

func (s *GRPCServer) GetProducts(ctx context.Context, in *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error) {
	return nil, nil
}

func (s *GRPCServer) CreateProduct(ctx context.Context, in *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {
	res, err := s.productService.CreateProduct(in.GetTitle(), float64(in.GetPrice()), int(in.GetQuantity()), in.GetDescription())
	if err != nil {
		return nil, err
	}
	return &productv1.CreateProductResponse{
		Id: strconv.Itoa(int(res)),
	}, nil
}

func (s *GRPCServer) UpdateProduct(ctx context.Context, in *productv1.UpdateProductRequest) (*productv1.Product, error) {
	return nil, nil
}

func (s *GRPCServer) DeleteProduct(ctx context.Context, in *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	return nil, nil
}

//func authorizationInterceptor(logger *slog.Logger, userService userService.User) grpc.UnaryServerInterceptor {
//	return func(ctx context.Context,
//		req interface{},
//		info *grpc.UnaryServerInfo,
//		handler grpc.UnaryHandler) (interface{}, error) {
//		const op = "grpcServer.authorizationInterceptor"
//
//		md, ok := metadata.FromIncomingContext(ctx)
//		if !ok {
//			return nil, fmt.Errorf("%s: %w", op, appError.InvalidMetadata)
//		}
//		authHeader := md.Get("authorization")
//
//		if len(authHeader) == 0 {
//			return nil, fmt.Errorf("%s: %w", op, appError.Unauthorized)
//		}
//
//		_, err :=
//
//		return handler(ctx, req)
//
//	}
//}
