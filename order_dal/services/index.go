package services

import (
	"context"
	"ecommerce_order/order_dal/interfaces"
	"ecommerce_order/order_dal/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)

type CustomerService struct {
	client              *mongo.Client
	OrderCollection     *mongo.Collection
	InventoryCollection *mongo.Collection
	ctx                 context.Context
}

func InitCustomerService(client *mongo.Client, collection1 *mongo.Collection, collection2 *mongo.Collection, ctx context.Context) interfaces.IOrder {
	return &CustomerService{client, collection1, collection2, ctx}
}

func (p *CustomerService) CreateOrder(input *models.Orders) (models.Orders, error) {
	var basePrice float32
	var discountvalue float32
	for i, val := range input.Items {
		filter := bson.M{"sku": val.Sku}
		inventoryResult := p.InventoryCollection.FindOne(p.ctx, filter)
		var inventoryDocument bson.M
		if err := inventoryResult.Decode(&inventoryDocument); err != nil {
			return models.Orders{}, err
		}
		num, err := strconv.ParseFloat(val.Quantity, 32) 
		if err != nil {
			fmt.Println("Conversion error:", err)

		}
		num32 := float32(num)
		price := inventoryDocument["price"].(bson.M)
		quantity := inventoryDocument["quantity"].(float64)
		if quantity < float64(num32) {
			fmt.Println("Lack of Quantity")
			return models.Orders{}, err
		}
		price64 := price["base"].(float64)
		discount := price["discount"].(float64)
		basePrice = float32(price64)
		discountvalue = float32(discount)
		fmt.Println(price64, discount, basePrice)
		input.Items[i].Price = num32 * basePrice
		input.Items[i].Discount =discountvalue
		input.Items[i].PreTaxTotal=input.Items[i].Price - ((discountvalue / 100) * 100)
		input.Items[i].Total = input.Items[i].PreTaxTotal
		fmt.Println(val.Price, val.Discount)
	}
	insertResult, err1 := p.OrderCollection.InsertOne(p.ctx, &input)
	if err1 != nil {

	}
	insertedID := insertResult.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": insertedID}
	var result models.Orders
	err2 := p.OrderCollection.FindOne(p.ctx, filter).Decode(&result)
	if err2 != nil {
		return result, err2
	}
	fmt.Println("Success")
	return result, nil
}
func (p *CustomerService) RemoveOrder(Customer_ID int32) (string, error) {
	filter := bson.M{"customerid": Customer_ID}
	_, err := p.OrderCollection.DeleteOne(p.ctx, filter)
	if err != nil {
		return "Unable to delete", err
	}
	return "Deleted Successfully", nil
}
func (p *CustomerService) GetAllOrder(CustomerId int32) ([]models.Orders, error) {
	filter := bson.M{"customerid": CustomerId}
	cursor, err := p.OrderCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
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

func (p *CustomerService) UpdateOrder(input *models.UpdateDetailsModel) (string, error) {
	var basePrice float32
	var discountvalue float32
	filter := bson.M{"sku": input.Sku}
	inventoryResult := p.InventoryCollection.FindOne(p.ctx, filter)
	var inventoryDocument bson.M
	if err := inventoryResult.Decode(&inventoryDocument); err != nil {

	}
	price := inventoryDocument["price"].(bson.M)
	price64 := price["base"].(float64)
	discount := price["discount"].(float64)
	discountvalue = float32(discount)
	basePrice = float32(price64)
	fmt.Println("The Price is ", basePrice)
	num, err := strconv.ParseFloat(input.Quantity, 32)
	if err != nil {
		fmt.Println("Conversion error:", err)
	}
	num32 := float32(num)
	input.Price = num32 * basePrice
	input.Discount = input.Price - ((discountvalue / 100) * 100)
	input.Total = input.Discount
	filter1 := bson.M{"customerid": input.Customer_ID}
	update := bson.M{
		"$set": bson.M{

			"items": []models.Items{
				{
					Sku:      input.Sku,
					Quantity: input.Quantity,
					Price:    input.Price,
					Discount: input.Discount,
					Total:    input.Total,
				},
			},
		},
	}
	_, err1 := p.OrderCollection.UpdateOne(context.Background(), filter1, update)
	if err1 != nil {
		return "updation failed", err1
	}
	return "Updation Success", nil

}

func (p *CustomerService) AddOrder(input *models.UpdateDetailsModel) (string, error) {
	var basePrice float32
	var discountvalue float32
	filter := bson.M{"sku": input.Sku}
	inventoryResult := p.InventoryCollection.FindOne(p.ctx, filter)
	var inventoryDocument bson.M
	if err := inventoryResult.Decode(&inventoryDocument); err != nil {
	}
	price := inventoryDocument["price"].(bson.M)
	price64 := price["base"].(float64)
	discount := price["discount"].(float64)
	discountvalue = float32(discount)
	basePrice = float32(price64)
	fmt.Println("The Price is ", basePrice)
	num, err := strconv.ParseFloat(input.Quantity, 32)
	if err != nil {
		fmt.Println("Conversion error:", err)
	}
	num32 := float32(num)
	fmt.Println("The quantity is ", num32)
	fmt.Println("The two values are ", num32, " ", basePrice)
	input.Price = num32 * basePrice
	input.Discount = input.Price - ((discountvalue / 100) * 100)
	input.Total = input.Discount
	itemsToAdd := []models.Items{
		{
			Sku:      input.Sku,
			Quantity: input.Quantity,
			Price:    input.Price,
			Discount: input.Discount,
			Total:    input.Total,
		},
	}
	filter1 := bson.M{"customerid": input.Customer_ID}
	update := bson.M{
		"$push": bson.M{
			"items": bson.M{
				"$each": itemsToAdd, 
			},
		},
	}
	_, err1 := p.OrderCollection.UpdateOne(context.Background(), filter1, update)
	if err1 != nil {
		return "updation failed", err1
	}
	return "Updation Success", nil

}
