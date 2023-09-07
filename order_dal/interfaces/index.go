
package interfaces

import (
	"ecommerce_order/order_dal/models"
)

type IOrder interface {
	CreateOrder(input *models.Orders) (models.Orders, error)
	RemoveOrder(Customer_ID int32) (string, error)
	GetAllOrder(CustomerId int32) ([]models.Orders, error)
	UpdateOrder(input *models.UpdateDetailsModel)(string,error)
	AddOrder(input *models.UpdateDetailsModel) (string, error)
}
