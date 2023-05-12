package urequest

type Cartreq struct {
	ProductId int
	UserID    int  `json:"-"`
}

type Addcount struct {
	UserID        int   `json:"-"`
	ProductId int `json:"product_id" binding:"required"`
    Count   uint `json:"count" binding:"omitempty,gte=1"`
}

type CartItems struct {
    ProductId int
    Qty      int
    Price         int
    Qty_In_Stock    int
}
