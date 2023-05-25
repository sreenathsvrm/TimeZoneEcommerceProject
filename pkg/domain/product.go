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
