package adapter

import (
	"context"
	"log"
	"mcx002/filestoreService/src"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioInterface interface {
	NewMinioClient(appConfig *src.AppConfig) (*minio.Client, error)
	NewBucket(appConfig *src.AppConfig, minioClient *minio.Client) error
}

type MinioAdapter struct{}

func NewMinioAdapter() *MinioAdapter {
	return &MinioAdapter{}
}

func (MinioAdapter) NewMinioClient(appConfig *src.AppConfig) (*minio.Client, error) {
	endpoint := appConfig.MinioEndpoint
	accessKeyID := appConfig.MinioAccessKeyID
	secretAccessKey := appConfig.MinioSecretAccessKey
	useSSL := appConfig.MinioUseSSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func (MinioAdapter) NewBucket(appConfig *src.AppConfig, minioClient *minio.Client) error {
	ctx := context.Background()

	// Make a new bucket called testbucket.
	bucketName := appConfig.MinioBucketName
	location := appConfig.MinioBucketLocation

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if !(errBucketExists == nil && exists) {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	return nil
}
