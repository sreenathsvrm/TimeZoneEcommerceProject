package repository

import (
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	"errors"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CouponDB struct {
	DB *gorm.DB
}


func NewCouponrepo(DB *gorm.DB)interfaces.CouponRepo {
	return &CouponDB{
		DB: DB,
	}
}


func (c *CouponDB)AddCoupon(ctx context.Context, coupon domain.Coupon) error  {
	createCoupen := `INSERT INTO coupons (code, discount_percent,usage_limits,maximum_discount_price,minimum_purchase_price,expiry_date)
		VALUES($1,$2,$3,$4,$5,$6)`
	 if c.DB.Exec(createCoupen,
		coupon.Code,
		coupon.DiscountPercent,
		coupon.UsageLimits,
		coupon.MaximumDiscountPrice,
		coupon.MinimumPurchasePrice,
		coupon.ExpiryDate).Error!=nil { 
			return errors.New("error")
		}
	return nil
}

func (c *CouponDB)UpdateCouponById(ctx context.Context,CouponId int,coupon urequest.Coupon)(updatedCoupon domain.Coupon,err error)  {
	updateCoupon := `UPDATE coupons SET discount_percent=$1, usage_limits=$2, maximum_discount_price=$3, minimum_purchase_price=$4, expiry_date=$5
		 WHERE id=$6
		 RETURNING id, discount_percent, usage_limits, maximum_discount_price, minimum_purchase_price, expiry_date`

	err = c.DB.Raw(updateCoupon,
		coupon.DiscountPercent,
		coupon.UsageLimits,
		coupon.MaximumDiscountPrice,
		coupon.MinimumPurchasePrice,
		coupon.ExpiryDate,
		CouponId).
		Scan(&updateCoupon).
		Error
	return updatedCoupon, err
}