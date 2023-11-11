package service

import (
	"mcx002/filestoreService/src/proto_gen"
	"time"
)

func (s *Service) GetHealth() *proto_gen.Health {
	health := &proto_gen.Health{
		Name:     s.AppConfig.Name,
		Version:  s.AppConfig.Version,
		Lifetime: time.Now().Unix() - s.AppConfig.ServerStartTime.Unix(),
	}

	return health
}
