package handler

import (
	"context"
	"microservices/services/product/entity"
	"microservices/services/product/genproto"
)

func (s *ProductHandler) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	var products []entity.Product
	_, err := s.DB.Find(&products)
	if err != nil {
		return nil, err
	}
	var productsResponse = make([]*pb.Product, len(products))
	for i := 0; i < len(products); i++ {
		productsResponse[i] = &pb.Product{
			Id:          products[i].Id,
			Name:        products[i].Name,
			Description: products[i].Description,
			Price: &pb.Money{
				Units:    products[i].PriceUnits,
				Decimals: products[i].PriceDecimals,
			},
			Stock: products[i].Stock,
		}
	}

	return &pb.SearchResponse{
		Products: productsResponse,
	}, nil
}
