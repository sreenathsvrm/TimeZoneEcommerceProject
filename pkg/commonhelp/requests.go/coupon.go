package requests

import "time"

type Coupon struct {
	DiscountPercent      float64 `json:"discountpercent" binding:"required"`
	UsageLimits          int   `json:"usagelimit" binding:"required"`
	MaximumDiscountPrice float64 `json:"maximumdiscountprice" binding:"required"`
	MinimumPurchasePrice float64  `json:"minimumpurchaseamount" binding:"required"`
	ExpiryDate           time.Time  `json:"expirationdate" binding:"required"`
}
  