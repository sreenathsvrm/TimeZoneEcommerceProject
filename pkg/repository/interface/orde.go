package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
)

type OrderRepo interface {
	OrderAll(ctx context.Context, UserID, paymentTypeId int) (domain.Orders, error)
	CancelOrder(ctx context.Context, orderId, userId int) error
	Listorders(ctx context.Context) ([]response.OrderResponse, error)
	Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error)
	ReturnOrder(userId, orderId int) (float64, error)
	AdminListorders(ctx context.Context, pagination urequest.Pagination) (orders []domain.Orders, err error)
	ListofOrderStatuses(ctx context.Context) (status []domain.OrderStatus, err error)
	UpdateOrderStatus(ctx context.Context, update urequest.Update) error
}
