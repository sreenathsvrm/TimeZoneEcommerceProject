package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
)

type CouponRepo interface {
	AddCoupon(ctx context.Context, Coupon domain.Coupon) (err error)
	UpdateCouponById(ctx context.Context, CouponId int, coupon urequest.Coupon) (updatedCoupon domain.Coupon, err error)
}
