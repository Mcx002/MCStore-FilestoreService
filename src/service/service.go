package service

import (
	"context"
	"mcx002/filestoreService/src"
	"mcx002/filestoreService/src/repository"
)

type Service struct {
	AppConfig  *src.AppConfig
	Context    *context.Context
	Repository repository.RepositoryInterface
}

func NewService(appConfig *src.AppConfig, repo repository.RepositoryInterface) *Service {
	return &Service{
		AppConfig:  appConfig,
		Context:    nil,
		Repository: repo,
	}
}
