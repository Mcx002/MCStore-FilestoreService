package adapter

import (
	"mcx002/filestoreService/src"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewRedis(t *testing.T) {
	godotenv.Load("../../.env.test")
	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Error("app config should be successfully loaded")
	}

	redisAdapter := NewRedisAdapter()
	redis := redisAdapter.NewRedis(appConfig)
	if redis == nil {
		t.Errorf("got redis nil; want redis not nil")
	}
}
