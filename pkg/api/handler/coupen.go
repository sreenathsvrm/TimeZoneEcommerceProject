package handler

import (
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	"ecommerce/pkg/domain"
	services "ecommerce/pkg/usecase/interface"
	"net/http"

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
   copier.Copy(&data,&body)
	err = cr.CouponUsecase.CreateCoupon(ctx,data)
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
