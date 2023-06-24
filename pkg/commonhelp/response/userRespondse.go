package response

import "time"

type UserValue struct {
	ID    uint    `json:"id" gorm:"unique;not null"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_time"`
 }


type Wishlist struct{
    ProductID uint   `json:"product_item_id"`
	ProductName   string `json:"product_name"`
	Price         uint   `json:"price"`
	Image         string `json:"image"`
	QtyInStock    uint   `json:"qty_in_stock"`

}