package repository

import (
	"context"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type cartDB struct {
	DB *gorm.DB
}

func NewecartRepository(DB *gorm.DB) interfaces.CartRepo {
	return &cartDB{
		DB: DB,
	}
}

func (c *cartDB) SaveCart(ctx context.Context, Userid int) (uint, error) {
	var cart domain.Cart
	query := `INSERT INTO carts (user_id,discount,total_price) VALUES($1, $2,$3) RETURNING id`
	if c.DB.Raw(query, Userid, 0,0).Scan(&cart).Error != nil {
		return 0, fmt.Errorf("faild to save cart for user")
	}
	fmt.Println(cart)
	return cart.Id, nil
}
func (c *cartDB) FindCartByUserID(ctx context.Context, UserID int) (domain.Cart, error) {
	var cart domain.Cart
	Find := `SELECT * FROM carts WHERE user_id = ?`
	if c.DB.Raw(Find, UserID).Scan(&cart).Error != nil {
		return cart, errors.New("faild to get cart for this id")
	}
	return cart, nil
}

func (c *cartDB) AddCartItem(ctx context.Context, CartItem domain.CartItem) error {
	Query := `INSERT INTO cart_items(cart_id,product_id,qty)VALUES($1,$2,$3)`
	if c.DB.Raw(Query, CartItem.CartID, CartItem.ProductId, 1).Scan(&CartItem).Error != nil {
		return errors.New("cant add this item")
	}

	return nil
}

func (c *cartDB) FindCartIDNproductId(ctx context.Context, cart_id uint, product_id uint) (cartItem domain.CartItem, err error) {

	Query := `SELECT * FROM cart_items WHERE cart_id = $1 AND product_id = $2`
	if c.DB.Raw(Query, cart_id, product_id).Scan(&cartItem).Error != nil {
		return cartItem, errors.New("cant find cartitem coresponding this cart id ,product id")
	}
	return cartItem, nil
}

func (c *cartDB) FindProduct(ctx context.Context, id uint) (response.Product, error) {
	var product response.Product
	query := `SELECT p.id,p.product_name as name,p.description,p.brand,p.prize,p.category_id,p.qty_in_stock,c.category_name,p.created_at,p.updated_at FROM products p 
		JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(query, id).Scan(&product).Error
	return product, err
}
func (c *cartDB) RemoveCartItem(ctx context.Context, CartItemid uint) error {
	RemoveQuery := `DELETE FROM cart_items WHERE id=$1	`
	if c.DB.Exec(RemoveQuery, CartItemid).Error != nil {
		return errors.New("faild to remove product_items from cart")
	}
	return nil
}


func (c *cartDB) AddQuantity(ctx context.Context, cartItemid uint, qty uint) error {

	query := `UPDATE cart_items SET qty = $1 WHERE id = $2`
	if c.DB.Exec(query, qty, cartItemid).Error != nil {
		return errors.New("faild to add  qty of ")
	}
	return nil
}

func (c *cartDB) FindCartlistByCartID(ctx context.Context, cartID uint) (cartitems []response.Cartres, err error) {

	query := `SELECT ci.product_id, p.product_name, ci.qty, p.prize, p.qty_in_stock 
	FROM cart_items ci 
	INNER JOIN products p ON ci.product_id = p.id 
	WHERE ci.cart_id = ?;
	`
	if c.DB.Raw(query,cartID).Scan(&cartitems).Error !=nil{
		return cartitems, errors.New("failed to show cartitems")
	}
	return cartitems,err
}