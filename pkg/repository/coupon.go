package repository

import (
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	"errors"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CouponDB struct {
	DB *gorm.DB
}

func NewCouponrepo(DB *gorm.DB) interfaces.CouponRepo {
	return &CouponDB{
		DB: DB,
	}
}

func (c *CouponDB) AddCoupon(ctx context.Context, coupon domain.Coupon) error {
	createCoupen := `INSERT INTO coupons (code, discount_percent,usage_limits,maximum_discount_price,minimum_purchase_price,expiry_date)
		VALUES($1,$2,$3,$4,$5,$6)`
	if c.DB.Exec(createCoupen,
		coupon.Code,
		coupon.DiscountPercent,
		coupon.UsageLimits,
		coupon.MaximumDiscountPrice,
		coupon.MinimumPurchasePrice,
		coupon.ExpiryDate).Error != nil {
		return errors.New("error")
	}
	return nil
}

func (c *CouponDB) UpdateCouponById(ctx context.Context, CouponId int, coupon requests.Coupon) (UpdatedCoupon domain.Coupon, err error) {
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
		Scan(&UpdatedCoupon).
		Error
	return UpdatedCoupon, err
}

func (c *CouponDB) DeleteCoupon(ctx context.Context, CouponId int) (err error) {
	delete := `DELETE FROM coupons WHERE id=?`

	fmt.Println(CouponId)
	err = c.DB.Exec(delete, CouponId).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CouponDB) ViewCoupons(ctx context.Context) ([]domain.Coupon, error) {
	var listofcoupens []domain.Coupon
	viewall := `SELECT * FROM coupons`
	err := c.DB.Raw(viewall).Scan(&listofcoupens).Error
	return listofcoupens, err

}

func (c *CouponDB) ViewCoupon(ctx context.Context, couponID int) (domain.Coupon, error) {
	var coupon domain.Coupon
	CoupenDetails := `SELECT * FROM coupons WHERE id=$1`
	err := c.DB.Raw(CoupenDetails, couponID).Scan(&coupon).Error
	return coupon, err
}

func (c *CouponDB) GetByCode(ctx context.Context, couponCode string) (coupon domain.Coupon, err error) {
	quary := `SELECT * FROM coupons WHERE code=$1`
	if err := c.DB.Raw(quary, couponCode).Scan(&coupon).Error; err != nil {
		return coupon, err
	}
	return coupon, nil
}

func (c *CouponDB) UpdateCouponByCode(ctx context.Context, code string, coupon domain.Coupon) error {
	var Updated domain.Coupon
	updateCoupon := `UPDATE coupons SET discount_percent=$1, usage_limits=$2, maximum_discount_price=$3, minimum_purchase_price=$4, expiry_date=$5
		 WHERE code=$6
		 RETURNING code, discount_percent, usage_limits, maximum_discount_price, minimum_purchase_price, expiry_date`

	err := c.DB.Raw(updateCoupon,
		coupon.DiscountPercent,
		coupon.UsageLimits,
		coupon.MaximumDiscountPrice,
		coupon.MinimumPurchasePrice,
		coupon.ExpiryDate,
		code).
		Scan(&Updated).Error
	return err
}

func (c *CouponDB) ApplyCoupontoCart(ctx context.Context, userID int, Code string) (float64, error) {
	tx := c.DB.Begin()
	var coupon domain.Coupon
	findCoupon := `SELECT * FROM coupons WHERE code = $1`
	err := tx.Raw(findCoupon, Code).First(&coupon).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("coupon not available")
		}
		tx.Rollback()
		return 0, err
	}

	fmt.Println(coupon.Id, coupon.Code)
	if coupon.Id == 0 {
		return 0, fmt.Errorf("coupons not available")
	}
	if coupon.ExpiryDate.Before(time.Now()) {
		tx.Rollback()
		return 0, fmt.Errorf("coupon expired")
	}

	// check whether the coupen is alredy used by the user in any other previous orders
	if coupon.UsageLimits <= 0 {
		return 0, errors.New("coupon usage limit has been reached")
	}
	// check whether the coupen is alresy added to the cart
	var cart domain.Cart
	getCartDetails := `SELECT * FROM carts WHERE user_id=?`
	err = tx.Raw(getCartDetails, userID).Scan(&cart).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if cart.Is_applied {
		return 0, fmt.Errorf("this coupen is allready applied")
	}
	var discount float64
	//check there is some thing inside the cart
	if cart.Total_price == 0 {
		tx.Rollback()
		return 0, fmt.Errorf("no product is in the cart to apply coupen")
	}

	fmt.Println(coupon)
	// Apply the discount to the cart or mark it as invalid
	if coupon.MinimumPurchasePrice < float64(cart.Total_price) {
		discount = (float64(cart.Total_price) * float64(coupon.DiscountPercent)) / 100
		if discount > float64(coupon.MaximumDiscountPrice) {
			discount = float64(coupon.MaximumDiscountPrice)
		}
	} else {
		return 0, fmt.Errorf("minimum cart prize required")
	}
	coupon.UsageLimits--

	updatecouponuse := `UPDATE coupons SET usage_limits =$1`
	err = tx.Exec(updatecouponuse, coupon.UsageLimits).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	fmt.Println(discount)
	// update the cart total with the subtotal - discount amount
	updatedCart := `UPDATE carts SET total_price=$1,is_applied='T' WHERE id=$2`
	err = tx.Exec(updatedCart, cart.Total_price-discount, cart.Id).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return (cart.Total_price - discount), nil

}

// func (c *CouponDB) RemoveCouponFromCart(ctx context.Context, userID int) error {
// 	tx := c.DB.Begin()

// 	// Check if the cart exists for the user
// 	var cart domain.Cart
// 	getCartDetails := `SELECT * FROM carts WHERE user_id=?`
// 	err := tx.Raw(getCartDetails, userID).Scan(&cart).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if !cart.Is_applied {
// 		// Coupon is not applied, nothing to remove
// 		return nil
// 	}

// 	// Reset the cart details
// 	resetCart := `UPDATE carts SET total_price=0, coupon_id=NULL, is_applied='F' WHERE id=?`
// 	err = tx.Exec(resetCart, cart.Id).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// Increase the usage limit of the previously applied coupon
// 	var coupon domain.Coupon
// 	getCoupon := `SELECT * FROM coupons WHERE id=?`
// 	err = tx.Raw(getCoupon, cart.Coupon_id).Scan(&coupon).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	coupon.UsageLimits++
// 	updateCouponUse := `UPDATE coupons SET usage_limits=? WHERE id=?`
// 	err = tx.Exec(updateCouponUse, coupon.UsageLimits, coupon.Id).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err = tx.Commit().Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	return nil
// }
