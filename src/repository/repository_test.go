package repository

import (
	"fmt"
	"io/ioutil"
	"log"
	"mcx002/filestoreService/src"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

var NewMinioClientMock func(appConfig *src.AppConfig) (*minio.Client, error)
var NewBucketMock func(appConfig *src.AppConfig, minioClient *minio.Client) error
var NewRedisMock func(appConfig *src.AppConfig) *redis.Client

type MinioAdapterMock struct{}
type RedisAdapterMock struct{}

func (MinioAdapterMock) NewMinioClient(appConfig *src.AppConfig) (*minio.Client, error) {
	return NewMinioClientMock(appConfig)
}

func (MinioAdapterMock) NewBucket(appConfig *src.AppConfig, minioClient *minio.Client) error {
	return NewBucketMock(appConfig, minioClient)
}

func (RedisAdapterMock) NewRedis(appConfig *src.AppConfig) *redis.Client {
	return NewRedisMock(appConfig)
}

func TestNewRepository(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	// Case failed connect to minio error
	minioAdapter := MinioAdapterMock{}
	redisAdapter := RedisAdapterMock{}
	appConfig := &src.AppConfig{}
	NewMinioClientMock = func(appConfig *src.AppConfig) (*minio.Client, error) {
		return nil, fmt.Errorf("error")
	}

	_, err := NewRepository(appConfig, minioAdapter, redisAdapter)
	if err == nil {
		t.Errorf("got error nil; want err failed connect to minio")
	}

	// Case failed create or get bucket error
	NewMinioClientMock = func(appConfig *src.AppConfig) (*minio.Client, error) {
		return &minio.Client{}, nil
	}
	NewBucketMock = func(appConfig *src.AppConfig, minioClient *minio.Client) error {
		return fmt.Errorf("error")
	}

	_, err = NewRepository(appConfig, minioAdapter, redisAdapter)
	if err == nil {
		t.Errorf("got error nil; want error failed create bucket")
	}

	// Case create repo success
	NewBucketMock = func(appConfig *src.AppConfig, minioClient *minio.Client) error {
		return nil
	}
	NewRedisMock = func(appConfig *src.AppConfig) *redis.Client {
		return &redis.Client{}
	}

	repo, err := NewRepository(appConfig, minioAdapter, redisAdapter)
	if repo == nil {
		t.Errorf("got repo nil; want repo not nil")
	}
	if err != nil {
		t.Errorf("got err; want err nil")
	}
}
