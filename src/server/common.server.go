package server

import (
	"context"
	"mcx002/filestoreService/src/proto_gen"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) GetHealth(ctx context.Context, in *empty.Empty) (*proto_gen.Health, error) {
	health := s.Service.GetHealth()

	return health, nil
}
