package usecase

import (
	"context"

	// "ecommerce/pkg/domain"

	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"

	"github.com/pkg/errors"
)

type CartUsecase struct {
	CartRepo interfaces.CartRepo
}

func NewCartUsecase(CartRepo interfaces.CartRepo) services.CartUsecase {
	return &CartUsecase{
		CartRepo: CartRepo,
	}
}

func (c *CartUsecase) AddCartItem(ctx context.Context, body urequest.Cartreq) error {

	//	a. product_id validate (find product by product_id),
	product, err := c.CartRepo.FindProduct(ctx, uint(body.ProductId))
	if err != nil {
		return errors.New("invalied product")
	}
	//   b. check product qty out of stock ; return error if  out of stock
	if product.Qty_in_stock == 0 {
		return errors.New("now product is unavailable ")
	}

	//  a. find user cart with user_id
	cart, err := c.CartRepo.FindCartByUserID(ctx, body.UserID)
	if err != nil {
		return errors.New("belongs to this userid have no cart")
		//	b. if cart not exitst ; create new cart with user_id on table cart
	} else if cart.Id == 0 {
		cartId, err := c.CartRepo.SaveCart(ctx, body.UserID)
		if err != nil {
			errors.New("unable to create cart for this id")
		}
		cart.Id = cartId
	}

	//a. add product_id and cart_id to table cart_items
	cartitem, err := c.CartRepo.FindCartIDNproductId(ctx, cart.Id, uint(body.UserID))
	if err != nil {
		return err
		//	b. with product already exists on cart
	} else if cartitem.Id != 0 {
		return errors.New("product is allready save in cart")
	}

	cartItem := domain.CartItem{
		CartID:    cart.Id,
		ProductId: uint(body.ProductId),
	}

	if err := c.CartRepo.AddCartItem(ctx, cartItem); err != nil {
		return err
	}

	return nil
}
func (c *CartUsecase) FindUserCart(ctx context.Context, userID int) (cart domain.Cart, err error) {
	cart, err = c.CartRepo.FindCartByUserID(ctx, userID)
	if err != nil {
		return cart, err
	}
	return cart, err
}

func (c *CartUsecase) RemoveFromCart(ctx context.Context, body urequest.Cartreq) error {

	//	a. product_id validate (find product by product_id),
	product, err := c.CartRepo.FindProduct(ctx, uint(body.ProductId))
	if err != nil {
		return errors.New("invalied product")
	}
	//   b. check product qty out of stock ; return error if  out of stock
	if product.Id == 0 {
		return errors.New("now product is unavailable ")
	}

	//  a. find user cart with user_id
	cart, err := c.CartRepo.FindCartByUserID(ctx, body.UserID)
	if err != nil {
		return errors.New("belongs to this userid have no cart")
		//	b. if cart not exitst ; create new cart with user_id on table cart
	} else if cart.Id == 0 {
		errors.New("unable to remove from cart becouse your cart is empty")

	}

	//a. check product is allready exit or not in your cart
	cartitem, err := c.CartRepo.FindCartIDNproductId(ctx, cart.Id, uint(body.ProductId))
	if err != nil {
		return err
	} else if cartitem.Id == 0 {
		return errors.New("product is not exist in your cart")
	}

	if err := c.CartRepo.RemoveCartItem(ctx, cartitem.Id); err != nil {
		return err
	}

	return nil

}



func (c *CartUsecase) AddQuantity(ctx context.Context, body urequest.Addcount) error {
	product, err := c.CartRepo.FindProduct(ctx, uint(body.ProductId))
	if err != nil {
		return errors.New("invalied product")
	}
	//   b. check product qty out of stock ; return error if  out of stock
	if product.Id == 0 {
		return errors.New("now product is unavailable ")
	}
    
	if body.Count<0{
		return errors.New("cant add -ve values to quantity")
	}

	if body.Count>uint(product.Qty_in_stock){
       return errors.New("insufficient product quantity in stock")
	}

	
	cart,err:= c.CartRepo.FindCartByUserID(ctx,body.UserID)
	if err!=nil{
		errors.New("user have no cart")
	}
  
	cartitem, err := c.CartRepo.FindCartIDNproductId(ctx, cart.Id, uint(body.ProductId))
	if err != nil {
		return err
	} else if cartitem.Id == 0 {
		return errors.New("product is not exist in your cart")
	}

	err=c.CartRepo.AddQuantity(ctx,cartitem.Id,body.Count); if err!=nil{
        return errors.New("unable to add more quantity")
	}
  
	return nil
}

func (c *CartUsecase)FindCartlistByCartID(ctx context.Context, cartID uint) ( []response.Cartres,  error) {
	  cartitems,err:=c.CartRepo.FindCartlistByCartID(ctx,cartID); if err !=nil{
       return cartitems,err
	  }
	 return cartitems,nil
}

