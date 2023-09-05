// this interface will provide the requried methods
package interfaces

import(
	models "ecommerce_order/order_dal/models"
)



type IOrder interface{
	CreateOrder(input *models.Orders) (models.Orders, error) 
	RemoveOrder(Customer_ID int32) (string, error)
	GetAllOrder(CustomerId int32) ([]models.Orders, error) 

	
}