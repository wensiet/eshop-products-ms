package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
)

func (m *MinIO) UploadImage(image []byte, name string) (string, error) {
	info, err := m.client.PutObject(
		context.Background(),
		m.bucketName, name,
		bytes.NewReader(image),
		int64(len(image)),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return "", err
	}

	return info.Key, nil
}

func (m *MinIO) CancelUpload(name string) error {
	return m.client.RemoveObject(context.Background(), m.bucketName, name, minio.RemoveObjectOptions{})
}
