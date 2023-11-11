package repository

import (
	"context"
	"mcx002/filestoreService/src"
	"mcx002/filestoreService/src/adapter"
	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type RepositoryInterface interface {
	GetMinioPresignedPutUrl(ctx context.Context, objName string) (*url.URL, error)
	GetMinioPresignedGetUrl(ctx context.Context, objName string) (*url.URL, error)
	GetCacheImageUrl(ctx context.Context, filename string) (string, error)
	SetCacheImageUrl(ctx context.Context, filename string, url string) error
}

type Repository struct {
	MinioClient *minio.Client
	AppConfig   *src.AppConfig
	Redis       *redis.Client
}

func NewRepository(
	appConfig *src.AppConfig,
	minioAdapter adapter.MinioInterface,
	redisAdapter adapter.RedisInterface) (*Repository, error) {
	minioClient, err := minioAdapter.NewMinioClient(appConfig)
	if err != nil {
		return nil, err
	}

	err = minioAdapter.NewBucket(appConfig, minioClient)
	if err != nil {
		return nil, err
	}

	return &Repository{
		MinioClient: minioClient,
		AppConfig:   appConfig,
		Redis:       redisAdapter.NewRedis(appConfig),
	}, nil
}
