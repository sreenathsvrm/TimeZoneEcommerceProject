package handler

import (
	"ecommerce/pkg/api/utilhandler"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"

	services "ecommerce/pkg/usecase/interface"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartUsecase services.CartUsecase
}

func NewCartHandler(CartUsecase services.CartUsecase) *CartHandler {
	return &CartHandler{
		CartUsecase: CartUsecase,
	}
}

// AddToCart godoc
// @summary api for add productItem to user cart
// @description user can add a stock in product to user cart
// @security ApiKeyAuth
// @id AddToCart
// @tags Cart
// @Param input body requests.Cartreq true "Input true info"
// @Router /cart/AddToCart [post]
// @Success 200 "Successfully productItem added to cart"
// @Failure 400 "can't add the product item into cart"
func (c *CartHandler) AddCartItem(ctx *gin.Context) {
	var body requests.Cartreq
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "faild to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	body.UserID, err = utilhandler.GetUserIdFromContext(ctx)
	fmt.Println(body, err)
	err = c.CartUsecase.AddCartItem(ctx, body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Add to cart these item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "successfully add to cart",
		Data:       body,
		Errors:     nil,
	})
}

// DeleteCartItem
// @Summary Admin can delete a category
// @ID delete-cartitem
// @Description user can delete their cartitems by id
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body requests.Cartreq{} true "Input Field"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /cart/RemoveFromCart [delete]
func (c *CartHandler) RemoveFromCart(ctx *gin.Context) {
	var body requests.Cartreq
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "faild to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	body.UserID, err = utilhandler.GetUserIdFromContext(ctx)
	fmt.Println(body, err)
	err = c.CartUsecase.RemoveFromCart(ctx, body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Unable to remove  these item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "successfully remove  from cart",
		Data:       body,
		Errors:     nil,
	})
}

// AddQuntity
// @Summary Admin can delete a category
// @ID Add-Qantity
// @Description user can delete their cartitems by id
// @Tags Cart
// @Accept json
// @Produce json
// @Param input body requests.Addcount{} true "Input Field"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /cart/Addcount [put]
func (c *CartHandler) Addcount(ctx *gin.Context) {

	var body requests.Addcount

	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "faild to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	body.UserID, err = utilhandler.GetUserIdFromContext(ctx)
	fmt.Println(body, err)
	err = c.CartUsecase.AddQuantity(ctx, body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Unable to Add count",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "count added",
		Data:       body,
		Errors:     nil,
	})

}

// viewCart godoc
// @summary api for get all cart item of user
// @description user can see all productItem that stored in cart
// @security ApiKeyAuth
// @id Cart
// @tags  Cart
// @Router /cart/viewcart [get]
// @Success 200 {object} response.Response{} "successfully got user cart items"
// @Failure 500 {object} response.Response{} "faild to get cart items"
func (c *CartHandler) ViewCartItems(ctx *gin.Context) {

	userID, err := utilhandler.GetUserIdFromContext(ctx)
	cart, err := c.CartUsecase.FindUserCart(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Unable to get cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if cart.Id == 0 {
		ctx.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "you are not add any products to cart",
			Data:       nil,
			Errors:     nil,
		})
		return
	}

	cartitems, err := c.CartUsecase.FindCartlistByCartID(ctx, cart.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Unable to get cartitems",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	if cartitems == nil {
		ctx.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "sorry no products in your cart",
			Data:       nil,
			Errors:     nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "your carts here",
		Data:       cartitems,
		Errors:     nil,
	})
}
