package storage

import (
	"cloud.google.com/go/storage"
	"context"
)

func NewBucketHandler(ctx context.Context, bucket string) (*storage.BucketHandle, error) {
	client, err := storage.NewClient(ctx)

	if err != nil {
		return nil, err
	}
	return client.Bucket(bucket), nil
}
