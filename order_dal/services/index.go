package services

import (
	"context"
	// "fmt"

	"log"

	"go.mongodb.org/mongo-driver/bson"

	"ecommerce_order/order_dal/interfaces"
	models "ecommerce_order/order_dal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	client          *mongo.Client
	OrderCollection *mongo.Collection

	ctx context.Context
}

func InitCustomerService(client *mongo.Client, collection *mongo.Collection, ctx context.Context) interfaces.IOrder {
	return &CustomerService{client, collection, ctx}
}

func (p *CustomerService) CreateOrder(input *models.Orders) (models.Orders, error) {
	
   insertResult, err1 := p.OrderCollection.InsertOne(p.ctx, &input)
   if err1 != nil {
   
   }


   insertedID := insertResult.InsertedID.(primitive.ObjectID)


   filter := bson.M{"_id": insertedID}
   var result models.Orders
   err := p.OrderCollection.FindOne(p.ctx, filter).Decode(&result)
   if err != nil {
     return result,err
   }
   return result,nil
}

func (p *CustomerService) RemoveOrder(Customer_ID int32) (string, error) {
    // fmt.Println(Customer_ID)
	filter := bson.M{"customerid": Customer_ID} // Replace with your filter criteria.

	// Delete the document that matches the filter.
	_, err := p.OrderCollection.DeleteOne(p.ctx,filter)

	if err != nil {
        // fmt.Println(Customer_ID)
		return "Unable to delete", err
	}
	return "Deleted Successfully", nil
}

func (p *CustomerService) GetAllOrder(CustomerId int32) ([]models.Orders, error) {
	filter := bson.M{"customerid": CustomerId}

	// Execute the find query.
	cursor, err := p.OrderCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate through the results and decode each document into a Person struct.
	var results []models.Orders
	for cursor.Next(context.Background()) {
		var result models.Orders
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return results, nil
}
