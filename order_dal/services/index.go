package services

import (
	"context"

	"log"
	"go.mongodb.org/mongo-driver/bson"
	
	"netxd_ecommerce/order_dal/interfaces"
	models "netxd_ecommerce/order_dal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	client *mongo.Client
	OrderCollection    *mongo.Collection
	
	ctx                   context.Context
}



func InitCustomerService(client *mongo.Client, collection *mongo.Collection, ctx context.Context) interfaces.IOrder {
	return &CustomerService{client, collection, ctx}
}


func (p *CustomerService) CreateOrder(input *models.Orders)(string,error){
    _,err :=p.OrderCollection.InsertOne(p.ctx,&input)    
    if err!=nil{
        return "Failed",err
    }
    
    return "success",nil
}


func (p *CustomerService) RemoveOrder(Customer_ID int32) (string, error) {

	filter := bson.M{"Customer_ID": Customer_ID} // Replace with your filter criteria.

    // Delete the document that matches the filter.
    _, err := p.OrderCollection.DeleteOne(context.Background(), filter)
	
    if err != nil {
        return "Unable to delete",err
    }
	return "Deleted Successfully", nil
}

func (p*CustomerService) GetAllOrder(CustomerId string)(*models.Orders,error){
	filter := bson.M{"customerid":CustomerId }

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
	return &results,nil
}