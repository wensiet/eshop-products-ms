package redis

import (
	"encoding/json"
	models "eshop-products-ms/internal/models/product"
	"fmt"
	"time"
)

func (r Redis) CacheProduct(product models.Product) error {
	const ttl = time.Minute * 5
	return r.client.Set(fmt.Sprintf("product:%d", product.ID), product, ttl).Err()
}

func (r Redis) GetProduct(productID string) (models.Product, error) {
	key := fmt.Sprintf("product:%s", productID)
	val, err := r.client.Get(key).Result()
	if err != nil {
		return models.Product{}, err
	}
	var product models.Product
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
