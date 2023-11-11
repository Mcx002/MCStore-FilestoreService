package adapter

import (
	"context"
	"mcx002/filestoreService/src"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewMinioClient(t *testing.T) {
	godotenv.Load("../../.env.test")
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	minioAdapter := NewMinioAdapter()
	_, err = minioAdapter.NewMinioClient(appConfig)
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	appConfig.MinioEndpoint = ""
	_, err = minioAdapter.NewMinioClient(appConfig)
	if err == nil {
		t.Error("got err nil; want err not nil")
	}
}

func TestNewBucket(t *testing.T) {
	godotenv.Load("../../.env.test")
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	// Prepare minio client
	minioAdapter := NewMinioAdapter()
	minioClient, err := minioAdapter.NewMinioClient(appConfig)
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	// should return err nil
	err = minioAdapter.NewBucket(appConfig, minioClient)
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	// should create new bucket
	appConfig.MinioBucketName = "test-bucket"
	err = minioAdapter.NewBucket(appConfig, minioClient)
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	// should remove created bucket
	minioClient.RemoveBucket(context.Background(), "test-bucket")
}
