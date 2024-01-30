package minio

import (
	"context"
	"eshop-products-ms/internal/config"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	client     *minio.Client
	bucketName string
}

func createBucketIfNotExists(client *minio.Client, bucketName string) error {
	err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "eu-central-1"})
	if err != nil {
		exists, errBucketExists := client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			return nil
		}
		return err
	}
	return nil
}

func New() (*MinIO, error) {
	conf := config.Get()
	dsn := fmt.Sprintf("%s:%d", conf.MinIO.Host, conf.MinIO.Port)
	client, err := minio.New(dsn, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.MinIO.AccessKey, conf.MinIO.SecretKey, ""),
		Secure: conf.MinIO.SSLMode,
	})
	if err != nil {
		return nil, err
	}

	err = createBucketIfNotExists(client, conf.MinIO.Bucket)
	if err != nil {
		return nil, err
	}

	return &MinIO{client: client, bucketName: conf.MinIO.Bucket}, nil
}
