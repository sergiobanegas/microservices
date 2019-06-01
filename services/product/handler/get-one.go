package handler

import (
	"context"
	"microservices/services/product/entity"
	"microservices/services/product/genproto"
)

func (s *ProductHandler) GetOne(ctx context.Context, req *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	var product entity.Product
	_, err := s.DB.FindOne(req.Id, &product)
	if err != nil {
		return nil, err
	}
	return &pb.GetOneResponse{
		Product: &pb.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price: &pb.Money{
				Units:    product.PriceUnits,
				Decimals: product.PriceDecimals,
			},
			Stock: product.Stock,
		},
	}, nil
}
