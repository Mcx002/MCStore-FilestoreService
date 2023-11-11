package service

import (
	"context"
	"log"

	"github.com/google/uuid"
)

func (s *Service) GetUploadUrl(ctx context.Context) (string, string, error) {
	id := uuid.New().String()
	url, err := s.Repository.GetMinioPresignedPutUrl(ctx, id)
	if err != nil {
		return "", "", err
	}
	return id, url.String(), nil
}

func (s *Service) GetImageUrl(ctx context.Context, filename string) (string, error) {
	urlCache, err := s.Repository.GetCacheImageUrl(ctx, filename)
	if err == nil {
		return urlCache, nil
	}

	url, err := s.Repository.GetMinioPresignedGetUrl(ctx, filename)
	if err != nil {
		return "", err
	}

	urlString := url.String()

	err = s.Repository.SetCacheImageUrl(ctx, filename, urlString)
	if err != nil {
		log.Printf("failed to set the cache: %v", err)
	}

	return urlString, nil
}
