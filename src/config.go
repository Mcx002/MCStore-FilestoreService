package src

import (
	"fmt"
	"mcx002/filestoreService/src/utils"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	ServerStartTime time.Time
	Host            string
	Port            int32

	// Service
	Name    string
	Version string

	// Minio
	MinioEndpoint        string
	MinioAccessKeyID     string
	MinioSecretAccessKey string
	MinioUseSSL          bool
	MinioBucketName      string
	MinioBucketLocation  string
	MinioExpiredTime     time.Duration

	// Redis
	RedisHost     string
	RedisPort     int32
	RedisPassword string
}

func NewAppConfig() (*AppConfig, error) {
	// Get Port Env
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("failed to get env port :%v", err)
	}

	minioUseSSL := utils.StringToBool(os.Getenv("MINIO_USE_SSL"))

	minioExpiredTime := time.Second * 24 * 60 * 60
	minioExpiredTimeEnv, err := strconv.Atoi(os.Getenv("MINIO_EXPIRED_TIME"))
	if err == nil {
		minioExpiredTime = time.Second * time.Duration(minioExpiredTimeEnv)
	}

	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return nil, fmt.Errorf("failed to load redis port :%v", err)
	}

	return &AppConfig{
		ServerStartTime: time.Now(),
		Host:            os.Getenv("HOST"),
		Port:            int32(port),
		Name:            os.Getenv("SERVICE_NAME"),
		Version:         os.Getenv("SERVICE_VERSION"),

		MinioEndpoint:        os.Getenv("MINIO_ENDPOINT"),
		MinioAccessKeyID:     os.Getenv("MINIO_ACCESS_KEY_ID"),
		MinioSecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		MinioUseSSL:          minioUseSSL,
		MinioBucketName:      os.Getenv("MINIO_BUCKET_NAME"),
		MinioBucketLocation:  os.Getenv("MINIO_BUCKET_LOCATION"),
		MinioExpiredTime:     minioExpiredTime,

		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     int32(redisPort),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}, nil
}
