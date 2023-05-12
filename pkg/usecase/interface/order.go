package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
)

type Orderusecase interface {
	PlaceOrder(ctx context.Context, UserID, paymentTypeId int) (domain.Orders, error)
	Razorpay(ctx context.Context, UserID, paymentMethodId int) (response.RazorPayResponse, error)
	VerifyRazorPay(ctx context.Context, body urequest.RazorPayRequest) error
	CancelOrder(ctx context.Context, orderId, userId int) error
	Listorders(ctx context.Context, userid int) ([]domain.Orders, error)
	Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error)
	ReturnOrder(userId, orderId int) (float64, error)
}
