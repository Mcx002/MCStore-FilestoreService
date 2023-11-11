package server

import (
	"mcx002/filestoreService/src"
	pb "mcx002/filestoreService/src/proto_gen"
	"mcx002/filestoreService/src/service"
)

type Server struct {
	pb.UnimplementedFilestoreServer
	AppConfig *src.AppConfig
	Service   *service.Service
}

func NewServer(appConfig *src.AppConfig, svc *service.Service) *Server {
	return &Server{
		AppConfig: appConfig,
		Service:   svc,
	}
}
