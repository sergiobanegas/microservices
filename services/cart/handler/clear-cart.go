package handler

import (
	"context"
	"microservices/services/cart/genproto"
)

func (s *CartHandler) Clear(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
	clientId, err := s.TokenService.GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	err = s.RedisRepository.Delete(clientId)
	if err != nil {
		return nil, err
	}
	return &pb.ClearCartResponse{
		Id: req.Id,
	}, nil
}
