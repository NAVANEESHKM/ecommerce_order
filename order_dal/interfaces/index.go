// this interface will provide the requried methods
package interfaces

import (
	// "context"
	"ecommerce_order/order_dal/models"
	ecommerce_order "ecommerce_order/order_proto"
)

type IOrder interface {
	CreateOrder(input *models.Orders) (*ecommerce_order.CustomerResponse, error)
	RemoveOrder(Customer_ID int32) (string, error)
	GetAllOrder(CustomerId int32) ([]models.Orders, error)
	// checkItemAvailability(ctx context.Context, sku string, quantity string) (bool, error)
}
