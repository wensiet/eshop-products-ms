package storage

import (
	"eshop-products-ms/internal/storage/caching/redis"
	"eshop-products-ms/internal/storage/database/postgres"
	"eshop-products-ms/internal/storage/objects/minio"
	"gorm.io/gorm"
	"time"
)

type runtimeTransaction struct {
	TX      *gorm.DB
	Expires time.Time
}

type Storage struct {
	Postgres       *postgres.Postgres
	Redis          *redis.Redis
	MinIO          *minio.MinIO
	TransactionMap map[string]runtimeTransaction
}

// New Create Storage level instance
func New() (*Storage, error) {
	postgresInstance, err := postgres.New()
	if err != nil {
		return nil, err
	}
	redisInstance, err := redis.New()
	if err != nil {
		return nil, err
	}
	minioInstance, err := minio.New()
	return &Storage{
		Postgres:       postgresInstance,
		Redis:          redisInstance,
		MinIO:          minioInstance,
		TransactionMap: map[string]runtimeTransaction{},
	}, err
}
