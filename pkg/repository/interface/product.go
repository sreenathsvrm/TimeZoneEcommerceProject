package interfaces

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
)

type ProductRepo interface {
	Addcategory(ctx context.Context, req requests.Category) (response.Category, error)
	UpdatCategory(ctx context.Context, category requests.Category, id int) (response.Category, error)
	DeleteCatagory(ctx context.Context, Id int) error
	Listallcatagory(ctx context.Context) ([]response.Category, error)
	ShowCatagory(ctx context.Context, Id int) (response.Category, error)
	SaveProduct(ctx context.Context, product requests.Product) (response.Product, error)
	UpdateProduct(ctx context.Context, id int, product requests.Product) (response.Product, error)
	DeleteProduct(ctx context.Context, id int) error
	ViewAllProducts(ctx context.Context, pagination requests.Pagination) (products []response.Product, err error)
	ViewProduct(ctx context.Context, id int) (response.Product, error)
}
