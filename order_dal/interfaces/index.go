// this interface will provide the requried methods
package interfaces

import(
	models "netxd_ecommerce/order_dal/models"
)



type IOrder interface{
	CreateOrder(input *models.Orders)(string,error)
	RemoveOrder(Customer_ID int32) (string, error)
	GetAllOrder(CustomerId string)([]models.Orders,error)

	
}