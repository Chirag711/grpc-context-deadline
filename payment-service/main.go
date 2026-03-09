package main

import (
	"context"
	"fmt"
	pb "grpc-context-deadline/grpc-context-deadline/pkg/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) PayOrder(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	fmt.Println("payment received via:", req.Method, " Rs", req.Amount)

	conn, err := grpc.Dial(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pb.NewBankServiceClient(conn)

	res, err := client.ProcessPayment(context.Background(), &pb.BankRequest{
		OrderId: req.OrderId,
		Amount:  req.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &pb.PaymentResponse{
		Status: res.Status,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("fail to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(grpcServer, &server{})

	fmt.Println("Payment service is runing on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve %v", err)
	}
}
