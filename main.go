package main

import (
	grpcApp "eshop-products-ms/internal/app/grpc"
	"eshop-products-ms/internal/config"
	"eshop-products-ms/internal/migrations"
	productService2 "eshop-products-ms/internal/services/product"
	userService2 "eshop-products-ms/internal/services/user"
	"eshop-products-ms/internal/storage/postgres"
	"github.com/getsentry/sentry-go"
	"log/slog"
)

var storage *postgres.Storage
var logger *slog.Logger
var productService productService2.Product
var userService userService2.User
var conf config.Config

func init() {
	conf = config.Get()

	logger = config.GetLogger()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://56d8e39d5f2248010f851a6e23cf3dd2@o4506655030378496.ingest.sentry.io/4506655062753280",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		logger.Error("failed to init sentry %s", err)
	}

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
