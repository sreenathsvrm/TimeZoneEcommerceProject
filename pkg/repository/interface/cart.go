package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
)

type CartRepo interface {
	SaveCart(ctx context.Context, Userid int) (uint, error)
	AddCartItem(ctx context.Context, Cartitem domain.CartItem) error
	FindCartIDNproductId(ctx context.Context, cart_id uint, product_id uint) (cartItem domain.CartItem, err error)
	FindCartByUserID(ctx context.Context, UserID int) (domain.Cart, error)
	FindProduct(ctx context.Context, id uint) (response.Product, error)
	RemoveCartItem(ctx context.Context, CartItemid uint) error
	AddQuantity(ctx context.Context, cartItemid uint, qty uint) error
	FindCartlistByCartID(ctx context.Context, cartID uint) (cartitems []response.Cartres, err error) 
}
