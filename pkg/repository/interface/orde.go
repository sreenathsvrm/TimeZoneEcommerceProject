package interfaces

import (
	"context"
	"ecommerce/pkg/domain"
)


type OrderRepo interface {
	OrderAll(ctx context.Context,UserID, paymentTypeId int) (domain.Orders, error)
	CancelOrder(ctx context.Context,orderId, userId int) error
	Listorders(ctx context.Context,userid int)([]domain.Orders,error)
	Listorder(ctx context.Context,Orderid int,UserId int)(order domain.Orders, err error ) 
	ReturnOrder(userId, orderId int) (float64,error)
}
