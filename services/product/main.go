package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"microservices/pkg"
	"microservices/services/product/genproto"
	"strconv"
)

func main() {
	ctx := context.Background()

	server := grpc.NewServer()
	db, err := gorm.Open("mysql", "root:admin@/microservices?charset=utf8,utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Error connecting to db")
	}
	defer db.Close()
	repository := pkg.NewMysqlRepository(db)
	migrate(repository)
	productServiceServer := NewProductServiceServer(repository)
	pb.RegisterProductServiceServer(server, productServiceServer)

	consulConfig := &pkg.ConsulConfig{
		ServiceName: "product-service",
		Address:     pkg.GetOutboundIP(),
		GRPCPort:    1112,
		HTTPPort:    8080,
	}

	go func() {
		_ = runHttpServer(ctx, consulConfig)
	}()

	err = pkg.RunGrpcServer(ctx, server, consulConfig)
	if err != nil {
		panic("Error starting server")
	}
}

func runHttpServer(ctx context.Context, consulConfig *pkg.ConsulConfig) error {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, consulConfig.Address+":"+strconv.Itoa(consulConfig.GRPCPort), opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	return pkg.RunHttpServer(ctx, consulConfig.HTTPPort, mux)
}
