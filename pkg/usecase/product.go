package usecase

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
)

type ProductUsecase struct {
	ProductRepo interfaces.ProductRepo
}

func NewProductUsecase(ProductRepo interfaces.ProductRepo) services.ProductUsecase {
	return &ProductUsecase{
		ProductRepo: ProductRepo,
	}
}

func (P *ProductUsecase) Addcategory(ctx context.Context, req urequest.Category) (response.Category, error) {
	addcatagory, err := P.ProductRepo.Addcategory(ctx, req)
	return addcatagory, err
}

func (P *ProductUsecase) UpdatCategory(ctx context.Context, category urequest.Category, id int) (response.Category, error) {
	updatedcatagory, err := P.ProductRepo.UpdatCategory(ctx, category, id)
	return updatedcatagory, err
}
func (P *ProductUsecase) DeleteCatagory(ctx context.Context, Id int) error {
	err := P.ProductRepo.DeleteCatagory(ctx, Id)
	return err
}

func (p *ProductUsecase) Listallcatagory(ctx context.Context) ([]response.Category, error) {
	Allcatagory, err := p.ProductRepo.Listallcatagory(ctx)

	return Allcatagory, err
}

func (p *ProductUsecase) ShowCatagory(ctx context.Context, Id int) (response.Category, error) {

	yourcategory, err := p.ProductRepo.ShowCatagory(ctx, Id)

	return yourcategory, err

}

func (p *ProductUsecase) SaveProduct(ctx context.Context, product urequest.Product) (response.Product, error) {
	newproduct, err := p.ProductRepo.SaveProduct(ctx, product)
	return newproduct, err
}

func (p *ProductUsecase) UpdateProduct(ctx context.Context, id int, product urequest.Product) (response.Product, error) {
	updateproduct, err := p.ProductRepo.UpdateProduct(ctx, id, product)
	return updateproduct, err

}

func (p *ProductUsecase) DeleteProduct(ctx context.Context, id int) error {

	err := p.ProductRepo.DeleteCatagory(ctx, id)

	return err
}

func (p *ProductUsecase) ViewAllProducts(ctx context.Context, pagination urequest.Pagination) (products []response.Product, err error) {
	allProducts, err := p.ProductRepo.ViewAllProducts(ctx, pagination)
	return allProducts, err
}


func (p *ProductUsecase) VeiwProduct(ctx context.Context,id int)(response.Product,error)  {
	product,err:=p.ProductRepo.ViewProduct(ctx,id)
	return product,err
}