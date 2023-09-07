package main

import (
	"context"
	// "fmt"

	"log"
	"net/http"
	// models "ecommerce_order/order_dal/models"
	pb "ecommerce_order/order_proto"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)
var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)
func main() {
    // r := gin.Default()
	r := gin.Default()
    conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewOrderServiceClient(conn)
    r.POST("/createorder", func(c *gin.Context) {
        var request pb.CustomerOrder
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        response, err := client.CreateOrder(c.Request.Context(), &request)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"value": response})
    })

	



    r.Run(":8080")
	// _, err1 := client.CreateOrder(context.Background(), &pb.CustomerOrder{
	// 	CustomerId:    "11",
	// 	PaymentId:     "your_payment_id",
	// 	PaymentStatus: "your_payment_status",
	// 	Status:        "your_status",
	// 	Currency:      "your_currency",
	// 	Items:[]*pb.Items{
	// 		{
	// 			Sku:         "SKU002",
	// 			Quantity:    5,
				
	// 			// Discount:    5.67,  // Your discount value
	// 			// PreTaxTotal: 18.01, // Your pre-tax total value
	// 			// Tax:         1.23,  // Your tax value
	// 			// Total:       19.24, // Your total value
	// 		},
	// 		// {
	// 		// 	Sku:         "SKU001",
	// 		// 	Quantity:    3,
	// 		// },
	// 		// Add more items if needed
	// 	},
	// 	Shipping: []*pb.Shipping{
	// 		{
	// 		   Address:[]*pb.Address{
	// 			{
    //                 Street1: "Anna Nagar",
	// 				Street2: "Gandhi nagar",
	// 				City: "Chennai",
	// 				State: "TamilNadu",
	// 				Country: "India",
	// 				Zip: "56456",
	// 			},

	// 		   },
	// 		   Origin: []*pb.Origin{
	// 			{
	// 				Street1: "Anna Nagar",
	// 				Street2: "Gandhi nagar",
	// 				City: "Chennai",
	// 				State: "TamilNadu",
	// 				Country: "India",
	// 				Zip: "56456",
	// 			},
	// 		   },
	// 		},

	// 	},
	// 	Carrier:  "your_carrier",
	// 	Tracking: "your_tracking",
	// })
	// r.POST("/create", func(c *gin.Context) {
	// 	var request pb.CustomerOrder
	// 	if err := c.ShouldBindJSON(&request); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	response, err := client.CreateOrder(c.Request.Context(), &request)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"value": response})
	// })

	// _, err1 := client.UpdateOrderDetails(context.Background(), &pb.UpdateOrderRequest{
	// 	       Customer_ID:    "11",
	// 			Sku:         "SKU002",
	// 			Quantity:    7,
		
	// 		// Add more items if needed
	// 	})


	// 	_, err1 := client.AddOrderDetails(context.Background(), &pb.UpdateOrderRequest{
	// 		Customer_ID:   "97",
	// 		 Sku:         "SKU002",
	// 		 Quantity:    7.0,
	 
	// 	 // Add more items if needed
	//  })



	// 	// Shipping: []*pb.Shipping{
			// {
			//    Address:[]*pb.Address{
			// 	{
            //         Street1: "Anna Nagar",
			// 		Street2: "Gandhi nagar",
			// 		City: "Chennai",
			// 		State: "TamilNadu",
			// 		Country: "India",
			// 		Zip: "56456",
			// 	},

			//    },
			//    Origin: []*pb.Origin{
			// 	{
			// 		Street1: "Anna Nagar",
			// 		Street2: "Gandhi nagar",
			// 		City: "Chennai",
			// 		State: "TamilNadu",
			// 		Country: "India",
			// 		Zip: "56456",
			// 	},
			//    },
			// },

		// },
			//)


	// _, err1 := client.RemoveOrderCustomer(context.Background(), &pb.RemoveOrderRequest{
	// 	CustomerId:345,
	// })

	// 	_, err1 := client.GetOrderDetails(context.Background(), &pb.GetOrderRequest{
	// 	CustomerId:76,
	// })
	// if err1 != nil {
	// 	log.Fatalf("Failed to call SayHello: %v", err1)
	// }
	// r.Run(":8080")
}
