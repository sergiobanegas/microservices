package handler

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"microservices/services/cart/genproto"
)

func (s *CartHandler) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {

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
	cartString, err := s.RedisRepository.Get(clientId)
	if err != nil {
		return nil, err
	}
	updatedCartString, err := updateCart(cartString, productResponse, req)

	_, err = s.RedisRepository.Save(clientId, updatedCartString)

	if err != nil {
		return nil, err
	}

	return &pb.AddProductResponse{
		Id: req.Id,
	}, nil
}

func updateCart(cartString []byte, product *pb.Product, req *pb.AddProductRequest) (string, error) {
	var cartProducts []pb.CartProduct
	if len(cartString) > 0 {
		err := json.Unmarshal(cartString, &cartProducts)
		if err != nil {
			return "", err
		}
		var found = false
		for i := 0; i < len(cartProducts); i++ {
			if cartProducts[i].Product.Id == req.Id {
				found = true
				if product.Stock < cartProducts[i].Quantity+req.Quantity {
					return "", status.Error(codes.InvalidArgument, "Quantity is greater than product stock")
				}
				cartProducts[i] = pb.CartProduct{
					Quantity: cartProducts[i].Quantity + req.Quantity,
					Product:  cartProducts[i].Product,
				}
				break
			}
		}
		if !found {
			cartProducts = append(cartProducts, pb.CartProduct{Product: product, Quantity: req.Quantity})
		}
	} else {
		cartProducts = append(cartProducts, pb.CartProduct{Product: product, Quantity: req.Quantity})
	}
	redisProduct, err := json.Marshal(cartProducts)
	if err != nil {
		return "", err
	}
	return string(redisProduct), nil
}
