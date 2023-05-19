package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
)

type CouponUseCase interface {
	CreateCoupon(ctx context.Context, coupon domain.Coupon) error
	UpdateCouponById(ctx context.Context, CouponId int, coupon urequest.Coupon) (domain.Coupon, error)
}
