package s3client

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadContentToDirectory(directory string, keyId string, fileName string, fileContent multipart.File) (*string, error) {
	client, err := GetS3Client()
	if err != nil {
		return nil, err
	}

	currentDate := time.Now().Format("01022006")
	objectKey := fmt.Sprintf("%s/%s_%s_%s", directory, keyId, currentDate, fileName)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   fileContent,
	})
	if err != nil {
		return nil, err
	}

	return &objectKey, nil
}
