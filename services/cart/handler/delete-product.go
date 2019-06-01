package handler

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"microservices/services/cart/genproto"
)

func (s *CartHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	clientId, err := s.TokenService.GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	productsString, err := s.RedisRepository.Get(clientId)
	if err != nil {
		return nil, err
	}
	var cartProducts []pb.CartProduct

	if len(productsString) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Product not found in the cart")
	}

	err = json.Unmarshal(productsString, &cartProducts)
	if err != nil {
		return nil, err
	}

	var index = -1
	for i, product := range cartProducts {
		if product.Product.Id == req.Id {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, status.Error(codes.InvalidArgument, "Product not found in the cart")
	}
	updatedCart := append(cartProducts[:index], cartProducts[index+1:]...)
	redisProduct, err := json.Marshal(updatedCart)

	_, err = s.RedisRepository.Save(clientId, string(redisProduct))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProductResponse{
		Id: req.Id,
	}, nil
}
