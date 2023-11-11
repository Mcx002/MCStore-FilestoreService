package service

import (
	"mcx002/filestoreService/src"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetHealth(t *testing.T) {
	// load env
	godotenv.Load("../../.env.test")
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	// Create Service
	svc := NewService(appConfig, &RepositoryMock{})

	health := svc.GetHealth()

	if health.Name != appConfig.Name {
		t.Errorf("got health.Name = %s; want %s", health.Name, appConfig.Name)
	}
}
