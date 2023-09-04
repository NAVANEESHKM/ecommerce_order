package main

import (
	"context"
	"fmt"
	"net"

	pro "netxd_ecommerce/order_proto" // Import the generated Go code
	"netxd_ecommerce/order_config/config"
	"netxd_ecommerce/order_config/constants"
	"netxd_ecommerce/order_dal/services"
     "netxd_ecommerce/order_controller/controller"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	OrderCollection := config.GetCollection(client, "DataBase", "Order")
	controller.OrderService = services.InitCustomerService(client,OrderCollection ,context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)	
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() //get the pointer reference of the server
	pro.RegisterOrderServiceServer(s, &controller.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
