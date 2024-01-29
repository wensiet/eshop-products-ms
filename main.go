package main

import (
	grpcApp "eshop-products-ms/internal/app/grpc"
	"eshop-products-ms/internal/config"
	"eshop-products-ms/internal/migrations"
	productService2 "eshop-products-ms/internal/services/product"
	userService2 "eshop-products-ms/internal/services/user"
	"eshop-products-ms/internal/storage/postgres"
	"github.com/wensiet/logmod"
	"log/slog"
)

var storage *postgres.Storage
var logger *slog.Logger
var productService productService2.Product
var userService userService2.User
var conf config.Config

func init() {
	conf = config.Get()

	logger = logmod.New(logmod.Options{
		Env:     conf.Env,
		Service: "eshop-products-ms",
		Loki: struct {
			Host string
			Port int
		}{
			Host: conf.Loki.Host,
			Port: conf.Loki.Port,
		},
	})

	var err error
	storage, err = postgres.NewConnection()
	if err != nil {
		panic(err)
	}

	productService = *productService2.New(logger, storage, storage)

	userService = *userService2.New(logger, storage)

	migrations.MustMigrate(*storage)
}

func main() {
	server := grpcApp.New(logger, productService, userService, conf.GRPC.Port)
	server.MustRun()
}
