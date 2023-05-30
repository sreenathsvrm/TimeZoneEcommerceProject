package usecase

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
	"math/rand"
	"time"

	"github.com/pkg/errors"
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

	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
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

func (c *couponUsecase) UpdateCouponById(ctx context.Context, CouponId int, coupon requests.Coupon) (domain.Coupon, error) {

	updated, err := c.CouponRepo.UpdateCouponById(ctx, CouponId, coupon)
	return updated, err
}

func (c *couponUsecase) DeleteCoupon(ctx context.Context, CouponId int) (err error) {
	err = c.CouponRepo.DeleteCoupon(ctx, CouponId)
	return err
}

func (c *couponUsecase) ViewCoupon(ctx context.Context, couponID int) (domain.Coupon, error) {
	coupon, err := c.CouponRepo.ViewCoupon(ctx, couponID)
	return coupon, err
}

func (c *couponUsecase) ViewCoupons(ctx context.Context) ([]domain.Coupon, error) {
	allcoupons, err := c.CouponRepo.ViewCoupons(ctx)
	return allcoupons, err
}

func (c *couponUsecase) ApplyCoupontoCart(ctx context.Context, userID int, Code string) (float64, error) {
	Total_price, err := c.CouponRepo.ApplyCoupontoCart(ctx, userID, Code)
	return Total_price, err
}
