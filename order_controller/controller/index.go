package controller

import (
	"context"
	"netxd_ecommerce/order_dal/interfaces"
	models "netxd_ecommerce/order_dal/models"
	pro "netxd_ecommerce/order_proto"
)

type RPCServer struct {
	pro.UnimplementedOrderServiceServer
}

var (
	OrderService interfaces.IOrder
)

func (s *RPCServer) RemoveOrderCustomer(ctx context.Context, req *pro.RemoveOrderRequest) (*pro.RemoveOrderResponse, error) {

	result, err := OrderService.RemoveOrder(req.Customer_ID)
	if err != nil {
		return nil, err
	}
	responseCustomer := &pro.RemoveOrderResponse{
		Status: result,
	}

	return responseCustomer, nil
}

func (s *RPCServer) CreateOrder(ctx context.Context, req *pro.CustomerOrder) (string, error) {
	dbInsert := &models.Orders{
		CustomerId:    req.CustomerId,
		PaymentId:     req.PaymentId,
		PaymentStatus: req.PaymentStatus,
		Status:        req.Status,
		Currency:      req.Currency,
		Items:         req.Items,
		Carrier:       req.Carrier,
		Tracking:      req.Tracking,
	}
	for _, protoShipping := range req.Shipping {
		shipping := models.Shipping{}
		for _, protoAddress := range protoShipping.Address {
			address := models.Address{
				Street1: protoAddress.Street1,
				Street2: protoAddress.Street2,
				City:    protoAddress.City,
				State:   protoAddress.State,
				Country: protoAddress.Country,
				Zip:     protoAddress.Zip,
			}
			shipping.Address = append(shipping.Address, address)
		}
		for _, protoOrigin := range protoShipping.Origin {
			origin := models.Origin{
				Street1: protoOrigin.Street1,
				Street2: protoOrigin.Street2,
				City:    protoOrigin.City,
				State:   protoOrigin.State,
				Country: protoOrigin.Country,
				Zip:     protoOrigin.Zip,
			}
			shipping.Origin = append(shipping.Origin, origin)
		}
		dbInsert.Shipping = append(dbInsert.Shipping, shipping)
	}
	var details interfaces.IOrder
	_, err := details.CreateOrder(dbInsert)
	if err != nil {
		return "failed", err

	}
	return "success", nil
}
func (s *RPCServer) GetOrderDetails(ctx context.Context, req *pro.GetOrderRequest) (*pro.GetOrderResponse, error) {
	result, err := OrderService.GetAllOrder(req.CustomerId)
	if err != nil {
		return nil, err
	}
	responseCustomer := &pro.RemoveOrderResponse{
		Status: result,
	}

	return responseCustomer, nil

}