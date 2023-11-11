package repository

import (
	"context"
	"net/url"
)

func (r *Repository) GetMinioPresignedPutUrl(ctx context.Context, objName string) (*url.URL, error) {
	object, err := r.MinioClient.PresignedPutObject(
		ctx,
		r.AppConfig.MinioBucketName,
		objName,
		r.AppConfig.MinioExpiredTime)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (r *Repository) GetMinioPresignedGetUrl(ctx context.Context, objName string) (*url.URL, error) {
	object, err := r.MinioClient.PresignedGetObject(
		ctx,
		r.AppConfig.MinioBucketName,
		objName,
		r.AppConfig.MinioExpiredTime,
		url.Values{})
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (r *Repository) GetCacheImageUrl(ctx context.Context, filename string) (string, error) {
	val, err := r.Redis.Get(ctx, filename).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *Repository) SetCacheImageUrl(ctx context.Context, filename string, url string) error {
	err := r.Redis.Set(ctx, filename, url, r.AppConfig.MinioExpiredTime).Err()
	if err != nil {
		return err
	}

	return nil
}
