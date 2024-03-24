package s3client

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFile(path string) (io.ReadCloser, *string, error) {
	client, err := GetS3Client()
	if err != nil {
		return nil, nil, err
	}

	result, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, nil, err
	}

	return result.Body, result.ContentType, nil
}
