package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	// "ecommerce/pkg/domain"
)

type CartUsecase interface {
	AddCartItem(ctx context.Context, body urequest.Cartreq) error
	RemoveFromCart(ctx context.Context, body urequest.Cartreq) error
	FindUserCart(ctx context.Context, userID int) (cart domain.Cart, err error)
	AddQuantity(ctx context.Context, body urequest.Addcount) error
	FindCartlistByCartID(ctx context.Context, cartID uint) (cartitems []response.Cartres, err error)
}
