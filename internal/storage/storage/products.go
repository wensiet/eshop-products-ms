package storage

import (
	appError "eshop-products-ms/internal/apperror"
	"eshop-products-ms/internal/config"
	models "eshop-products-ms/internal/models/product"
	"github.com/google/uuid"
	"time"
)

func (s Storage) SaveProduct(title string, price float64, quantity int, description string) (uint, error) {
	product := models.Product{Title: title, Price: price, Quantity: quantity, Description: description}
	err := s.Postgres.DB.Save(&product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (s Storage) Product(id string) (models.Product, error) {
	cached, err := s.Redis.GetProduct(id)
	if err == nil {
		return cached, nil
	}
	var product models.Product
	err = s.Postgres.DB.First(&product, id).Error
	if err != nil {
		return models.Product{}, err
	}
	cacheErr := s.Redis.CacheProduct(product)
	if cacheErr != nil {
		appError.LogIfNotApp(cacheErr, config.GetLogger())
	}
	return product, nil
}

func (s Storage) ManyProducts(limit, offset int) ([]models.Product, error) {
	var products []models.Product
	err := s.Postgres.DB.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s Storage) UpdateProduct(product *models.Product) error {
	err := s.Postgres.DB.Save(&product).Error
	if err != nil {
		return err
	}
	_ = s.Redis.CacheProduct(*product)
	return nil
}

func (s Storage) BeginProductUpdateTransaction(product *models.Product) (string, time.Time, error) {
	transactionID := uuid.New().String()

	tx := s.Postgres.DB.Begin()

	err := tx.Save(&product).Error
	if err != nil {
		return "", time.Time{}, err
	}

	expiration := time.Now().Add(time.Second * 5)
	s.TransactionMap[transactionID] = runtimeTransaction{TX: tx, Expires: expiration}

	go endTransactionOnTimeout(transactionID, s.TransactionMap[transactionID])

	return transactionID, expiration, nil
}

func (s Storage) EndProductUpdateTransaction(transactionID string, success bool) error {
	transaction := s.TransactionMap[transactionID]

	s.TransactionMap[transactionID] = runtimeTransaction{}

	if transaction.Expires.Before(time.Now()) {
		return appError.ExpiredTransaction
	}
	if success {
		return transaction.TX.Commit().Error
	}
	return transaction.TX.Rollback().Error
}

func endTransactionOnTimeout(transactionID string, transaction runtimeTransaction) {
	defer func() {
		logger := config.GetLogger()

		err := transaction.TX.Rollback().Error
		if err == nil {
			logger.With("op", "storage.endTransactionOnTimeout").With("transactionID", transactionID).Warn("transaction expired")
		}
	}()
	time.Sleep(time.Until(transaction.Expires))
}
