package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/domain"
)

type CouponUseCase interface {
	CreateCoupon(ctx context.Context, coupon domain.Coupon) error
	UpdateCouponById(ctx context.Context, CouponId int, coupon requests.Coupon) (domain.Coupon, error)
	DeleteCoupon(ctx context.Context, CouponId int) (err error)
	ViewCoupon(ctx context.Context, couponID int) (domain.Coupon, error)
	ViewCoupons(ctx context.Context) ([]domain.Coupon, error)
	ApplyCoupontoCart(ctx context.Context, userID int, Code string) (float64, error)
}
