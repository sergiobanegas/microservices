package pkg

import (
	"google.golang.org/grpc"
)

func ConnectToGRPC(serviceName string) (*grpc.ClientConn, error) {
	serviceAddress, err := LookupServiceWithConsul(serviceName)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	return conn, err
}
