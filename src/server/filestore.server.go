package server

import (
	"context"
	"mcx002/filestoreService/src/proto_gen"
)

func (s *Server) GetUploadUrl(ctx context.Context, in *proto_gen.FileDto) (*proto_gen.FileDto, error) {
	s.Service.Context = &ctx

	id, url, err := s.Service.GetUploadUrl(ctx)

	if err != nil {
		return nil, err
	}

	fileDto := &proto_gen.FileDto{
		Id:  id,
		Url: url,
	}

	return fileDto, nil
}

func (s *Server) GetImageUrl(ctx context.Context, in *proto_gen.FileDto) (*proto_gen.FileDto, error) {
	url, err := s.Service.GetImageUrl(ctx, in.Id)

	if err != nil {
		return nil, err
	}

	fileDto := &proto_gen.FileDto{
		Id:  in.Id,
		Url: url,
	}

	return fileDto, nil
}
