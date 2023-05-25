package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
)

type OrderRepo interface {
	OrderAll(ctx context.Context, UserID, paymentTypeId int) (domain.Orders, error)
	CancelOrder(ctx context.Context, orderId, userId int) error
	Listorders(ctx context.Context) ([]response.OrderResponse, error)
	Listorder(ctx context.Context, Orderid int, UserId int) (order domain.Orders, err error)
	ReturnOrder(userId, orderId int) (float64, error)
	AdminListorders(ctx context.Context, pagination requests.Pagination) (orders []domain.Orders, err error)
	ListofOrderStatuses(ctx context.Context) (status []domain.OrderStatus, err error)
	UpdateOrderStatus(ctx context.Context, update requests.Update) error
	// FindWalletByUserID(ctx context.Context, userID uint) (wallet domain.Wallet, err error)
	// SaveWallet(ctx context.Context, userID uint) (walletID uint, err error)
	// UpdateWallet(ctx context.Context, walletID, upateTotalAmount uint) error
	// SaveWalletTransaction(ctx context.Context, walletTrx domain.Transaction) error
	// FindWalletTransactions(ctx context.Context, walletID uint, pagination requests.Pagination) (transaction []domain.Transaction, err error)
}
