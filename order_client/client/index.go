package main

import (
	"context"
	// "fmt"

	"log"

	pb "netxd_ecommerce/order_proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	_, err1 := client.RemoveOrderCustomer(context.Background(), &pb.RemoveOrderRequest{
		Customer_ID:12,
	})
	if err1 != nil {
		log.Fatalf("Failed to call SayHello: %v", err1)
	}

	
}
