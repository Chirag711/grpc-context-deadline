package main

import (
	"context"
	"fmt"
	pb "grpc-context-deadline/grpc-context-deadline/pkg/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Connectin faild... %v", err)
	}

	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	res, err := client.PayOrder(context.Background(), &pb.PaymentRequest{
		OrderId: "ORD123456",
		Amount:  500,
		Method:  "GooglePay",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Payment Status : ", res.Status)

}
