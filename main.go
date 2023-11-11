package main

import (
	"fmt"
	"log"
	"mcx002/filestoreService/src"
	"mcx002/filestoreService/src/adapter"
	"mcx002/filestoreService/src/proto_gen"
	"mcx002/filestoreService/src/repository"
	"mcx002/filestoreService/src/server"
	"mcx002/filestoreService/src/service"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("cannot load env :%v", err)
	}

	// Load Env Data
	appConfig, err := src.NewAppConfig()
	if err != nil {
		log.Fatalf("failed to load appConfig :%v", err)
	}

	// Init adapter
	minioAdapter := adapter.NewMinioAdapter()
	redisAdapter := adapter.NewRedisAdapter()

	// Init repo
	repo, err := repository.NewRepository(appConfig, minioAdapter, redisAdapter)
	if err != nil {
		log.Fatalf("failed to load the repo :%v", err)
	}

	// Init service
	svc := service.NewService(appConfig, repo)
	if err != nil {
		log.Fatalf("failed to load the service :%v", err)
	}

	// Init server
	svr := server.NewServer(appConfig, svc)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	// Listen to TCP protocol connection
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port))
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	// Register the server
	s := grpc.NewServer()
	proto_gen.RegisterFilestoreServer(s, svr)
	log.Printf("server listening at %v", lis.Addr())

	// Run Server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
