package domain

import "time"

type Users struct {
	ID        uint   `gorm:"primaryKey;unique;not null"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Mobile    string `json:"mobile" binding:"required,eq=10" gorm:"unique; not null"`
	Password  string `json:"password" gorm:"not null"`
	IsBlocked bool   `gorm:"default:false"`
	CreatedAt time.Time
}

type UserStatus struct {
	ID                uint `gorm:"primaryKey"`
	UsersID           uint
	Users             Users `gorm:"foreignKey:UsersID"`
	BlockedAt         time.Time
	BlockedBy         uint
	ReasonForBlocking string
}

type Cart struct {
	Id          uint `gorm:"primaryKey;unique;not null"`
	User_id     uint
	Users       Users `gorm:"foreignKey:User_id"`
	Total_price uint  `json:"total_price" gorm:"not null"`
}

type CartItem struct {
	Id        uint `json:"id" gorm:"primaryKey;not null"`
	CartID    uint `json:"cart_id"`
	Cart      Cart
	ProductId uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"-"`
	Qty       uint    `json:"qty" gorm:"not null"`
}

type Address struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Users       Users  `gorm:"foreignKey:UserID" json:"-"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
}