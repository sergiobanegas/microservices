package main

import (
	"context"
	"microservices/pkg"
	"microservices/services/cart/genproto"
	"microservices/services/cart/handler"
)

type cartServiceServer struct {
	redisRepository *pkg.RedisRepository
	tokenService    *pkg.TokenService
	handler         *handler.CartHandler
}

func NewCartServiceServer(redisRepository *pkg.RedisRepository, tokenService *pkg.TokenService) pb.CartServiceServer {
	return &cartServiceServer{redisRepository,
		tokenService,
		&handler.CartHandler{RedisRepository: redisRepository, TokenService: tokenService},
	}
}

func (s *cartServiceServer) HealthCheck(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: "UP"}, nil
}

func (s *cartServiceServer) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	return s.handler.AddProduct(ctx, req)
}

func (s *cartServiceServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	return s.handler.GetCart(ctx, req)
}

func (s *cartServiceServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	return s.handler.UpdateProduct(ctx, req)
}

func (s *cartServiceServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	return s.handler.DeleteProduct(ctx, req)
}

func (s *cartServiceServer) Clear(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
	return s.handler.Clear(ctx, req)
}
