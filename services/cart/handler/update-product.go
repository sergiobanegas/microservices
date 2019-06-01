package handler

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"microservices/services/cart/genproto"
)

func (s *CartHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	productResponse, err := getProduct(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	err = validateQuantity(productResponse, req.Quantity)
	if err != nil {
		return nil, err
	}

	clientId, err := s.TokenService.GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	productsString, err := s.RedisRepository.Get(clientId)
	var cartProducts []pb.CartProduct

	if len(productsString) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Product not found in the cart")
	}
	err = json.Unmarshal(productsString, &cartProducts)
	if err != nil {
		return nil, err
	}

	found := false
	for i := 0; i < len(cartProducts); i++ {
		if cartProducts[i].Product.Id == req.Id {
			found = true
			cartProducts[i] = pb.CartProduct{
				Quantity: req.Quantity,
				Product:  cartProducts[i].Product,
			}
			break
		}
	}
	if !found {
		return nil, errors.New("product not found")
	}

	redisProduct, err := json.Marshal(cartProducts)

	_, err = s.RedisRepository.Save(clientId, string(redisProduct))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProductResponse{
		Id: req.Id,
	}, nil
}
