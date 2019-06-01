package main

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"microservices/pkg"
	"microservices/services/auth/genproto"
	"microservices/services/auth/handler"
)

type authServiceServer struct {
	handler *handler.AuthHandler
}

func NewAuthServiceServer(tokenService *pkg.TokenService) pb.AuthServiceServer {
	return &authServiceServer{&handler.AuthHandler{TokenService: tokenService}}
}

func (authServiceServer) HealthCheck(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: "UP"}, nil
}

func (s authServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.handler.Login(ctx, req)
}
