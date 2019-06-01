package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"microservices/pkg"
	"microservices/services/checkout/genproto"
	"strconv"
)

func main() {
	ctx := context.Background()

	server := grpc.NewServer()
	checkoutServiceServer := NewCheckoutServiceServer(&pkg.TokenService{})
	pb.RegisterCheckoutServiceServer(server, checkoutServiceServer)

	consulConfig := &pkg.ConsulConfig{
		ServiceName: "checkout-service",
		Address:     pkg.GetOutboundIP(),
		GRPCPort:    1115,
		HTTPPort:    8086,
	}

	go func() {
		_ = runHttpServer(ctx, consulConfig)
	}()

	err := pkg.RunGrpcServer(ctx, server, consulConfig)
	if err != nil {
		panic("Error starting server")
	}
}

func runHttpServer(ctx context.Context, consulConfig *pkg.ConsulConfig) error {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterCheckoutServiceHandlerFromEndpoint(ctx, mux, consulConfig.Address+":"+strconv.Itoa(consulConfig.GRPCPort), opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	return pkg.RunHttpServer(ctx, consulConfig.HTTPPort, mux)
}
