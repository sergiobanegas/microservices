PROTO_DIR=../../api/proto
PROTO_OUT_DIR=./genproto

proto:
	@ mkdir -p $(PROTO_OUT_DIR)
	@ protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:$(PROTO_OUT_DIR) -I$(PROTO_DIR) $(PROTO_DIR)/*.proto
	@ protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:$(PROTO_OUT_DIR) -I$(PROTO_DIR) $(PROTO_DIR)/checkout-service.proto
	@ echo "Proto files generated"

install:
	@ go build -o run

run:
	@ ./run
