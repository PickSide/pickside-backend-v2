package s3client

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	bucketName = "pickside-image-storage"
)

func GetS3Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error loading AWS configuration: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return client, nil
}