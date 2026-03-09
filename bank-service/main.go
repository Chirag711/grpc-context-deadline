package main

import (
	"context"
	"fmt"
	pb "grpc-context-deadline/grpc-context-deadline/pkg/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBankServiceServer
}

func (s *server) ProcessPayment(ctx context.Context, req *pb.BankRequest) (*pb.BankResponse, error) {
	fmt.Println("bank processsing payment for order:", req.OrderId)

	return &pb.BankResponse{
		Status: "Payment successful",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("fail to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBankServiceServer(grpcServer, &server{})

	fmt.Println("Bank service is runing on port 50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve %v", err)
	}
}
