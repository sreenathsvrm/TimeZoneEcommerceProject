package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
)

type CartUsecase interface {
	AddCartItem(ctx context.Context, body requests.Cartreq) error
	RemoveFromCart(ctx context.Context, body requests.Cartreq) error
	FindUserCart(ctx context.Context, userID int) (cart domain.Cart, err error)
	AddQuantity(ctx context.Context, body requests.Addcount) error
	FindCartlistByCartID(ctx context.Context, cartID uint) (cartitems []response.Cartres, err error)
}
