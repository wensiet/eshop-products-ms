package tests_test

import (
	models "eshop-products-ms/internal/models/product"
	myRedis "eshop-products-ms/internal/storage/caching/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

type RedisTestsFixtures struct {
	ProductFixture struct {
		Product models.Product
		Passed  bool
	}
	RedisClientFixture struct {
		Redis  *myRedis.Redis
		Passed bool
	}
	ProductIsCachedFixture struct {
		ProductID string
		Passed    bool
	}
}

var redisTestsFixtures RedisTestsFixtures

func TestInitProductFixture(t *testing.T) {
	product := models.Product{
		Title:       "Product fixture",
		Price:       1,
		Quantity:    1,
		Description: "This is a product fixture",
	}
	product.ID = 999
	redisTestsFixtures.ProductFixture.Product = product
	redisTestsFixtures.ProductFixture.Passed = true
}

func TestRedisConnection(t *testing.T) {
	conn, err := myRedis.NewConnection()
	assert.NoError(t, err, "Failed to initialize Redis client")

	redisTestsFixtures.RedisClientFixture.Passed = true
	redisTestsFixtures.RedisClientFixture.Redis = conn
}

func TestRedisCacheProduct(t *testing.T) {
	if !redisTestsFixtures.ProductFixture.Passed || !redisTestsFixtures.RedisClientFixture.Passed {
		t.Skip("Skipping Redis cache product test")
	}
	err := redisTestsFixtures.RedisClientFixture.Redis.CacheProduct(redisTestsFixtures.ProductFixture.Product)
	assert.NoError(t, err, "Failed to cache product")
}

func TestRedisGetProduct(t *testing.T) {
	if !redisTestsFixtures.ProductFixture.Passed || !redisTestsFixtures.RedisClientFixture.Passed {
		t.Skip("Skipping Redis get product test")
	}
	product, err := redisTestsFixtures.RedisClientFixture.Redis.GetProduct("999")
	assert.NoError(t, err, "Failed to get product from Redis")
	assert.Equal(t, redisTestsFixtures.ProductFixture.Product, product, "Product from Redis is not equal to the fixture")
}

func TestRedisGetProductFail(t *testing.T) {
	if !redisTestsFixtures.ProductFixture.Passed || !redisTestsFixtures.RedisClientFixture.Passed {
		t.Skip("Skipping Redis get product test")
	}
	product, err := redisTestsFixtures.RedisClientFixture.Redis.GetProduct("998")
	assert.Error(t, err, "Failed to get product from Redis")
	assert.Equal(t, models.Product{}, product, "Product from Redis is not equal to the fixture")
}
