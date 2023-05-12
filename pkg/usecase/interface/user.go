package interfaces

import (
	"context"

	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
)

type UserUseCase interface {
	UserSignup(ctx context.Context, user urequest.Fusersign) (response.UserValue, error)
	UserLogin(ctx context.Context, user urequest.Flogin) (string, error)
	OtpLogin(mobno string) (string, error)
	AddAdress(ctx context.Context, UserID int, address urequest.AddressReq) (domain.Address, error)
	UpdateAdress(ctx context.Context, UserID int, address urequest.AddressReq) (domain.Address, error)
	VeiwAdress(ctx context.Context, UserID int) (domain.Address, error)
}
