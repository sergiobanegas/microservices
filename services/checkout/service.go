package main

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"microservices/pkg"
	"microservices/services/checkout/genproto"
	"microservices/services/checkout/handler"
)

type checkoutServiceServer struct {
	handler *handler.CheckoutHandler
}

func NewCheckoutServiceServer(tokenService *pkg.TokenService) pb.CheckoutServiceServer {
	return &checkoutServiceServer{&handler.CheckoutHandler{TokenService: tokenService}}
}

func (checkoutServiceServer) HealthCheck(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: "UP"}, nil
}

func (s checkoutServiceServer) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	return s.handler.Checkout(ctx, req)
}
