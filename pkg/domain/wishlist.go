package domain



type WishList struct {
	ID            uint `json:"-" gorm:"primaryKey;not null"`
	UserID        uint `json:"user_id" gorm:"not null"`
	User          Users  `json:"-"`
	ProductID     uint `json:"product_id" gorm:"not null"`
	Product        Product   `json:"-"`
}

