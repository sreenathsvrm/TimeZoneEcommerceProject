package interfaces

import (
	"context"

	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
)

type UserRepository interface {
	UserSignup (ctx context.Context, user urequest.Fusersign)(response.UserValue, error)
    UserLogin (ctx context.Context, Email string)(domain.Users,error)
	OtpLogin(mbnum string)(int ,error)
	AddAdress(ctx context.Context,UserID int,address urequest.AddressReq)(domain.Address,error)
	UpdateAdress(ctx context.Context, UserID int, address urequest.AddressReq) (domain.Address, error)
	VeiwAdress(ctx context.Context, UserID int) (domain.Address, error)
}
