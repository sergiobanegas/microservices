package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"microservices/pkg"
	"microservices/services/cart/genproto"
)

func validateQuantity(product *pb.Product, quantity int64) error {
	if product.Stock == 0 {
		return status.Error(codes.InvalidArgument, "Product out of stock")
	} else if product.Stock < quantity {
		return status.Error(codes.InvalidArgument, "Quantity is greater than product stock")
	}
	return nil
}

func getProduct(ctx context.Context, id string) (*pb.Product, error) {
	conn, err := pkg.ConnectToGRPC("product-service")
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return nil, err
	}

	c := pb.NewProductServiceClient(conn)

	getProductRequest := &pb.GetOneRequest{
		Id: id,
	}
	productResponse, err := c.GetOne(ctx, getProductRequest)
	if err != nil {
		return nil, status.Error(codes.NotFound, "The product doesn't exists")
	}
	return productResponse.Product, nil
}
