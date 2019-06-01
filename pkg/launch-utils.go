package pkg

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func RunGrpcServer(ctx context.Context, server *grpc.Server, consulConfig *ConsulConfig) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			DeregisterService(consulConfig.ServiceName)
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	RegisterServiceWithConsul(consulConfig)
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(consulConfig.GRPCPort))
	if err != nil {
		panic("Error starting the server: " + err.Error())
	}
	print("Starting server...")
	return server.Serve(lis)
}

func RunHttpServer(ctx context.Context, httpPort int, mux *runtime.ServeMux) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(httpPort),
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
