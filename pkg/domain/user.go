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
	Is_applied    bool
	Discount     float64   `json:"discount" gorm:"not null"`
	Total_price float64  `json:"total_price" gorm:"not null"`
	 
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

type Wallet struct {
	ID          uint `json:"wallet_id" gorm:"primaryKey;not null"`
	UserID      uint `json:"user_id" gorm:"not null"`
	TotalAmount uint `json:"total_amount" gorm:"not null"`
}

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

type Transaction struct {
	TransactionID   uint            `json:"transction_id" gorm:"primaryKey;not null"`
	WalletID        uint            `json:"wallet_id" gorm:"not null"`
	Wallet          Wallet          `json:"-"`
	TransactionDate time.Time       `json:"transaction_time" gorm:"not null"`
	Amount          uint            `josn:"amount" gorm:"not null"`
	TransactionType TransactionType `json:"transaction_type" gorm:"not null"`
}