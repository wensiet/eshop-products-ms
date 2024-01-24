package main

import (
	grpcServer "eshop-products-ms/internal/grpc/server"
	"eshop-products-ms/internal/storage/postgres"
	"log/slog"
	"os"
)

func init() {
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	storage, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	//migrations.MustMigrate(*storage)
	//service := productService.New(logger, storage, storage)
	////productID, err := service.CreateProduct("Apple", 100, 1, "Fruit")
	////if err != nil {
	////	panic(err)
	////}
	////product, err := service.GetProductByID(strconv.Itoa(int(productID)))
	////if err != nil {
	////	panic(err)
	////}
	////fmt.Println(product)
	//err = service.AddImage("some-path2", "1")
	//if err != nil {
	//	panic(err)
	//}
	grpcServer.MustRun(logger, *storage)
}
