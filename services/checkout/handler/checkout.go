package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"microservices/pkg"
	"microservices/services/checkout/genproto"
	"strconv"
)

func (s CheckoutHandler) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	clientId, err := s.TokenService.GetAuthToken(ctx)
	if err != nil {
		return nil, err
	}
	contextWithToken := metadata.AppendToOutgoingContext(ctx, "AccessToken", clientId)
	cartConnection, err := pkg.ConnectToGRPC("cart-service")
	if cartConnection != nil {
		defer cartConnection.Close()
	}
	if err != nil {
		return nil, err
	}
	cartServiceClient := pb.NewCartServiceClient(cartConnection)
	cart, err := cartServiceClient.GetCart(contextWithToken, &pb.GetCartRequest{})
	if err != nil {
		return nil, err
	}
	amount, err := getTotalAmountFromCartProducts(cart.Products)
	if err != nil {
		return nil, err
	}
	paymentConnection, err := pkg.ConnectToGRPC("payment-service")
	if paymentConnection != nil {
		defer paymentConnection.Close()
	}
	if err != nil {
		return nil, err
	}
	paymentServiceClient := pb.NewPaymentServiceClient(paymentConnection)
	paymentResponse, err := paymentServiceClient.ExecutePayment(contextWithToken, &pb.ExecutePaymentRequest{
		Amount:     amount,
		CardNumber: req.CardNumber,
	})
	if err != nil {
		return nil, err
	}
	_, err = cartServiceClient.Clear(contextWithToken, &pb.ClearCartRequest{})
	if err != nil {
		return nil, err
	}
	return &pb.CheckoutResponse{TransactionId: paymentResponse.TransactionId}, nil
}

func getTotalAmountFromCartProducts(cartProducts []*pb.CartProduct) (float64, error) {
	var amount = 0.00
	for _, product := range cartProducts {
		price := product.Product.Price
		floatString := fmt.Sprintf("%d.%d", price.Units, price.Decimals)
		num, err := strconv.ParseFloat(floatString, 32)
		if err != nil {
			return 0.00, err
		}

		amount += num * float64(product.Quantity)
	}
	return amount, nil
}
