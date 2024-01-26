package postgres

import (
	"eshop-products-ms/internal/config"
	"eshop-products-ms/internal/storage/caching/redis"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB    *gorm.DB
	Redis *redis.Redis
}

func NewConnection() (*Storage, error) {
	conf := config.Get()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Database.Host, conf.Database.User, conf.Database.Pass, conf.Database.Name, conf.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	redisInstance, err := redis.NewConnection()
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db, Redis: redisInstance}, err
}
