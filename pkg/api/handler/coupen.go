package handler

import (
	"ecommerce/pkg/api/utilhandler"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	services "ecommerce/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CouponHandler struct {
	CouponUsecase services.CouponUseCase
}

func NewCoupenHandler(CouponUsecase services.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		CouponUsecase: CouponUsecase,
	}
}

// AddCoupon godoc
// @summary api for add Coupons for ecommerce
// @description Admin can add coupon
// @security ApiKeyAuth
// @id AddCoupon
// @tags Coupon
// @Param input body urequest.Coupon true "Input true info"
// @Router /admin/coupon/AddCoupons [post]
// @Success 200 "Successfully productItem added to cart"
// @Failure 400 "can't add the product item into cart"
func (cr *CouponHandler) AddCoupon(ctx *gin.Context) {
	var body urequest.Coupon
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    " body binding faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var data domain.Coupon
	copier.Copy(&data, &body)
	err = cr.CouponUsecase.CreateCoupon(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "fAILED TO CREATE COUPON",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "successfully Create a coupon",
		Data:       body,
		Errors:     nil,
	})

}

// UpdateCoupon
// @Summary Admin can update existing coupon
// @ID update-coupon
// @Description Admin can update existing coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param CouponID path int true "CouponID"
// @Param coupon_details body urequest.Coupon true "details of coupon to be updated"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/Update/{CouponID} [patch]
func (c *CouponHandler) UpdateCoupon(ctx *gin.Context) {
	id := ctx.Param("CouponID")
	CouponID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var body urequest.Coupon
	err = ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    " body binding faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updated, err := c.CouponUsecase.UpdateCouponById(ctx, CouponID, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "fAILED TO UPDATE COUPON",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "successfully updated a coupon",
		Data:       updated,
		Errors:     nil,
	})
}

// REMOVECOUPON
// @Summary Admin can delete a coupon
// @ID delete-coupon
// @Description Admin can delete a coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param CouponID path string true "CouponID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/Delete/{CouponID} [delete]
func (cr *CouponHandler) DeleteCoupon(ctx *gin.Context) {
	Id := ctx.Param("CouponID")
	CouponID, err := strconv.Atoi(Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.CouponUsecase.DeleteCoupon(ctx, CouponID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't delete coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupon deleted",
		Data:       nil,
		Errors:     nil,
	})

}

// FindCouponByID
// @Summary Admins  can see Coupons with coupon_id
// @ID find-Coupon-by-id
// @Description Admins can see Coupons with coupon_id
// @Tags Coupon
// @Accept json
// @Produce json
// @Param id path string true "CouponID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/Viewcoupon/{id} [get]
func (cr *CouponHandler) ViewCoupon(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find coupon with this id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Coupon, err := cr.CouponUsecase.ViewCoupon(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupon",
		Data:       Coupon,
		Errors:     nil,
	})
}

// ListAllcoupons
// @Summary for geting all order status list
// @ID List-all-coupons
// @Description Endpoint for getting all coupons
// @Tags Coupon
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/couponlist [get]
func (cr *CouponHandler) Coupons(ctx *gin.Context) {
	Coupons, err := cr.CouponUsecase.ViewCoupons(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't List the coupons",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    " here the list of coupons ",
		Data:       Coupons,
		Errors:     nil,
	})

}

// ApplayCoupon
// @Summary User can apply a coupon to the cart
// @ID applay-coupon-to-cart
// @Description User can apply coupon to the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param code query string true "code"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /coupon/apply/{code} [patch]
func (cr *CouponHandler) ApplyCoupon(ctx *gin.Context) {
	UserID, err := utilhandler.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Code := ctx.Query("code")
	discount, err := cr.CouponUsecase.ApplyCoupontoCart(ctx, UserID, Code)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "can't apply coupen",
				Data:       nil,
				Errors:     err.Error(),
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupen applyed",
		Data:       []interface{}{"rate after coupen applaid is ", discount},
		Errors:     nil,
	})

}
