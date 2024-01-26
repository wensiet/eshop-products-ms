package main

import (
	grpcApp "eshop-products-ms/internal/app/grpc"
	"eshop-products-ms/internal/config"
	"eshop-products-ms/internal/migrations"
	productService2 "eshop-products-ms/internal/services/product"
	userService2 "eshop-products-ms/internal/services/user"
	"eshop-products-ms/internal/storage/postgres"
	"log/slog"
	"os"
)

var storage *postgres.Storage
var logger *slog.Logger
var productService productService2.Product
var userService userService2.User
var conf config.Config

func init() {
	conf = config.Get()

	switch conf.Env {
	case "local":
		logger = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "production":
		logger = slog.New(slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case "test":
		logger = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

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
