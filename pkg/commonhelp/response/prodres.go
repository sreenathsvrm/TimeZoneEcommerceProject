package response

type Category struct {
	ID           uint   `json:"id" gorm:"unique;not null"`
	CategoryName string `json:"name"`
}

type Product struct {
	Id           int `json:",omitempty"`
	Name         string
	Description  string
	Prize        int
	Qty_in_stock int
	Brand        string
	Category_Id  uint
	CategoryName string
}

type Cartres struct {
	Product_Id uint   `json:"product_item_id"`
	ProductName   string `json:"product_name"`
	Prize         uint   `json:"prize"`
	Qty_in_stock    uint   `json:"qty_in_stock"`
	Qty           uint   `json:"qty"`
}