package repository

import (
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	interfaces "ecommerce/pkg/repository/interface"
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type productDB struct {
	DB *gorm.DB
}

func NewproductRepository(DB *gorm.DB) interfaces.ProductRepo {
	return &productDB{
		DB: DB,
	}
}

func (p *productDB) Addcategory(ctx context.Context, req urequest.Category) (response.Category, error) {
	var categoryname response.Category
	query := `INSERT INTO categories (category_name ,created_at)VAlues($1,NOW())RETURNING id,category_name as name`
	err := p.DB.Raw(query, req.Name).Scan(&categoryname).Error
	return categoryname, err

}

func (c *productDB) UpdatCategory(ctx context.Context, category urequest.Category, id int) (response.Category, error) {
	var updatedCategory response.Category
	query := `UPDATE  categories SET category_name = $1 , updated_at =NOW() WHERE id=$2 RETURNING id,category_name `
	err := c.DB.Raw(query, category.Name, id).Scan(&updatedCategory).Error
	return updatedCategory, err
}

func (c *productDB) DeleteCatagory(ctx context.Context, Id int) error {
	Query := `DELETE FROM categories WHERE id=?`
	err := c.DB.Exec(Query, Id).Error
	return err

}

func (c *productDB) Listallcatagory(ctx context.Context) ([]response.Category, error) {
	var Allcatagory []response.Category
	Query := `SELECT * FROM categories`
	err := c.DB.Raw(Query).Scan(&Allcatagory).Error
	fmt.Println(Allcatagory)
	return Allcatagory, err
}

func (c *productDB) ShowCatagory(ctx context.Context, Id int) (response.Category, error) {
	var catagory response.Category

	Query := `SELECT *FROM categories WHERE id=$1`
	err := c.DB.Raw(Query, Id).Scan(&catagory).Error
	return catagory, err
}

func (c *productDB) SaveProduct(ctx context.Context, product urequest.Product) (response.Product, error) {
	var Newproduct response.Product
	var exits bool
	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, product.Category_Id).Scan(&exits)
	if !exits {
		return response.Product{}, fmt.Errorf("this catagory is not found ")
	}

	query := `INSERT INTO products (product_name, description ,brand ,prize,qty_in_stock,category_id, created_at)VALUES($1,$2,$3,$4,$5,$6,NOW())
	RETURNING id, product_name as name, description, brand, prize, category_id `
	fmt.Println(product)
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.Prize, product.Qty_in_stock,product.Category_Id).
		Scan(&Newproduct).Error

	return Newproduct, err

}

func (c *productDB) UpdateProduct(ctx context.Context, id int, product urequest.Product) (response.Product, error) {

	var Newproduct response.Product

	query := `UPDATE products SET product_name = $1, description = $2, brand = $3, prize = $4, qty_in_stock = $5, category_id = $6, updated_at = NOW() WHERE id = $7 
	RETURNING id, product_name as name, description, brand, prize, qty_in_stock, category_id`



	err := c.DB.Raw(query, product.Name, product.Description, product.Brand,
		product.Prize,product.Qty_in_stock, product.Category_Id, id).Scan(&Newproduct).Error

	return Newproduct, err

}

func (c *productDB) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

func (c *productDB) ViewAllProducts(ctx context.Context, pagination urequest.Pagination) (products []response.Product, err error) {

	limit := pagination.Page
	offset := (pagination.PerPage - 1) * limit

	// aliase :: p := product; c := category

	Query := `SELECT p.id,p.product_name,p.description,p.brand,p.prize,p.qty_in_stock,p.category_id,c.category_name,p.created_at,p.updated_at
	FROM products p LEFT JOIN categories c ON p.category_id=c.id
	ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	if c.DB.Raw(Query, limit, offset).Scan(&products).Error != nil {
		return products, errors.New("failed to get products from database")
	}
	return products, nil
}

func (c *productDB) ViewProduct(ctx context.Context, id int) (response.Product, error) {
	var product response.Product
	query := `SELECT p.id,p.product_name as name,p.description,p.brand,p.prize,p.category_id,p.qty_in_stock,c.category_name,p.created_at,p.updated_at FROM products p 
		JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(query, id).Scan(&product).Error
	return product, err
}
