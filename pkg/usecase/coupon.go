package usecase

import (
	"context"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
	"errors"
	"math/rand"
	"time"
)

type couponUsecase struct {
	CouponRepo interfaces.CouponRepo
}

func NewCouponUseCase(repo interfaces.CouponRepo) services.CouponUseCase {
	return &couponUsecase{
		CouponRepo: repo,
	}
}

const (
	couponCodeLength = 8
)

func generateCouponCode() string {
	// Define the characters allowed in the coupon code
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Initialize the random seed
	rand.Seed(time.Now().UnixNano())

	// Generate a random coupon code
	code := make([]byte, couponCodeLength)
	for i := 0; i < couponCodeLength; i++ {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

func (c *couponUsecase) CreateCoupon(ctx context.Context, coupon domain.Coupon) error {
	// Validate coupon data
	if coupon.DiscountPercent <= 0 {
		return errors.New("invalid discount amount")
	}
	if coupon.ExpiryDate.Before(time.Now()) {
		return errors.New("coupon has already expired")
	}

	if coupon.UsageLimits < 0 {
		return errors.New("invalid usage limits")
	}

	// Generate a unique coupon code if needed
	if coupon.Code == "" {
		coupon.Code = generateCouponCode()
	}

	err := c.CouponRepo.AddCoupon(ctx, coupon)
	if err != nil {
		return err
	}

	return nil
}

func (c *couponUsecase) UpdateCouponById(ctx context.Context, CouponId int, coupon urequest.Coupon) (domain.Coupon, error) {

	updated,err:=c.CouponRepo.UpdateCouponById(ctx,CouponId,coupon)
    return updated ,err
}
