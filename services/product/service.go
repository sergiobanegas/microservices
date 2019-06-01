package main

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"microservices/pkg"
	"microservices/services/product/genproto"
	"microservices/services/product/handler"
)

type productServiceServer struct {
	handler *handler.ProductHandler
}

func NewProductServiceServer(db *pkg.MysqlRepository) pb.ProductServiceServer {
	return &productServiceServer{&handler.ProductHandler{DB: db}}
}

func (s *productServiceServer) HealthCheck(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: "UP"}, nil
}

func (s *productServiceServer) GetOne(ctx context.Context, req *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	return s.handler.GetOne(ctx, req)
}

func (s *productServiceServer) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	return s.handler.Search(ctx, req)
}
