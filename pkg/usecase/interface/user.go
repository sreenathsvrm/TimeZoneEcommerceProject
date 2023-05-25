package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
)

type UserUseCase interface {
	UserSignup(ctx context.Context, user requests.Usersign) (response.UserValue, error)
	UserLogin(ctx context.Context, user requests.Login) (string, error)
	OtpLogin(mobno string) (string, error)
	AddAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error)
	UpdateAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error)
	VeiwAdress(ctx context.Context, UserID int) (domain.Address, error)
	AddToWishList(ctx context.Context, wishList domain.WishList) error
	ListWishlist(ctx context.Context, userID uint) ([]response.Wishlist, error)
	RemoveFromWishList(ctx context.Context, wishList domain.WishList) error
}
