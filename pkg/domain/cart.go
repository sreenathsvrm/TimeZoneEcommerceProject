package domain

type Cart struct {
	Id          uint `gorm:"primaryKey;unique;not null"`
	User_id     uint
	Users       Users `gorm:"foreignKey:User_id"`
	Is_applied  bool
	Discount    float64 `json:"discount" gorm:"not null"`
	Total_price float64 `json:"total_price" gorm:"not null"`
}

type CartItem struct {
	Id        uint `json:"id" gorm:"primaryKey;not null"`
	CartID    uint `json:"cart_id"`
	Cart      Cart
	ProductId uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"-"`
	Qty       uint    `json:"qty" gorm:"not null"`
}
