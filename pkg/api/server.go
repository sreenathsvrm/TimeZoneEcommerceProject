package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ecommerce/cmd/api/docs"
	handler "ecommerce/pkg/api/handler"
	"ecommerce/pkg/api/middleware"
	//	middleware "ecommerce/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler,
	AdminHandler *handler.AdminHandler, ProductHandler *handler.ProductHandler, CartHandler *handler.CartHandler, OrderHandler *handler.OrderHandler,CouponHandler *handler.CouponHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	engine.LoadHTMLGlob("./*html")

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT

	// Auth middleware

	user := engine.Group("/")
	{

		user.POST("signup", userHandler.UserSignup)
		user.POST("login", userHandler.UserLogin)
		user.POST("otp/send", otpHandler.SendOtp)
		user.POST("otp/verify", otpHandler.ValidateOtp)
		user.GET("home", userHandler.Home)
	
		

	}
	user.Use(middleware.UserAuth)
	{

		user.POST("SaveAddress", userHandler.AddAdress)
		user.PATCH("UpdateAddress", userHandler.UpdateAdress)
		user.GET("viewAddress", userHandler.VeiwAddress)
		user.POST("Addwishlist/:id",userHandler.AddToWishList)
		user.DELETE("/Removewishlist/:id",userHandler.RemoveFromWishList)
		user.GET("wishlist",userHandler.GetWishList)
		user.POST("logout", userHandler.UserLogout)

		category := user.Group("/category")
		{
			category.GET("showall/", ProductHandler.ListCategories)
			category.GET("disply/:id", ProductHandler.DisplayCategory)
		}
		product := user.Group("/product")
		{
			product.GET("ViewAllProducts", ProductHandler.ViewAllProducts)
		}
		cart := user.Group("/cart")
		{
			cart.POST("/AddToCart", CartHandler.AddCartItem)
			cart.DELETE("/RemoveFromCart", CartHandler.RemoveFromCart)
			cart.PUT("/Addcount", CartHandler.Addcount)
			cart.GET("/viewcart", CartHandler.ViewCartItems)
			
		}
		coupon:=user.Group("/coupon")
		{
			coupon.GET("/coupons",CouponHandler.UserCoupons)
			coupon.PATCH("/apply/:code", CouponHandler.ApplyCoupon)
		}
		order := user.Group("/order")
		{
			order.POST("/orderAll/:payment_id", OrderHandler.CashonDElivery)
			order.GET("/razor/:payment_id", OrderHandler.RazorpayCheckout)
			order.POST("/razor/success", OrderHandler.RazorpayVerify)
			order.PATCH("/cancel/:orderId", OrderHandler.CancelOrder)
			order.GET("/view/:order_id", OrderHandler.ListOrder)
			order.GET("/listall", OrderHandler.ListAllOrders)
			order.PATCH("/return/:orderId", OrderHandler.ReturnOrder)
		}
	}

	// admin side
	admin := engine.Group("/admin")

	admin.POST("/signup", AdminHandler.SaveAdmin)
	admin.POST("/login", AdminHandler.LoginAdmin)
	admin.POST("/logout", AdminHandler.AdminLogout)

	admin.Use(middleware.AdminAuth)
	{
		admin.GET("/findall", AdminHandler.FindAllUser)
		admin.GET("/finduser/:user_id", AdminHandler.FindUserByID)
		admin.PATCH("/block", AdminHandler.BlockUser)
		admin.PATCH("/unblock/:user_id", AdminHandler.UnblockUser)
		admin.GET("/salesreport",AdminHandler.ViewSalesReport)
		admin.GET("/salesreport/download",AdminHandler.DownloadSalesReport)
		//category
		category := admin.Group("/category")
		{
			category.POST("add", ProductHandler.Addcategory)
			category.PATCH("update/:id", ProductHandler.UpdatCategory)
			category.DELETE("delete/:category_id", ProductHandler.DeleteCategory)
			category.GET("showall/", ProductHandler.ListCategories)
			category.GET("disply/:id", ProductHandler.DisplayCategory)
		}
		product := admin.Group("/product")
		{

			product.POST("save", ProductHandler.SaveProduct)
			product.PATCH("updateproduct/:id", ProductHandler.UpdateProduct)
			product.DELETE("delete/:product_id", ProductHandler.DeleteProduct)
			product.GET("ViewAllProducts", ProductHandler.ViewAllProducts)
			product.GET("ViewProduct/:id", ProductHandler.VeiwProduct)
		}

		order := admin.Group("/order")
		{
			order.GET("/Status", OrderHandler.Statuses)
			order.GET("/Allorders", OrderHandler.AllOrders)
			order.PATCH("/UpdateStatus", OrderHandler.UpdateOrderStatus)
		}
		coupon:=admin.Group("/coupon")
		{
			coupon.POST("/AddCoupons",CouponHandler.AddCoupon)
			coupon.PATCH("/Update/:CouponID",CouponHandler.UpdateCoupon)
			coupon.DELETE("/Delete/:CouponID",CouponHandler.DeleteCoupon)
			coupon.GET("/Viewcoupon/:id",CouponHandler.ViewCoupon)
			coupon.GET("/couponlist",CouponHandler.Coupons)

		}
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
