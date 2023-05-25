package interfaces

import (
	"context"

	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
)

type UserRepository interface {
	UserSignup(ctx context.Context, user requests.Usersign) (response.UserValue, error)
	UserLogin(ctx context.Context, Email string) (domain.Users, error)
	OtpLogin(mbnum string) (int, error)
	AddAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error)
	UpdateAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error)
	VeiwAdress(ctx context.Context, UserID int) (domain.Address, error)
	RemoveWishListItem(ctx context.Context, wishList domain.WishList) error
	SaveWishListItem(ctx context.Context, wishList domain.WishList) error
	FindAllWishListItemsByUserID(ctx context.Context, userID uint) ([]response.Wishlist, error)
	FindWishListItem(ctx context.Context, productID, userID uint) (domain.WishList, error)
	FindProduct(ctx context.Context, id uint) (response.Product, error)
}
