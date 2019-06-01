package com.microservices.payment;

import io.grpc.stub.StreamObserver;
import org.lognet.springboot.grpc.GRpcService;
import pb.ExecutePaymentResponse;
import pb.PaymentServiceGrpc;

import java.util.UUID;

@GRpcService
public class PaymentService extends PaymentServiceGrpc.PaymentServiceImplBase {

    @Override
    public void executePayment(final pb.ExecutePaymentRequest request,
                               StreamObserver<ExecutePaymentResponse> responseObserver) {
        ExecutePaymentResponse response = ExecutePaymentResponse.newBuilder().setTransactionId(UUID.randomUUID().toString()).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

}
