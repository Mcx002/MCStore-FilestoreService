package repository

import (
	"context"
	"mcx002/filestoreService/src"
	"mcx002/filestoreService/src/adapter"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestGetMinioPresignedPutUrl(t *testing.T) {
	godotenv.Load("../../.env.test")
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	repo, err := NewRepository(appConfig, adapter.NewMinioAdapter(), adapter.NewRedisAdapter())
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	_, err = repo.GetMinioPresignedPutUrl(nil, "objName")
	if err == nil {
		t.Error("got err nil; want err not nil")
	}

	url, err := repo.GetMinioPresignedPutUrl(context.Background(), "objName")
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}
	if url == nil {
		t.Error("got url nil; want url not nil")
	}
}

func TestGetMinioPresignedGetUrl(t *testing.T) {
	godotenv.Load("../../.env.test")
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	repo, err := NewRepository(appConfig, adapter.NewMinioAdapter(), adapter.NewRedisAdapter())
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	_, err = repo.GetMinioPresignedGetUrl(nil, "objName")
	if err == nil {
		t.Error("got err nil; want err not nil")
	}

	url, err := repo.GetMinioPresignedGetUrl(context.Background(), "objName")
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}
	if url == nil {
		t.Error("got url nil; want url not nil")
	}
}

func TestGetCacheImageUrl(t *testing.T) {
	godotenv.Load("../../.env.test")
	ctx := context.Background()
	filename := "objName"
	value := "valObjName"

	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	repo, err := NewRepository(appConfig, adapter.NewMinioAdapter(), adapter.NewRedisAdapter())
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	_, err = repo.GetCacheImageUrl(ctx, filename)
	if err == nil {
		t.Error("got err nil; want err not nil")
	}

	err = repo.Redis.Set(ctx, filename, value, time.Second*5).Err()
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	url, err := repo.GetCacheImageUrl(ctx, filename)
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	if url != value {
		t.Errorf("got url = %s; want url = %s", url, value)
	}

	err = repo.Redis.Del(ctx, filename).Err()
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}
}

func TestSetCacheImageUrl(t *testing.T) {
	godotenv.Load("../../.env.test")
	ctx := context.Background()
	filename := "objName"
	value := "valObjName"

	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	repo, err := NewRepository(appConfig, adapter.NewMinioAdapter(), adapter.NewRedisAdapter())
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	err = repo.SetCacheImageUrl(ctx, filename, value)
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}

	str, err := repo.Redis.Get(ctx, filename).Result()
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}
	if str != value {
		t.Errorf("got str = %s; want str = %s", str, value)
	}

	err = repo.Redis.Del(ctx, filename).Err()
	if err != nil {
		t.Errorf("got err %v; want err nil", err)
	}
}
