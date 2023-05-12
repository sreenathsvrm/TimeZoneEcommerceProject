package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	
)

type ProductUsecase interface{
	Addcategory(ctx context.Context,req urequest.Category)(response.Category,error)
	UpdatCategory(ctx context.Context,category urequest.Category, id int) (response.Category, error)
	DeleteCatagory(ctx context.Context,Id int)error
	Listallcatagory(ctx context.Context) ([]response.Category, error)
	ShowCatagory(ctx context.Context, Id int) (response.Category, error) 
	SaveProduct(ctx context.Context, product urequest.Product) (response.Product, error)
	UpdateProduct(ctx context.Context,id int ,product urequest.Product) (response.Product,error)
	DeleteProduct(ctx context.Context,id int)error
	ViewAllProducts(ctx context.Context, pagination urequest.Pagination) (products []response.Product, err error) 
    VeiwProduct(ctx context.Context,id int)(response.Product,error)
}