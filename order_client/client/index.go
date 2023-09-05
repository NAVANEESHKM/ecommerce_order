package main

import (
	"context"
	// "fmt"

	"log"

	pb "ecommerce_order/order_proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)



	// _,err1:=client.CreateOrder(context.Background(),&pb.CustomerOrder{
	// 	//   Id:"1",
	// 	  CustomerId: 76,
	// 	  PaymentId: "8678",
	// 	  PaymentStatus: "success",
	// })


	// _, err1 := client.RemoveOrderCustomer(context.Background(), &pb.RemoveOrderRequest{
	// 	CustomerId:345,
	// })


		_, err1 := client.GetOrderDetails(context.Background(), &pb.GetOrderRequest{
		CustomerId:345,
	})
	if err1 != nil {
		log.Fatalf("Failed to call SayHello: %v", err1)
	}

	
}