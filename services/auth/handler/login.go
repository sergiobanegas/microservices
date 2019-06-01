package handler

import (
	"context"
	"microservices/services/auth/genproto"
)

func (s AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.TokenService.GenerateToken(req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}
