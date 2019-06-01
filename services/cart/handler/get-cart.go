package handler

import (
	"context"
	"encoding/json"
	"microservices/services/cart/genproto"
)

func (s *CartHandler) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	clientId, err := s.TokenService.GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	productsString, err := s.RedisRepository.Get(clientId)
	if err != nil {
		return nil, err
	}
	var cartProducts = make([]*pb.CartProduct, 0)
	if len(productsString) > 0 {
		err = json.Unmarshal(productsString, &cartProducts)
		if err != nil {
			return nil, err
		}
	}
	return &pb.GetCartResponse{
		Products: cartProducts,
	}, nil
}
