package domain

import "time"

type Coupon struct {
	Id                   uint   `gorm:"primaryKey;unique;not null"`
	Code                 string ``
	DiscountPercent      float64
	UsageLimits          int
	MaximumDiscountPrice float64
	MinimumPurchasePrice float64
	ExpiryDate           time.Time
}
