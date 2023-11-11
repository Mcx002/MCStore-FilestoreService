package src

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestNewAppConfig(t *testing.T) {
	// load .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatalf("cannot load env :%v", err)
	}

	// load env
	appConfig, err := NewAppConfig()

	// err should nil
	if err != nil {
		t.Errorf("err = %v; want error nil", err)
	}

	// appConfig should running well
	if appConfig.Name != "Test Filestore Service" {
		t.Errorf("appConfig.Name = %s; want Filestore Service", appConfig.Name)
	}

	// Case Minio Expired Time got default value
	// manipulate env
	os.Setenv("MINIO_EXPIRED_TIME", "")

	// err should nil
	if err != nil {
		t.Errorf("err = %v; want error nil", err)
	}

	// load env
	appConfig, _ = NewAppConfig()

	// validate env
	if appConfig.MinioExpiredTime != (time.Second * 24 * 60 * 60) {
		t.Errorf("appConfig.Name = %s; want Filestore Service", appConfig.Name)
	}

	// Case load env redis port err
	// set REDIS_PORT to empty
	os.Setenv("REDIS_PORT", "")

	// load env
	_, err = NewAppConfig()

	// err should not nil
	if err == nil {
		t.Errorf("err = %v; want err failed to load redis port", err)
	}

	// Case load env port err
	// set PORT to empty
	os.Setenv("PORT", "")

	// load env
	_, err = NewAppConfig()

	// err should not nil
	if err == nil {
		t.Errorf("err = %v; want err failed to get env port", err)
	}

}
