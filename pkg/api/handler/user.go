package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce/pkg/api/utilhandler"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/commonhelp/urequest"
	services "ecommerce/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

// type Response struct {
// 	ID    uint   `copier:"must"`
// 	Name  string `copier:"must"`
// 	Email string `copier:"must"`
// }

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary UserSignUp
// @ID UserSignup
// @Description Create a new user with the specified details.
// @Tags Users
// @Accept json
// @Produce json
// @Param   inputs   body     urequest.Fusersign{}   true  "Input Field"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /signup [post]
func (cr *UserHandler) UserSignup(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	// defer cancel()
	var user urequest.Fusersign
	err := c.Bind(&user)
	fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	UserValue, err := cr.userUseCase.UserSignup(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       UserValue,
		Errors:     nil,
	})

}

// LoginWithEmail
// @Summary User Login
// @ID UserLogin
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param   input   body     urequest.Flogin{}   true  "Input Field"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /login [post]
func (cr *UserHandler) UserLogin(c *gin.Context) {

	var user urequest.Flogin
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ss, err := cr.userUseCase.UserLogin(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logined successfuly",
		Data:       nil,
		Errors:     nil,
	})
}

//Home
//@Summery Homepage
//ID homepage
//Discription landing page for users
//tags HOME
// @Success 200 "success"
// @Failure 400 "failed"
// @Router /home [GET]
func (cr *UserHandler) Home(c *gin.Context) {

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "WELCOME to HOMEPAGE",
		Data:       nil,
		Errors:     nil,
	})

}

// UserLogout
// @Summary User Login
// @ID UserLogout
// @Description Logout as a user exit from the ecommerce site
// @Tags Users
// @Success 200 "success"
// @Failure 400 "failed"
// @Router /logout [post]
func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logout successfully",
		Data:       nil,
		Errors:     nil,
	})

}

// @Summary AddAdrress_for_user
// @ID Add_Adress
// @Description Create a new user with the specified details.
// @Tags Users
// @Accept json
// @Produce json
// @Param   inputs   body     urequest.AddressReq{} true  "Input Field"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /SaveAddress [post]
func (cr *UserHandler) AddAdress(c *gin.Context) {
	var newAddress urequest.AddressReq
	err := c.Bind(&newAddress)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "failed to read request body",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	UserID, err := utilhandler.GetUserIdFromContext(c)

	Address, err := cr.userUseCase.AddAdress(c, UserID, newAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant add this adress",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "SuccessFully Added YOUR ADRESS",
		Data:       Address,
		Errors:     nil,
	})
}




// @Summary updateAdrress_for_user
// @ID Update_Adress
// @Description Update user Adresses.
// @Tags Users
// @Accept json
// @Produce json
// @Param   inputs   body     urequest.AddressReq{} true  "Input Field"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /UpdateAddress [patch]
func (cr *UserHandler) UpdateAdress(c *gin.Context) {
	var UpdatedAddress urequest.AddressReq
	err := c.Bind(&UpdatedAddress)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "failed to read request body",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	UserID, err := utilhandler.GetUserIdFromContext(c)

	Address,err:=cr.userUseCase.UpdateAdress(c,UserID,UpdatedAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant update this adress",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "SuccessFully updated YOUR ADRESS",
		Data:       Address,
		Errors:     nil,
	})
}


// viewAdress godoc
// @summary api for get address of user
// @description user can see their Adress
// @security ApiKeyAuth
// @id User_Address
// @tags  Users
// @Router /viewAddress [get]
// @Success 200 {object} response.Response{} "successfully get Address"
// @Failure 500 {object} response.Response{} "faild to get Address"
func (c *UserHandler)VeiwAddress(ctx *gin.Context)  {

	userID,err:=utilhandler.GetUserIdFromContext(ctx)
	Adress,err:=c.userUseCase.VeiwAdress(ctx,userID)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message:    "something Went Wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
     {
		ctx.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "this your Adress",
			Data:       Adress,
			Errors:     nil,
		})
		return
	}
}
