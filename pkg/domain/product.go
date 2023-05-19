package domain

import "time"

type Category struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	CategoryName string `gorm:"unique;not null"`
	Created_at   time.Time
	Updated_at   time.Time
}

type Product struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	ProductName  string `gorm:"unique;not null"`
	Description  string
	Brand        string
	Prize        int
	Qty_in_stock int
	Category_id  uint
	Category     Category `gorm:"foreignKey:Category_id"`
	Created_at   time.Time
	Updated_at   time.Time
}

type PaymentMethod struct {
	ID            uint `gorm:"primaryKey"`
	PaymentMethod string  `json:"payment_method"`
}

type PaymentStatus struct{
	ID   uint   `gorm:"primaryKey"`
	PaymentStatus string `json:"payment_status,omitempty"`
}

type Orders struct {
	ID                uint           `gorm:"primaryKey"`
	UserID            uint           `json:"user_id"`
	Users             Users          `gorm:"foreignKey:UserID" json:"-"`
	OrderDate         time.Time      `json:"order_date"`
	PaymentMethodID   uint           `json:"payment_method_id"`
	PaymentMethod     PaymentMethod  `gorm:"foreignKey:PaymentMethodID" json:"-"`
	ShippingAddressID uint           `json:"shipping_address_id"`
	Address           Address        `gorm:"foreignKey:ShippingAddressID" json:"-"`
	OrderTotal        float64        `json:"order_total"`
	OrderStatusID     uint           `json:"order_status_id"`
	OrderStatus       OrderStatus    `gorm:"foreignKey:OrderStatusID" json:"-"`
	DeliveryUpdatedAt time.Time      `json:"delivery_time"`
}

type OrderLine struct {
	ID            uint         `gorm:"primaryKey"`
	ProductID     uint          `json:"product_id"`
	Product      Product       ` json:"-"`
	OrderID       uint         `json:"order_Id"`
	Order         Orders       
	Qty      int         `json:"qty"`
	Price         float64     `json:"price"`
}

type OrderStatus struct {
	ID          uint `gorm:"primaryKey"`
	OrderStatus string
}

type DeliveryStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}

type PaymentDetails struct {
	ID              uint          `gorm:"primaryKey" json:"id,omitempty"`
	OrdersID        uint          `json:"order_id,omitempty"`
	Orders          Orders        `gorm:"foreignKey:OrdersID" json:"-"`
	OrderTotal      float64       `json:"order_total"`
	PaymentMethodID   uint          `json:"payment_method_id"`
	PaymentMethod    PaymentMethod   `gorm:"foreignKey:PaymentMethodID"`
	PaymentStatusID uint          `json:"payment_status_id,omitempty"`
	PaymentStatus   PaymentStatus `gorm:"foreignKey:PaymentStatusID" json:"-"`
	UpdatedAt       time.Time
}
