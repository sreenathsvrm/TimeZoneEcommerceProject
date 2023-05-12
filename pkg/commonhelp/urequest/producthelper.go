package urequest



type Category struct {
	Name string `json:"name" validate:"required"`
}

type Product struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Prize       int
	Qty_in_stock int
	Category_Id  string `json:"categoryid" validate:"required"`
}

