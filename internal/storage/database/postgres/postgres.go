package postgres

import (
	"eshop-products-ms/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func New() (*Postgres, error) {
	conf := config.Get()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Database.Host, conf.Database.User, conf.Database.Pass, conf.Database.Name, conf.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Postgres{DB: db}, err
}
