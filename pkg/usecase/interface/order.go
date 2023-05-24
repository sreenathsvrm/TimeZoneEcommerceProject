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
	Listorders(ctx context.Context, userid int) ([]response.OrderResponse, error)
	Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error)
	ReturnOrder(userId, orderId int) (float64, error)
	ListofOrderStatuses(ctx context.Context) (status []domain.OrderStatus, err error)
	AdminListorders(ctx context.Context, pagination urequest.Pagination) (orders []domain.Orders,err error)
	UpdateOrderStatus(ctx context.Context, update urequest.Update) ( error) 

	// GetUserWallet(ctx context.Context, userID uint) (wallet domain.Wallet, err error)
	// GetUserWalletTransactions(ctx context.Context,userID uint, pagination urequest.Pagination) (transactions []domain.Transaction, err error)
}
